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

type VPCInitParameters struct {

	// Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
	// Enable routing between Private Networks in the VPC
	EnableRouting *bool `json:"enableRouting,omitempty" tf:"enable_routing,omitempty"`

	// The name for the VPC. If not provided it will be randomly generated.
	// The name of the VPC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the Project the VPC is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region of the VPC.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags to associate with the VPC.
	// The tags associated with the VPC
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VPCObservation struct {

	// Date and time of VPC's creation (RFC 3339 format).
	// The date and time of the creation of the private network
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
	// Enable routing between Private Networks in the VPC
	EnableRouting *bool `json:"enableRouting,omitempty" tf:"enable_routing,omitempty"`

	// The ID of the VPC.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Defines whether the VPC is the default one for its Project.
	// Defines whether the VPC is the default one for its Project
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name for the VPC. If not provided it will be randomly generated.
	// The name of the VPC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Organization ID the VPC is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to provider project_id) The ID of the Project the VPC is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region of the VPC.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags to associate with the VPC.
	// The tags associated with the VPC
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Date and time of VPC's last update (RFC 3339 format).
	// The date and time of the last update of the private network
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type VPCParameters struct {

	// Enable routing between Private Networks in the VPC. Note that you will not be able to deactivate it afterwards.
	// Enable routing between Private Networks in the VPC
	// +kubebuilder:validation:Optional
	EnableRouting *bool `json:"enableRouting,omitempty" tf:"enable_routing,omitempty"`

	// The name for the VPC. If not provided it will be randomly generated.
	// The name of the VPC
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the Project the VPC is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region of the VPC.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags to associate with the VPC.
	// The tags associated with the VPC
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
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
	InitProvider VPCInitParameters `json:"initProvider,omitempty"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPC is the Schema for the VPCs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec"`
	Status            VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
