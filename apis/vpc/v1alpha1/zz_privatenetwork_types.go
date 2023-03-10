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

type PrivateNetworkObservation struct {

	// The date and time of the creation of the private network
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the private network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization ID the private network is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The date and time of the last update of the private network
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type PrivateNetworkParameters struct {

	// The name of the private network. If not provided it will be randomly generated.
	// The name of the private network
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the private network is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The tags associated with the private network.
	// The tags associated with private network
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the private network should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// PrivateNetworkSpec defines the desired state of PrivateNetwork
type PrivateNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateNetworkParameters `json:"forProvider"`
}

// PrivateNetworkStatus defines the observed state of PrivateNetwork.
type PrivateNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateNetwork is the Schema for the PrivateNetworks API. Manages Scaleway VPC Private Networks.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type PrivateNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateNetworkSpec   `json:"spec"`
	Status            PrivateNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateNetworkList contains a list of PrivateNetworks
type PrivateNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateNetwork `json:"items"`
}

// Repository type metadata.
var (
	PrivateNetwork_Kind             = "PrivateNetwork"
	PrivateNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateNetwork_Kind}.String()
	PrivateNetwork_KindAPIVersion   = PrivateNetwork_Kind + "." + CRDGroupVersion.String()
	PrivateNetwork_GroupVersionKind = CRDGroupVersion.WithKind(PrivateNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateNetwork{}, &PrivateNetworkList{})
}
