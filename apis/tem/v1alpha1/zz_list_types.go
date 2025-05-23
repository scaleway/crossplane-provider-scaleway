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

type ListInitParameters struct {

	// The ID of the domain affected by the blocklist. Must be in the format {region}/{domain_id}.
	// The ID of the domain affected by the blocklist.
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/tem/v1alpha1.Domain
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Reference to a Domain in tem to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDRef *v1.Reference `json:"domainIdRef,omitempty" tf:"-"`

	// Selector for a Domain in tem to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDSelector *v1.Selector `json:"domainIdSelector,omitempty" tf:"-"`

	// The email address to block.
	// Email address to block.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reason for blocking the email address.
	// Reason for blocking the emails.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The region in which the blocklist is created. Defaults to the provider's region.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Type of the blocklist. Possible values are:
	// Type of the blocked list. (mailbox_full or mailbox_not_found)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ListObservation struct {

	// The ID of the domain affected by the blocklist. Must be in the format {region}/{domain_id}.
	// The ID of the domain affected by the blocklist.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The email address to block.
	// Email address to block.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the blocklist, in the format {region}/{id}.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reason for blocking the email address.
	// Reason for blocking the emails.
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The region in which the blocklist is created. Defaults to the provider's region.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Type of the blocklist. Possible values are:
	// Type of the blocked list. (mailbox_full or mailbox_not_found)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ListParameters struct {

	// The ID of the domain affected by the blocklist. Must be in the format {region}/{domain_id}.
	// The ID of the domain affected by the blocklist.
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/tem/v1alpha1.Domain
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// Reference to a Domain in tem to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDRef *v1.Reference `json:"domainIdRef,omitempty" tf:"-"`

	// Selector for a Domain in tem to populate domainId.
	// +kubebuilder:validation:Optional
	DomainIDSelector *v1.Selector `json:"domainIdSelector,omitempty" tf:"-"`

	// The email address to block.
	// Email address to block.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the project this blocklist belongs to. Defaults to the provider's project ID.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reason for blocking the email address.
	// Reason for blocking the emails.
	// +kubebuilder:validation:Optional
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The region in which the blocklist is created. Defaults to the provider's region.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Type of the blocklist. Possible values are:
	// Type of the blocked list. (mailbox_full or mailbox_not_found)
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ListSpec defines the desired state of List
type ListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ListParameters `json:"forProvider"`
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
	InitProvider ListInitParameters `json:"initProvider,omitempty"`
}

// ListStatus defines the observed state of List.
type ListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// List is the Schema for the Lists API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type List struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ListSpec   `json:"spec"`
	Status ListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ListList contains a list of Lists
type ListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []List `json:"items"`
}

// Repository type metadata.
var (
	List_Kind             = "List"
	List_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: List_Kind}.String()
	List_KindAPIVersion   = List_Kind + "." + CRDGroupVersion.String()
	List_GroupVersionKind = CRDGroupVersion.WithKind(List_Kind)
)

func init() {
	SchemeBuilder.Register(&List{}, &ListList{})
}
