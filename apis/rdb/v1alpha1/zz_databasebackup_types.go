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

type DatabaseBackupInitParameters struct {

	// Name of the database (e.g. my-database).
	// Name of the database of this backup.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Expiration date (Format ISO 8601).
	// Expiration date (Format ISO 8601). Cannot be removed.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the user is created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/rdb/v1alpha1.Instance
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Name of the database (e.g. my-database).
	// Name of the backup.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DatabaseBackupObservation struct {

	// Creation date (Format ISO 8601).
	// Creation date (Format ISO 8601).
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Name of the database (e.g. my-database).
	// Name of the database of this backup.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Expiration date (Format ISO 8601).
	// Expiration date (Format ISO 8601). Cannot be removed.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// The ID of the backup, which is of the form {region}/{id}, e.g. fr-par/11111111-1111-1111-1111-111111111111
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the user is created
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Name of the instance of the backup.
	// Name of the instance of the backup.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Name of the database (e.g. my-database).
	// Name of the backup.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Size of the backup (in bytes).
	// Size of the backup (in bytes).
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Updated date (Format ISO 8601).
	// Updated date (Format ISO 8601).
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type DatabaseBackupParameters struct {

	// Name of the database (e.g. my-database).
	// Name of the database of this backup.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Expiration date (Format ISO 8601).
	// Expiration date (Format ISO 8601). Cannot be removed.
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the user is created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/rdb/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Name of the database (e.g. my-database).
	// Name of the backup.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// DatabaseBackupSpec defines the desired state of DatabaseBackup
type DatabaseBackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseBackupParameters `json:"forProvider"`
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
	InitProvider DatabaseBackupInitParameters `json:"initProvider,omitempty"`
}

// DatabaseBackupStatus defines the observed state of DatabaseBackup.
type DatabaseBackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseBackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DatabaseBackup is the Schema for the DatabaseBackups API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type DatabaseBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseName) || (has(self.initProvider) && has(self.initProvider.databaseName))",message="spec.forProvider.databaseName is a required parameter"
	Spec   DatabaseBackupSpec   `json:"spec"`
	Status DatabaseBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseBackupList contains a list of DatabaseBackups
type DatabaseBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseBackup `json:"items"`
}

// Repository type metadata.
var (
	DatabaseBackup_Kind             = "DatabaseBackup"
	DatabaseBackup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseBackup_Kind}.String()
	DatabaseBackup_KindAPIVersion   = DatabaseBackup_Kind + "." + CRDGroupVersion.String()
	DatabaseBackup_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseBackup_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseBackup{}, &DatabaseBackupList{})
}
