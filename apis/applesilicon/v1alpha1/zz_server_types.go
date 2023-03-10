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

type ServerObservation struct {

	// The date and time of the creation of the Apple Silicon server.
	// The date and time of the creation of the server
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The minimal date and time on which you can delete this server due to Apple licence
	DeletableAt *string `json:"deletableAt,omitempty" tf:"deletable_at,omitempty"`

	// The ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4 address of the server (IPv4 address).
	// IPv4 address of the server
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The organization ID the server is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The state of the server. Check the possible values on
	// our sdk.
	// The state of the server
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time of the last update of the Apple Silicon server.
	// The date and time of the last update of the server
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// URL of the VNC.
	// VNC url use to connect remotely to the desktop GUI
	VncURL *string `json:"vncUrl,omitempty" tf:"vnc_url,omitempty"`
}

type ServerParameters struct {

	// The name of the server.
	// Name of the server
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the server is
	// associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The commercial type of the server. You find all the available types on
	// the pricing page. Updates to this field will recreate a new
	// resource.
	// Type of the server
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which
	// the server should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Manages Scaleway Apple silicon M1 as-a-Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
