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

type NetworkInitParameters struct {

	// The hub ID to which the Network will be attached to.
	// The ID of the hub on which this network will be created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/iot/v1alpha1.Hub
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// Reference to a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDRef *v1.Reference `json:"hubIdRef,omitempty" tf:"-"`

	// Selector for a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDSelector *v1.Selector `json:"hubIdSelector,omitempty" tf:"-"`

	// The name of the IoT Network you want to create (e.g. my-net).
	// The name of the network
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Network is attached to.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The prefix that will be prepended to all topics for this Network.
	// The prefix that will be prepended to all topics for this Network
	TopicPrefix *string `json:"topicPrefix,omitempty" tf:"topic_prefix,omitempty"`

	// The network type to create (e.g. sigfox).
	// The type of the network
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkObservation struct {

	// The date and time the Network was created.
	// The date and time of the creation of the network
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The endpoint to use when interacting with the network.
	// The endpoint to use when interacting with the network
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The hub ID to which the Network will be attached to.
	// The ID of the hub on which this network will be created
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// The ID of the Network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoT Network you want to create (e.g. my-net).
	// The name of the network
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Network is attached to.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The prefix that will be prepended to all topics for this Network.
	// The prefix that will be prepended to all topics for this Network
	TopicPrefix *string `json:"topicPrefix,omitempty" tf:"topic_prefix,omitempty"`

	// The network type to create (e.g. sigfox).
	// The type of the network
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkParameters struct {

	// The hub ID to which the Network will be attached to.
	// The ID of the hub on which this network will be created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/iot/v1alpha1.Hub
	// +kubebuilder:validation:Optional
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// Reference to a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDRef *v1.Reference `json:"hubIdRef,omitempty" tf:"-"`

	// Selector for a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDSelector *v1.Selector `json:"hubIdSelector,omitempty" tf:"-"`

	// The name of the IoT Network you want to create (e.g. my-net).
	// The name of the network
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Network is attached to.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The prefix that will be prepended to all topics for this Network.
	// The prefix that will be prepended to all topics for this Network
	// +kubebuilder:validation:Optional
	TopicPrefix *string `json:"topicPrefix,omitempty" tf:"topic_prefix,omitempty"`

	// The network type to create (e.g. sigfox).
	// The type of the network
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
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
	InitProvider NetworkInitParameters `json:"initProvider,omitempty"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Network is the Schema for the Networks API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   NetworkSpec   `json:"spec"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
