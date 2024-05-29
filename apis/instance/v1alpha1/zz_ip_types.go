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

type IPInitParameters struct {

	// (Defaults to provider project_id) The ID of the project the IP is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The tags associated with the IP.
	// The tags associated with the ip
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to nat) The type of the IP (nat, routed_ipv4, routed_ipv6), more information in the documentation
	// The type of instance IP
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the IP should be reserved.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type IPObservation struct {

	// The IP address.
	// The IP address
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the IP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization ID the IP is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The IP Prefix.
	// The IP prefix
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (Defaults to provider project_id) The ID of the project the IP is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The reverse dns attached to this IP
	// The reverse DNS for this IP
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	// The ID of the IP.
	// The server associated with this IP
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The tags associated with the IP.
	// The tags associated with the ip
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to nat) The type of the IP (nat, routed_ipv4, routed_ipv6), more information in the documentation
	// The type of instance IP
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the IP should be reserved.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type IPParameters struct {

	// (Defaults to provider project_id) The ID of the project the IP is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The tags associated with the IP.
	// The tags associated with the ip
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to nat) The type of the IP (nat, routed_ipv4, routed_ipv6), more information in the documentation
	// The type of instance IP
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the IP should be reserved.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// IPSpec defines the desired state of IP
type IPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPParameters `json:"forProvider"`
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
	InitProvider IPInitParameters `json:"initProvider,omitempty"`
}

// IPStatus defines the observed state of IP.
type IPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IP is the Schema for the IPs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type IP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPSpec   `json:"spec"`
	Status            IPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPList contains a list of IPs
type IPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IP `json:"items"`
}

// Repository type metadata.
var (
	IP_Kind             = "IP"
	IP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IP_Kind}.String()
	IP_KindAPIVersion   = IP_Kind + "." + CRDGroupVersion.String()
	IP_GroupVersionKind = CRDGroupVersion.WithKind(IP_Kind)
)

func init() {
	SchemeBuilder.Register(&IP{}, &IPList{})
}
