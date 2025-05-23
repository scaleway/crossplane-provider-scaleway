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

type DeploymentInitParameters struct {

	// Some models (e.g Meta Llama) require end-user license agreements. Set true to accept.
	// Whether or not the deployment is accepting eula
	AcceptEula *bool `json:"acceptEula,omitempty" tf:"accept_eula,omitempty"`

	// The maximum size of the pool.
	// The maximum size of the pool
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the pool.
	// The minimum size of the pool
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The model id used for the deployment.
	// The model id used for the deployment
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/inference/v1alpha1.Model
	ModelID *string `json:"modelId,omitempty" tf:"model_id,omitempty"`

	// Reference to a Model in inference to populate modelId.
	// +kubebuilder:validation:Optional
	ModelIDRef *v1.Reference `json:"modelIdRef,omitempty" tf:"-"`

	// Selector for a Model in inference to populate modelId.
	// +kubebuilder:validation:Optional
	ModelIDSelector *v1.Selector `json:"modelIdSelector,omitempty" tf:"-"`

	// The deployment name.
	// The deployment name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (scw inference node-type list)
	// The node type to use for the deployment
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Configuration of the deployment's private endpoint.
	// List of endpoints
	PrivateEndpoint []PrivateEndpointInitParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// (Defaults to provider project_id) The ID of the project the deployment is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Configuration of the deployment's public endpoint.
	// Public endpoints
	PublicEndpoint []PublicEndpointInitParameters `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`

	// The number of bits each model parameter should be quantized to
	Quantization *float64 `json:"quantization,omitempty" tf:"quantization,omitempty"`

	// (Defaults to provider region) The region in which the deployment is created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags associated with the deployment.
	// The tags associated with the deployment
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DeploymentObservation struct {

	// Some models (e.g Meta Llama) require end-user license agreements. Set true to accept.
	// Whether or not the deployment is accepting eula
	AcceptEula *bool `json:"acceptEula,omitempty" tf:"accept_eula,omitempty"`

	// The date and time of the creation of the deployment.
	// The date and time of the creation of the deployment
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the deployment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum size of the pool.
	// The maximum size of the pool
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the pool.
	// The minimum size of the pool
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The model id used for the deployment.
	// The model id used for the deployment
	ModelID *string `json:"modelId,omitempty" tf:"model_id,omitempty"`

	// The model name used for the deployment. Model names can be found in Console or using Scaleway's CLI (scw inference model list)
	// The model name to use for the deployment
	ModelName *string `json:"modelName,omitempty" tf:"model_name,omitempty"`

	// The deployment name.
	// The deployment name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (scw inference node-type list)
	// The node type to use for the deployment
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Configuration of the deployment's private endpoint.
	// List of endpoints
	PrivateEndpoint []PrivateEndpointObservation `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// (Defaults to provider project_id) The ID of the project the deployment is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Configuration of the deployment's public endpoint.
	// Public endpoints
	PublicEndpoint []PublicEndpointObservation `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`

	// The number of bits each model parameter should be quantized to
	Quantization *float64 `json:"quantization,omitempty" tf:"quantization,omitempty"`

	// (Defaults to provider region) The region in which the deployment is created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The size of the pool.
	// The size of the pool
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The status of the deployment.
	// The status of the deployment
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The tags associated with the deployment.
	// The tags associated with the deployment
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The date and time of the last update of the deployment.
	// The date and time of the last update of the deployment
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type DeploymentParameters struct {

	// Some models (e.g Meta Llama) require end-user license agreements. Set true to accept.
	// Whether or not the deployment is accepting eula
	// +kubebuilder:validation:Optional
	AcceptEula *bool `json:"acceptEula,omitempty" tf:"accept_eula,omitempty"`

	// The maximum size of the pool.
	// The maximum size of the pool
	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// The minimum size of the pool.
	// The minimum size of the pool
	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// The model id used for the deployment.
	// The model id used for the deployment
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/inference/v1alpha1.Model
	// +kubebuilder:validation:Optional
	ModelID *string `json:"modelId,omitempty" tf:"model_id,omitempty"`

	// Reference to a Model in inference to populate modelId.
	// +kubebuilder:validation:Optional
	ModelIDRef *v1.Reference `json:"modelIdRef,omitempty" tf:"-"`

	// Selector for a Model in inference to populate modelId.
	// +kubebuilder:validation:Optional
	ModelIDSelector *v1.Selector `json:"modelIdSelector,omitempty" tf:"-"`

	// The deployment name.
	// The deployment name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (scw inference node-type list)
	// The node type to use for the deployment
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Configuration of the deployment's private endpoint.
	// List of endpoints
	// +kubebuilder:validation:Optional
	PrivateEndpoint []PrivateEndpointParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// (Defaults to provider project_id) The ID of the project the deployment is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Configuration of the deployment's public endpoint.
	// Public endpoints
	// +kubebuilder:validation:Optional
	PublicEndpoint []PublicEndpointParameters `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`

	// The number of bits each model parameter should be quantized to
	// +kubebuilder:validation:Optional
	Quantization *float64 `json:"quantization,omitempty" tf:"quantization,omitempty"`

	// (Defaults to provider region) The region in which the deployment is created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The tags associated with the deployment.
	// The tags associated with the deployment
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateEndpointInitParameters struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// The ID of the private network to use.
	// The id of the private network
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`
}

type PrivateEndpointObservation struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// The ID of the deployment.
	// The id of the private endpoint
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the private network to use.
	// The id of the private network
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// The URL of the endpoint.
	// The URL of the endpoint.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PrivateEndpointParameters struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	// +kubebuilder:validation:Optional
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// The ID of the private network to use.
	// The id of the private network
	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`
}

type PublicEndpointInitParameters struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// Enable or disable public endpoint.
	// Enable or disable public endpoint
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type PublicEndpointObservation struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// The ID of the deployment.
	// The id of the public endpoint
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Enable or disable public endpoint.
	// Enable or disable public endpoint
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The URL of the endpoint.
	// The URL of the endpoint.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PublicEndpointParameters struct {

	// Disable the authentication on the endpoint.
	// Disable the authentication on the endpoint.
	// +kubebuilder:validation:Optional
	DisableAuth *bool `json:"disableAuth,omitempty" tf:"disable_auth,omitempty"`

	// Enable or disable public endpoint.
	// Enable or disable public endpoint
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
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
	InitProvider DeploymentInitParameters `json:"initProvider,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Deployment is the Schema for the Deployments API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeType) || (has(self.initProvider) && has(self.initProvider.nodeType))",message="spec.forProvider.nodeType is a required parameter"
	Spec   DeploymentSpec   `json:"spec"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
