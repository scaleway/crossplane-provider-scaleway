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

type ContainerInitParameters struct {

	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	// The amount of vCPU computing resources to allocate to each container. Defaults to 70.
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Boolean controlling whether the container is on a production environment.
	// This allows you to control your production environment
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the container.
	// The container description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The environment variables of the container.
	// The environment variables to be injected into your container at runtime.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Allow both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	// The maximum the number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency *float64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum of number of instances this container can scale to. Default to 20.
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	// The memory computing resources in MB to allocate to each container. Defaults to 128.
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum of running container instances continuously. Defaults to 0.
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the container name.
	// The container name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The container namespace ID of the container.
	// The container namespace associated
	// +crossplane:generate:reference:type=ContainerNamespace
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a ContainerNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a ContainerNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// The port to expose the container. Defaults to 8080.
	// The port to expose the container. Defaults to 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The privacy type define the way to authenticate to your container. Please check our dedicated section.
	// The privacy type define the way to authenticate to your container
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// The communication protocol http1 or h2c. Defaults to http1.
	// The communication protocol http1 or h2c. Defaults to http1.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Defaults to provider region) The region in which the container was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The registry image address. e.g: "rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE".
	// The scaleway registry image address
	RegistryImage *string `json:"registryImage,omitempty" tf:"registry_image,omitempty"`

	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
	RegistrySha256 *string `json:"registrySha256,omitempty" tf:"registry_sha256,omitempty"`

	SecretEnvironmentVariables map[string]*string `json:"secretEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// The container status.
	// The container status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ContainerObservation struct {

	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	// The amount of vCPU computing resources to allocate to each container. Defaults to 70.
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// The cron status of the container.
	// The cron status
	CronStatus *string `json:"cronStatus,omitempty" tf:"cron_status,omitempty"`

	// Boolean controlling whether the container is on a production environment.
	// This allows you to control your production environment
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the container.
	// The container description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The native domain name of the container
	// The native container domain name.
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// The environment variables of the container.
	// The environment variables to be injected into your container at runtime.
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// The error message of the container.
	// The error description
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	// Allow both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// The container's ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	// The maximum the number of simultaneous requests your container can handle at the same time. Defaults to 50.
	MaxConcurrency *float64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum of number of instances this container can scale to. Default to 20.
	// The maximum of number of instances this container can scale to. Default to 20.
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	// The memory computing resources in MB to allocate to each container. Defaults to 128.
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum of running container instances continuously. Defaults to 0.
	// The minimum of running container instances continuously. Defaults to 0.
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the container name.
	// The container name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The container namespace ID of the container.
	// The container namespace associated
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// The port to expose the container. Defaults to 8080.
	// The port to expose the container. Defaults to 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The privacy type define the way to authenticate to your container. Please check our dedicated section.
	// The privacy type define the way to authenticate to your container
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// The communication protocol http1 or h2c. Defaults to http1.
	// The communication protocol http1 or h2c. Defaults to http1.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Defaults to provider region) The region in which the container was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The registry image address. e.g: "rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE".
	// The scaleway registry image address
	RegistryImage *string `json:"registryImage,omitempty" tf:"registry_image,omitempty"`

	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
	RegistrySha256 *string `json:"registrySha256,omitempty" tf:"registry_sha256,omitempty"`

	// The container status.
	// The container status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ContainerParameters struct {

	// The amount of vCPU computing resources to allocate to each container. Defaults to 140.
	// The amount of vCPU computing resources to allocate to each container. Defaults to 70.
	// +kubebuilder:validation:Optional
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Boolean controlling whether the container is on a production environment.
	// This allows you to control your production environment
	// +kubebuilder:validation:Optional
	Deploy *bool `json:"deploy,omitempty" tf:"deploy,omitempty"`

	// The description of the container.
	// The container description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The environment variables of the container.
	// The environment variables to be injected into your container at runtime.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" tf:"environment_variables,omitempty"`

	// Allow both HTTP and HTTPS (enabled) or redirect HTTP to HTTPS (redirected). Defaults to enabled.
	// HTTP traffic configuration
	// +kubebuilder:validation:Optional
	HTTPOption *string `json:"httpOption,omitempty" tf:"http_option,omitempty"`

	// The maximum number of simultaneous requests your container can handle at the same time. Defaults to 50.
	// The maximum the number of simultaneous requests your container can handle at the same time. Defaults to 50.
	// +kubebuilder:validation:Optional
	MaxConcurrency *float64 `json:"maxConcurrency,omitempty" tf:"max_concurrency,omitempty"`

	// The maximum of number of instances this container can scale to. Default to 20.
	// The maximum of number of instances this container can scale to. Default to 20.
	// +kubebuilder:validation:Optional
	MaxScale *float64 `json:"maxScale,omitempty" tf:"max_scale,omitempty"`

	// The memory computing resources in MB to allocate to each container. Defaults to 256.
	// The memory computing resources in MB to allocate to each container. Defaults to 128.
	// +kubebuilder:validation:Optional
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// The minimum of running container instances continuously. Defaults to 0.
	// The minimum of running container instances continuously. Defaults to 0.
	// +kubebuilder:validation:Optional
	MinScale *float64 `json:"minScale,omitempty" tf:"min_scale,omitempty"`

	// The unique name of the container name.
	// The container name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The container namespace ID of the container.
	// The container namespace associated
	// +crossplane:generate:reference:type=ContainerNamespace
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a ContainerNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a ContainerNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// The port to expose the container. Defaults to 8080.
	// The port to expose the container. Defaults to 8080
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The privacy type define the way to authenticate to your container. Please check our dedicated section.
	// The privacy type define the way to authenticate to your container
	// +kubebuilder:validation:Optional
	Privacy *string `json:"privacy,omitempty" tf:"privacy,omitempty"`

	// The communication protocol http1 or h2c. Defaults to http1.
	// The communication protocol http1 or h2c. Defaults to http1.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Defaults to provider region) The region in which the container was created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The registry image address. e.g: "rg.fr-par.scw.cloud/$NAMESPACE/$IMAGE".
	// The scaleway registry image address
	// +kubebuilder:validation:Optional
	RegistryImage *string `json:"registryImage,omitempty" tf:"registry_image,omitempty"`

	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string.
	// The sha256 of your source registry image, changing it will re-apply the deployment. Can be any string
	// +kubebuilder:validation:Optional
	RegistrySha256 *string `json:"registrySha256,omitempty" tf:"registry_sha256,omitempty"`

	// The secret environment variables of the container.
	// The secret environment variables to be injected into your container at runtime.
	// +kubebuilder:validation:Optional
	SecretEnvironmentVariablesSecretRef *v1.SecretReference `json:"secretEnvironmentVariablesSecretRef,omitempty" tf:"-"`

	// The container status.
	// The container status
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	// The maximum amount of time in seconds during which your container can process a request before we stop it. Defaults to 300s.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
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
	InitProvider ContainerInitParameters `json:"initProvider,omitempty"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Container is the Schema for the Containers API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
