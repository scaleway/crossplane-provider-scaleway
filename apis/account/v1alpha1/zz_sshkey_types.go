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

type SSHKeyObservation struct {

	// The ID of the SSH key (UUID format).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization ID the SSH key is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type SSHKeyParameters struct {

	// The name of the SSH key.
	// The name of the SSH key
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the SSH key is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The public SSH key to be added.
	// The public SSH key
	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`
}

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHKeyParameters `json:"forProvider"`
}

// SSHKeyStatus defines the observed state of SSHKey.
type SSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKey is the Schema for the SSHKeys API. Manages Scaleway user SSH keys.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SSHKeySpec   `json:"spec"`
	Status            SSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKeyList contains a list of SSHKeys
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}

// Repository type metadata.
var (
	SSHKey_Kind             = "SSHKey"
	SSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHKey_Kind}.String()
	SSHKey_KindAPIVersion   = SSHKey_Kind + "." + CRDGroupVersion.String()
	SSHKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
