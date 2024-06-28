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

type UserInitParameters struct {

	// The email of the IAM user.
	// The description of the iam user
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (Defaults to provider organization_id) The ID of the organization the user is associated with.
	// ID of organization the resource is associated to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type UserObservation struct {

	// The ID of the account root user associated with the user.
	// The ID of the account root user associated with the iam user.
	AccountRootUserID *string `json:"accountRootUserId,omitempty" tf:"account_root_user_id,omitempty"`

	// The date and time of the creation of the iam user.
	// The date and time of the creation of the iam user
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether the iam user is deletable.
	// Whether or not the iam user is editable
	Deletable *bool `json:"deletable,omitempty" tf:"deletable,omitempty"`

	// The email of the IAM user.
	// The description of the iam user
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the user (UUID format).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date of the last login.
	// The date and time of last login of the iam user
	LastLoginAt *string `json:"lastLoginAt,omitempty" tf:"last_login_at,omitempty"`

	// Whether the MFA is enabled.
	// Whether or not the MFA is enabled
	Mfa *bool `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// (Defaults to provider organization_id) The ID of the organization the user is associated with.
	// ID of organization the resource is associated to.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The status of user invitation. Check the possible values in the api doc.
	// The status of user invitation.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The type of user. Check the possible values in the api doc.
	// The type of the iam user
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The date and time of the last update of the iam user.
	// The date and time of the last update of the iam user
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type UserParameters struct {

	// The email of the IAM user.
	// The description of the iam user
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (Defaults to provider organization_id) The ID of the organization the user is associated with.
	// ID of organization the resource is associated to.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
