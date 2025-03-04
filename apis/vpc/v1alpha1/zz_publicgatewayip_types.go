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

type PublicGatewayIPInitParameters struct {

	// (Defaults to provider project_id) The ID of the Project the Public Gateway IP is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The reverse domain name for the IP address
	// reverse domain name for the IP address
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	// The tags associated with the Public Gateway IP.
	// The tags associated with public gateway IP
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the Public Gateway IP should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PublicGatewayIPObservation struct {

	// The IP address itself.
	// the IP itself
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The date and time of the creation of the Public Gateway IP.
	// The date and time of the creation of the public gateway IP
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the Public Gateway IP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Organization ID the Public Gateway IP is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to provider project_id) The ID of the Project the Public Gateway IP is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The reverse domain name for the IP address
	// reverse domain name for the IP address
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	// The tags associated with the Public Gateway IP.
	// The tags associated with public gateway IP
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The date and time of the last update of the Public Gateway IP.
	// The date and time of the last update of the public gateway IP
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// (Defaults to provider zone) The zone in which the Public Gateway IP should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PublicGatewayIPParameters struct {

	// (Defaults to provider project_id) The ID of the Project the Public Gateway IP is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The reverse domain name for the IP address
	// reverse domain name for the IP address
	// +kubebuilder:validation:Optional
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	// The tags associated with the Public Gateway IP.
	// The tags associated with public gateway IP
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the Public Gateway IP should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// PublicGatewayIPSpec defines the desired state of PublicGatewayIP
type PublicGatewayIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicGatewayIPParameters `json:"forProvider"`
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
	InitProvider PublicGatewayIPInitParameters `json:"initProvider,omitempty"`
}

// PublicGatewayIPStatus defines the observed state of PublicGatewayIP.
type PublicGatewayIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicGatewayIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PublicGatewayIP is the Schema for the PublicGatewayIPs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type PublicGatewayIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicGatewayIPSpec   `json:"spec"`
	Status            PublicGatewayIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicGatewayIPList contains a list of PublicGatewayIPs
type PublicGatewayIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicGatewayIP `json:"items"`
}

// Repository type metadata.
var (
	PublicGatewayIP_Kind             = "PublicGatewayIP"
	PublicGatewayIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicGatewayIP_Kind}.String()
	PublicGatewayIP_KindAPIVersion   = PublicGatewayIP_Kind + "." + CRDGroupVersion.String()
	PublicGatewayIP_GroupVersionKind = CRDGroupVersion.WithKind(PublicGatewayIP_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicGatewayIP{}, &PublicGatewayIPList{})
}
