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

type LBInitParameters struct {

	// Defines whether to automatically assign a flexible public IP to the load-balancer.
	// Defines whether to automatically assign a flexible public IP to the load balancer
	AssignFlexibleIP *bool `json:"assignFlexibleIp,omitempty" tf:"assign_flexible_ip,omitempty"`

	// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
	// Defines whether to automatically assign a flexible public IPv6 to the load balancer
	AssignFlexibleIPv6 *bool `json:"assignFlexibleIpv6,omitempty" tf:"assign_flexible_ipv6,omitempty"`

	// The description of the load-balancer.
	// The description of the lb
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Deprecated) The ID of the associated LB IP. See below.
	// The load-balance public IP ID
	// +crossplane:generate:reference:type=IP
	IPID *string `json:"ipId,omitempty" tf:"ip_id,omitempty"`

	// Reference to a IP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDRef *v1.Reference `json:"ipIdRef,omitempty" tf:"-"`

	// Selector for a IP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDSelector *v1.Selector `json:"ipIdSelector,omitempty" tf:"-"`

	// The List of IP IDs to attach to the Load Balancer.
	// List of IP IDs to attach to the Load Balancer
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// The name of the load-balancer.
	// Name of the lb
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// will recreate the attachment.
	// List of private network to connect with your load balancer
	PrivateNetwork []PrivateNetworkInitParameters `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the load-balancer is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to false) The release_ip allow release the ip address associated with the load-balancers.
	// Release the IPs related to this load-balancer
	ReleaseIP *bool `json:"releaseIp,omitempty" tf:"release_ip,omitempty"`

	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check possible values.
	// Enforces minimal SSL version (in SSL/TLS offloading context)
	SSLCompatibilityLevel *string `json:"sslCompatibilityLevel,omitempty" tf:"ssl_compatibility_level,omitempty"`

	// The tags associated with the load-balancers.
	// Array of tags to associate with the load-balancer
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the load-balancer. Please check the migration section to upgrade the type.
	// The type of load-balancer you want to create
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone of the load-balancer.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LBObservation struct {

	// Defines whether to automatically assign a flexible public IP to the load-balancer.
	// Defines whether to automatically assign a flexible public IP to the load balancer
	AssignFlexibleIP *bool `json:"assignFlexibleIp,omitempty" tf:"assign_flexible_ip,omitempty"`

	// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
	// Defines whether to automatically assign a flexible public IPv6 to the load balancer
	AssignFlexibleIPv6 *bool `json:"assignFlexibleIpv6,omitempty" tf:"assign_flexible_ipv6,omitempty"`

	// The description of the load-balancer.
	// The description of the lb
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the load-balancer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load-balancer public IPv4 Address.
	// The load-balance public IPv4 address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// (Deprecated) The ID of the associated LB IP. See below.
	// The load-balance public IP ID
	IPID *string `json:"ipId,omitempty" tf:"ip_id,omitempty"`

	// The List of IP IDs to attach to the Load Balancer.
	// List of IP IDs to attach to the Load Balancer
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// The load-balancer public IPv6 Address.
	// The load-balance public IPv6 address
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// The name of the load-balancer.
	// Name of the lb
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the load-balancer is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// will recreate the attachment.
	// List of private network to connect with your load balancer
	PrivateNetwork []PrivateNetworkObservation `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the load-balancer is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region of the resource
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// (Defaults to false) The release_ip allow release the ip address associated with the load-balancers.
	// Release the IPs related to this load-balancer
	ReleaseIP *bool `json:"releaseIp,omitempty" tf:"release_ip,omitempty"`

	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check possible values.
	// Enforces minimal SSL version (in SSL/TLS offloading context)
	SSLCompatibilityLevel *string `json:"sslCompatibilityLevel,omitempty" tf:"ssl_compatibility_level,omitempty"`

	// The tags associated with the load-balancers.
	// Array of tags to associate with the load-balancer
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the load-balancer. Please check the migration section to upgrade the type.
	// The type of load-balancer you want to create
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone of the load-balancer.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LBParameters struct {

	// Defines whether to automatically assign a flexible public IP to the load-balancer.
	// Defines whether to automatically assign a flexible public IP to the load balancer
	// +kubebuilder:validation:Optional
	AssignFlexibleIP *bool `json:"assignFlexibleIp,omitempty" tf:"assign_flexible_ip,omitempty"`

	// Defines whether to automatically assign a flexible public IPv6 to the load-balancer.
	// Defines whether to automatically assign a flexible public IPv6 to the load balancer
	// +kubebuilder:validation:Optional
	AssignFlexibleIPv6 *bool `json:"assignFlexibleIpv6,omitempty" tf:"assign_flexible_ipv6,omitempty"`

	// The description of the load-balancer.
	// The description of the lb
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Deprecated) The ID of the associated LB IP. See below.
	// The load-balance public IP ID
	// +crossplane:generate:reference:type=IP
	// +kubebuilder:validation:Optional
	IPID *string `json:"ipId,omitempty" tf:"ip_id,omitempty"`

	// Reference to a IP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDRef *v1.Reference `json:"ipIdRef,omitempty" tf:"-"`

	// Selector for a IP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDSelector *v1.Selector `json:"ipIdSelector,omitempty" tf:"-"`

	// The List of IP IDs to attach to the Load Balancer.
	// List of IP IDs to attach to the Load Balancer
	// +kubebuilder:validation:Optional
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// The name of the load-balancer.
	// Name of the lb
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// will recreate the attachment.
	// List of private network to connect with your load balancer
	// +kubebuilder:validation:Optional
	PrivateNetwork []PrivateNetworkParameters `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the load-balancer is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to false) The release_ip allow release the ip address associated with the load-balancers.
	// Release the IPs related to this load-balancer
	// +kubebuilder:validation:Optional
	ReleaseIP *bool `json:"releaseIp,omitempty" tf:"release_ip,omitempty"`

	// Enforces minimal SSL version (in SSL/TLS offloading context). Please check possible values.
	// Enforces minimal SSL version (in SSL/TLS offloading context)
	// +kubebuilder:validation:Optional
	SSLCompatibilityLevel *string `json:"sslCompatibilityLevel,omitempty" tf:"ssl_compatibility_level,omitempty"`

	// The tags associated with the load-balancers.
	// Array of tags to associate with the load-balancer
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the load-balancer. Please check the migration section to upgrade the type.
	// The type of load-balancer you want to create
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Defaults to provider zone) The zone of the load-balancer.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNetworkInitParameters struct {

	// Set to true if you want to let DHCP assign IP addresses. See below.
	// Set to true if you want to let DHCP assign IP addresses
	DHCPConfig *bool `json:"dhcpConfig,omitempty" tf:"dhcp_config,omitempty"`

	// The ID of the Private Network to associate.
	// The Private Network ID
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// Define a local ip address of your choice for the load balancer instance. See below.
	// Define an IP address in the subnet of your private network that will be assigned to your load balancer instance
	StaticConfig []*string `json:"staticConfig,omitempty" tf:"static_config,omitempty"`
}

type PrivateNetworkObservation struct {

	// Set to true if you want to let DHCP assign IP addresses. See below.
	// Set to true if you want to let DHCP assign IP addresses
	DHCPConfig *bool `json:"dhcpConfig,omitempty" tf:"dhcp_config,omitempty"`

	// The ID of the Private Network to associate.
	// The Private Network ID
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// Define a local ip address of your choice for the load balancer instance. See below.
	// Define an IP address in the subnet of your private network that will be assigned to your load balancer instance
	StaticConfig []*string `json:"staticConfig,omitempty" tf:"static_config,omitempty"`

	// The status of private network connection
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Defaults to provider zone) The zone of the load-balancer.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNetworkParameters struct {

	// Set to true if you want to let DHCP assign IP addresses. See below.
	// Set to true if you want to let DHCP assign IP addresses
	// +kubebuilder:validation:Optional
	DHCPConfig *bool `json:"dhcpConfig,omitempty" tf:"dhcp_config,omitempty"`

	// The ID of the Private Network to associate.
	// The Private Network ID
	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId" tf:"private_network_id,omitempty"`

	// Define a local ip address of your choice for the load balancer instance. See below.
	// Define an IP address in the subnet of your private network that will be assigned to your load balancer instance
	// +kubebuilder:validation:Optional
	StaticConfig []*string `json:"staticConfig,omitempty" tf:"static_config,omitempty"`
}

// LBSpec defines the desired state of LB
type LBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBParameters `json:"forProvider"`
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
	InitProvider LBInitParameters `json:"initProvider,omitempty"`
}

// LBStatus defines the observed state of LB.
type LBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LB is the Schema for the LBs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type LB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   LBSpec   `json:"spec"`
	Status LBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBList contains a list of LBs
type LBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LB `json:"items"`
}

// Repository type metadata.
var (
	LB_Kind             = "LB"
	LB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LB_Kind}.String()
	LB_KindAPIVersion   = LB_Kind + "." + CRDGroupVersion.String()
	LB_GroupVersionKind = CRDGroupVersion.WithKind(LB_Kind)
)

func init() {
	SchemeBuilder.Register(&LB{}, &LBList{})
}
