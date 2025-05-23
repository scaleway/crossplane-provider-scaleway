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

type FunctionInitParameters struct {

	// Define whether the function should be deployed. Your function will be redeployed if you update the source zip file
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the function.
	// The description of the function
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The environment variables of the function.
	// The environment variables of the function
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Allows both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// Handler of the function, depends on the runtime. Refer to the dedicated documentation for the list of supported runtimes.
	// Handler of the function. Depends on the runtime https://developers.scaleway.com/en/products/functions/api/#create-a-function
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured max_scale value.
	// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory resources in MB to allocate to each function. Defaults to 256 MB.
	// Memory limit in MB for your function, defaults to 256MB
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a min_scale greater than 0 will cause your function to run constantly.
	// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function to run all the time.
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the function name.
	// The name of the function
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Functions namespace ID of the function.
	// The namespace ID associated with this function
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/function/v1alpha1.FunctionNamespace
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a FunctionNamespace in function to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a FunctionNamespace in function to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// The privacy type defines the way to authenticate to your function. Please check our dedicated section.
	// Privacy of the function. Can be either `private` or `public`
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// (Defaults to provider project_id) The ID of the project the functions namespace is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the namespace should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Runtime of the function. Runtimes can be fetched using specific route
	// Runtime of the function
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Execution environment of the function.
	// Execution environment of the function.
	Sandbox *string `json:"sandbox,omitempty" tf:"sandbox,omitempty"`

	SecretEnvironmentVariables map[string]*string `json:"secretEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
	// Holds the max duration (in seconds) the function is allowed for responding to a request
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Path to the zip file containing your function sources to upload.
	// Location of the zip file to upload containing your function sources
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`

	// The hash of your source zip file, changing it will redeploy the function. Can be any string, changing it will simply trigger a state change.
	// The hash of your source zip file, changing it will re-apply function. Can be any string
	ZipHash *string `json:"zipHash,omitempty" tf:"zip_hash,omitempty"`
}

type FunctionObservation struct {

	// The CPU limit in mVCPU for your function.
	// CPU limit in mCPU for your function
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Define whether the function should be deployed. Your function will be redeployed if you update the source zip file
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the function.
	// The description of the function
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The native domain name of the function.
	// The native function domain name.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The environment variables of the function.
	// The environment variables of the function
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Allows both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// Handler of the function, depends on the runtime. Refer to the dedicated documentation for the list of supported runtimes.
	// Handler of the function. Depends on the runtime https://developers.scaleway.com/en/products/functions/api/#create-a-function
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// The unique identifier of the function.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured max_scale value.
	// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory resources in MB to allocate to each function. Defaults to 256 MB.
	// Memory limit in MB for your function, defaults to 256MB
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a min_scale greater than 0 will cause your function to run constantly.
	// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function to run all the time.
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the function name.
	// The name of the function
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Functions namespace ID of the function.
	// The namespace ID associated with this function
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// The organization ID the function is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The privacy type defines the way to authenticate to your function. Please check our dedicated section.
	// Privacy of the function. Can be either `private` or `public`
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// (Defaults to provider project_id) The ID of the project the functions namespace is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the namespace should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Runtime of the function. Runtimes can be fetched using specific route
	// Runtime of the function
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Execution environment of the function.
	// Execution environment of the function.
	Sandbox *string `json:"sandbox,omitempty" tf:"sandbox,omitempty"`

	// The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
	// Holds the max duration (in seconds) the function is allowed for responding to a request
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Path to the zip file containing your function sources to upload.
	// Location of the zip file to upload containing your function sources
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`

	// The hash of your source zip file, changing it will redeploy the function. Can be any string, changing it will simply trigger a state change.
	// The hash of your source zip file, changing it will re-apply function. Can be any string
	ZipHash *string `json:"zipHash,omitempty" tf:"zip_hash,omitempty"`
}

type FunctionParameters struct {

	// Define whether the function should be deployed. Your function will be redeployed if you update the source zip file
	// +kubebuilder:validation:Optional
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the function.
	// The description of the function
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The environment variables of the function.
	// The environment variables of the function
	// +kubebuilder:validation:Optional
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Allows both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	// +kubebuilder:validation:Optional
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// Handler of the function, depends on the runtime. Refer to the dedicated documentation for the list of supported runtimes.
	// Handler of the function. Depends on the runtime https://developers.scaleway.com/en/products/functions/api/#create-a-function
	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// The maximum number of instances this function can scale to. Default to 20. Your function will scale automatically based on the incoming workload, but will never exceed the configured max_scale value.
	// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
	// +kubebuilder:validation:Optional
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory resources in MB to allocate to each function. Defaults to 256 MB.
	// Memory limit in MB for your function, defaults to 256MB
	// +kubebuilder:validation:Optional
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum number of function instances running continuously. Defaults to 0. Functions are billed when executed, and using a min_scale greater than 0 will cause your function to run constantly.
	// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function to run all the time.
	// +kubebuilder:validation:Optional
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the function name.
	// The name of the function
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Functions namespace ID of the function.
	// The namespace ID associated with this function
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/function/v1alpha1.FunctionNamespace
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a FunctionNamespace in function to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a FunctionNamespace in function to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// The privacy type defines the way to authenticate to your function. Please check our dedicated section.
	// Privacy of the function. Can be either `private` or `public`
	// +kubebuilder:validation:Optional
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// (Defaults to provider project_id) The ID of the project the functions namespace is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the namespace should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Runtime of the function. Runtimes can be fetched using specific route
	// Runtime of the function
	// +kubebuilder:validation:Optional
	Runtime *string `json:"runtime,omitempty" tf:"runtime,omitempty"`

	// Execution environment of the function.
	// Execution environment of the function.
	// +kubebuilder:validation:Optional
	Sandbox *string `json:"sandbox,omitempty" tf:"sandbox,omitempty"`

	// The secret environment variables of the function.
	// The secret environment variables to be injected into your function at runtime.
	// +kubebuilder:validation:Optional
	SecretEnvironmentVariablesSecretRef *v1.SecretReference `json:"secretEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// The maximum amount of time your function can spend processing a request before being stopped. Defaults to 300s.
	// Holds the max duration (in seconds) the function is allowed for responding to a request
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Path to the zip file containing your function sources to upload.
	// Location of the zip file to upload containing your function sources
	// +kubebuilder:validation:Optional
	ZipFile *string `json:"zipFile,omitempty" tf:"zip_file,omitempty"`

	// The hash of your source zip file, changing it will redeploy the function. Can be any string, changing it will simply trigger a state change.
	// The hash of your source zip file, changing it will re-apply function. Can be any string
	// +kubebuilder:validation:Optional
	ZipHash *string `json:"zipHash,omitempty" tf:"zip_hash,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
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
	InitProvider FunctionInitParameters `json:"initProvider,omitempty"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Function is the Schema for the Functions API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.handler) || (has(self.initProvider) && has(self.initProvider.handler))",message="spec.forProvider.handler is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privacy) || (has(self.initProvider) && has(self.initProvider.privacy))",message="spec.forProvider.privacy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.runtime) || (has(self.initProvider) && has(self.initProvider.runtime))",message="spec.forProvider.runtime is a required parameter"
	Spec   FunctionSpec   `json:"spec"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
