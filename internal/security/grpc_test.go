package security_test

import (
	"bytes"
	"context"
	"net"
	"testing"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func TestGRPCRejectsNonCanonicalMethod(t *testing.T) {
	const method = "/grpc.health.v1.Health/Check"
	server := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		requestedMethod, ok := grpc.Method(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "missing method")
		}
		if requestedMethod == method {
			return nil, status.Error(codes.PermissionDenied, "denied")
		}
		return handler(ctx, req)
	}))
	t.Cleanup(server.Stop)
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = listener.Close() })
	go func() { _ = server.Serve(listener) }()

	for _, tc := range []struct {
		name string
		path string
		code string
	}{
		{name: "canonical method is denied", path: method, code: "7"},
		{name: "missing leading slash is rejected", path: method[1:], code: "12"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := grpcStatus(t, listener.Addr().String(), tc.path); got != tc.code {
				t.Fatalf("grpc-status = %q, want %q", got, tc.code)
			}
		})
	}
}

func grpcStatus(t *testing.T, address, path string) string {
	t.Helper()
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Write([]byte(http2.ClientPreface)); err != nil {
		t.Fatal(err)
	}
	framer := http2.NewFramer(conn, conn)
	framer.ReadMetaHeaders = hpack.NewDecoder(4096, nil)
	if err := framer.WriteSettings(); err != nil {
		t.Fatal(err)
	}
	var headers bytes.Buffer
	encoder := hpack.NewEncoder(&headers)
	for _, field := range []hpack.HeaderField{
		{Name: ":method", Value: "POST"},
		{Name: ":scheme", Value: "http"},
		{Name: ":path", Value: path},
		{Name: ":authority", Value: address},
		{Name: "content-type", Value: "application/grpc"},
		{Name: "te", Value: "trailers"},
	} {
		if err := encoder.WriteField(field); err != nil {
			t.Fatal(err)
		}
	}
	if err := framer.WriteHeaders(http2.HeadersFrameParam{StreamID: 1, BlockFragment: headers.Bytes(), EndHeaders: true}); err != nil {
		t.Fatal(err)
	}
	if err := framer.WriteData(1, true, []byte{0, 0, 0, 0, 0}); err != nil {
		t.Fatal(err)
	}
	for {
		frame, err := framer.ReadFrame()
		if err != nil {
			t.Fatal(err)
		}
		switch frame := frame.(type) {
		case *http2.SettingsFrame:
			if !frame.IsAck() {
				if err := framer.WriteSettingsAck(); err != nil {
					t.Fatal(err)
				}
			}
		case *http2.MetaHeadersFrame:
			for _, field := range frame.Fields {
				if field.Name == "grpc-status" {
					return field.Value
				}
			}
		}
	}
}
