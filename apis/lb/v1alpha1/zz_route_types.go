/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteInitParameters struct {

	// The ID of the backend the route is associated with.
	// The backend ID destination of redirection
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Backend
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// Reference to a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDRef *v1.Reference `json:"backendIdRef,omitempty" tf:"-"`

	// Selector for a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDSelector *v1.Selector `json:"backendIdSelector,omitempty" tf:"-"`

	// The ID of the frontend the route is associated with.
	// The frontend ID origin of redirection
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Frontend
	FrontendID *string `json:"frontendId,omitempty" tf:"frontend_id,omitempty"`

	// Reference to a Frontend in lb to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDRef *v1.Reference `json:"frontendIdRef,omitempty" tf:"-"`

	// Selector for a Frontend in lb to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDSelector *v1.Selector `json:"frontendIdSelector,omitempty" tf:"-"`

	// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
	// Only one of match_sni and match_host_header should be specified.
	// Specifies the host of the server to which the request is being sent
	MatchHostHeader *string `json:"matchHostHeader,omitempty" tf:"match_host_header,omitempty"`

	// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
	// Only one of match_sni and match_host_header should be specified.
	// Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer
	MatchSni *string `json:"matchSni,omitempty" tf:"match_sni,omitempty"`
}

type RouteObservation struct {

	// The ID of the backend the route is associated with.
	// The backend ID destination of redirection
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// The date on which the route was created.
	// The date at which the route was created (RFC 3339 format)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the frontend the route is associated with.
	// The frontend ID origin of redirection
	FrontendID *string `json:"frontendId,omitempty" tf:"frontend_id,omitempty"`

	// The ID of the route
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
	// Only one of match_sni and match_host_header should be specified.
	// Specifies the host of the server to which the request is being sent
	MatchHostHeader *string `json:"matchHostHeader,omitempty" tf:"match_host_header,omitempty"`

	// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
	// Only one of match_sni and match_host_header should be specified.
	// Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer
	MatchSni *string `json:"matchSni,omitempty" tf:"match_sni,omitempty"`

	// The date on which the route was last updated.
	// The date at which the route was last updated (RFC 3339 format)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type RouteParameters struct {

	// The ID of the backend the route is associated with.
	// The backend ID destination of redirection
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Backend
	// +kubebuilder:validation:Optional
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// Reference to a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDRef *v1.Reference `json:"backendIdRef,omitempty" tf:"-"`

	// Selector for a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDSelector *v1.Selector `json:"backendIdSelector,omitempty" tf:"-"`

	// The ID of the frontend the route is associated with.
	// The frontend ID origin of redirection
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Frontend
	// +kubebuilder:validation:Optional
	FrontendID *string `json:"frontendId,omitempty" tf:"frontend_id,omitempty"`

	// Reference to a Frontend in lb to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDRef *v1.Reference `json:"frontendIdRef,omitempty" tf:"-"`

	// Selector for a Frontend in lb to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDSelector *v1.Selector `json:"frontendIdSelector,omitempty" tf:"-"`

	// The HTTP host header to match. Value to match in the HTTP Host request header from an incoming connection.
	// Only one of match_sni and match_host_header should be specified.
	// Specifies the host of the server to which the request is being sent
	// +kubebuilder:validation:Optional
	MatchHostHeader *string `json:"matchHostHeader,omitempty" tf:"match_host_header,omitempty"`

	// The Server Name Indication (SNI) value to match. Value to match in the Server Name Indication TLS extension (SNI) field from an incoming connection made via an SSL/TLS transport layer.
	// Only one of match_sni and match_host_header should be specified.
	// Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer
	// +kubebuilder:validation:Optional
	MatchSni *string `json:"matchSni,omitempty" tf:"match_sni,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Route is the Schema for the Routes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
