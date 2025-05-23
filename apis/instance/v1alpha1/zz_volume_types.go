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

type VolumeInitParameters struct {

	// If set, the new volume will be created from this snapshot. Only one of size_in_gb and from_snapshot_id should be specified.
	// Create a volume based on a image
	FromSnapshotID *string `json:"fromSnapshotId,omitempty" tf:"from_snapshot_id,omitempty"`

	// If true, consider that this volume may have been migrated and no longer exists.
	MigrateToSbs *bool `json:"migrateToSbs,omitempty" tf:"migrate_to_sbs,omitempty"`

	// The name of the volume. If not provided it will be randomly generated.
	// The name of the volume
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the volume is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The size of the volume. Only one of size_in_gb and from_snapshot_id should be specified.
	// The size of the volume in gigabyte
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// A list of tags to apply to the volume.
	// The tags associated with the volume
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the volume. The possible values are: b_ssd (Block SSD), l_ssd (Local SSD), scratch (Local Scratch SSD).
	// The volume type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the volume should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type VolumeObservation struct {

	// If set, the new volume will be created from this snapshot. Only one of size_in_gb and from_snapshot_id should be specified.
	// Create a volume based on a image
	FromSnapshotID *string `json:"fromSnapshotId,omitempty" tf:"from_snapshot_id,omitempty"`

	// The ID of the volume.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If true, consider that this volume may have been migrated and no longer exists.
	MigrateToSbs *bool `json:"migrateToSbs,omitempty" tf:"migrate_to_sbs,omitempty"`

	// The name of the volume. If not provided it will be randomly generated.
	// The name of the volume
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the volume is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to provider project_id) The ID of the project the volume is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The id of the associated server.
	// The server associated with this volume
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The size of the volume. Only one of size_in_gb and from_snapshot_id should be specified.
	// The size of the volume in gigabyte
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// A list of tags to apply to the volume.
	// The tags associated with the volume
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the volume. The possible values are: b_ssd (Block SSD), l_ssd (Local SSD), scratch (Local Scratch SSD).
	// The volume type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the volume should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type VolumeParameters struct {

	// If set, the new volume will be created from this snapshot. Only one of size_in_gb and from_snapshot_id should be specified.
	// Create a volume based on a image
	// +kubebuilder:validation:Optional
	FromSnapshotID *string `json:"fromSnapshotId,omitempty" tf:"from_snapshot_id,omitempty"`

	// If true, consider that this volume may have been migrated and no longer exists.
	// +kubebuilder:validation:Optional
	MigrateToSbs *bool `json:"migrateToSbs,omitempty" tf:"migrate_to_sbs,omitempty"`

	// The name of the volume. If not provided it will be randomly generated.
	// The name of the volume
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the volume is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The size of the volume. Only one of size_in_gb and from_snapshot_id should be specified.
	// The size of the volume in gigabyte
	// +kubebuilder:validation:Optional
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// A list of tags to apply to the volume.
	// The tags associated with the volume
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the volume. The possible values are: b_ssd (Block SSD), l_ssd (Local SSD), scratch (Local Scratch SSD).
	// The volume type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone in which the volume should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// VolumeSpec defines the desired state of Volume
type VolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeParameters `json:"forProvider"`
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
	InitProvider VolumeInitParameters `json:"initProvider,omitempty"`
}

// VolumeStatus defines the observed state of Volume.
type VolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Volume is the Schema for the Volumes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Volume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   VolumeSpec   `json:"spec"`
	Status VolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeList contains a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Volume `json:"items"`
}

// Repository type metadata.
var (
	Volume_Kind             = "Volume"
	Volume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Volume_Kind}.String()
	Volume_KindAPIVersion   = Volume_Kind + "." + CRDGroupVersion.String()
	Volume_GroupVersionKind = CRDGroupVersion.WithKind(Volume_Kind)
)

func init() {
	SchemeBuilder.Register(&Volume{}, &VolumeList{})
}
