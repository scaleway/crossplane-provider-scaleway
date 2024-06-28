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

type BucketInitParameters struct {

	// (Deprecated) The canned ACL you want to apply to the bucket.
	// ACL of the bucket: either 'private', 'public-read', 'public-read-write' or 'authenticated-read'.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleInitParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and not recoverable
	// Delete objects in bucket
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRule []LifecycleRuleInitParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The name of the bucket.
	// The name of the bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable object lock
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which the bucket should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with this bucket
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	// Allow multiple versions of an object in the same bucket
	Versioning []VersioningInitParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketObservation struct {

	// (Deprecated) The canned ACL you want to apply to the bucket.
	// ACL of the bucket: either 'private', 'public-read', 'public-read-write' or 'authenticated-read'.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The endpoint URL of the bucket
	// API URL of the bucket
	APIEndpoint *string `json:"apiEndpoint,omitempty" tf:"api_endpoint,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRule []CorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// The endpoint URL of the bucket
	// Endpoint of the bucket
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and not recoverable
	// Delete objects in bucket
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The name of the bucket.
	// The name of the bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable object lock
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which the bucket should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with this bucket
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	// Allow multiple versions of an object in the same bucket
	Versioning []VersioningObservation `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type BucketParameters struct {

	// (Deprecated) The canned ACL you want to apply to the bucket.
	// ACL of the bucket: either 'private', 'public-read', 'public-read-write' or 'authenticated-read'.
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// A rule of Cross-Origin Resource Sharing (documented below).
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and not recoverable
	// Delete objects in bucket
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The name of the bucket.
	// The name of the bucket
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable object lock
	// +kubebuilder:validation:Optional
	ObjectLockEnabled *bool `json:"objectLockEnabled,omitempty" tf:"object_lock_enabled,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which the bucket should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with this bucket
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A state of versioning (documented below)
	// Allow multiple versions of an object in the same bucket
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`
}

type CorsRuleInitParameters struct {

	// Specifies which headers are allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleObservation struct {

	// Specifies which headers are allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleParameters struct {

	// Specifies which headers are allowed.
	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// Specifies which origins are allowed.
	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Specifies expose header in the response.
	// +kubebuilder:validation:Optional
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// Specifies time in seconds that browser can cache the response for a preflight request.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type ExpirationInitParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type ExpirationObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type ExpirationParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days" tf:"days,omitempty"`
}

type LifecycleRuleInitParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Specifies if the configuration rule is Enabled or Disabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	// Specifies a period in the object's expire
	Expiration []ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// Unique identifier for the rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// The prefix identifying one or more objects to which the rule applies
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with the bucket lifecycle
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a period in the object's transitions (documented below).
	// Define when objects transition to another storage class
	Transition []TransitionInitParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LifecycleRuleObservation struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Specifies if the configuration rule is Enabled or Disabled
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	// Specifies a period in the object's expire
	Expiration []ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// Unique identifier for the rule
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// The prefix identifying one or more objects to which the rule applies
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with the bucket lifecycle
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a period in the object's transitions (documented below).
	// Define when objects transition to another storage class
	Transition []TransitionObservation `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LifecycleRuleParameters struct {

	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Specifies if the configuration rule is Enabled or Disabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Specifies a period in the object's expire (documented below).
	// Specifies a period in the object's expire
	// +kubebuilder:validation:Optional
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	// Unique identifier for the rule
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object key prefix identifying one or more objects to which the rule applies.
	// The prefix identifying one or more objects to which the rule applies
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// A list of tags (key / value) for the bucket.
	// The tags associated with the bucket lifecycle
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies a period in the object's transitions (documented below).
	// Define when objects transition to another storage class
	// +kubebuilder:validation:Optional
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type TransitionInitParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Scaleway storage class STANDARD, GLACIER, ONEZONE_IA  to which you want the object to transition.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type TransitionObservation struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Scaleway storage class STANDARD, GLACIER, ONEZONE_IA  to which you want the object to transition.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type TransitionParameters struct {

	// Specifies the number of days after object creation when the specific rule action takes effect.
	// Specifies the number of days after object creation when the specific rule action takes effect
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Specifies the Scaleway storage class STANDARD, GLACIER, ONEZONE_IA  to which you want the object to transition.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type VersioningInitParameters struct {

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type VersioningObservation struct {

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type VersioningParameters struct {

	// The element value can be either Enabled or Disabled. If a rule is disabled, Scaleway S3 doesn't perform any of the actions defined in the rule.
	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Bucket is the Schema for the Buckets API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
