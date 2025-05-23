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

type PrivateIpsInitParameters struct {
}

type PrivateIpsObservation struct {

	// The private IP address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the private NIC.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateIpsParameters struct {
}

type PrivateNICInitParameters struct {

	// IPAM ip list, should be for internal use only
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network.
	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/ipam/v1alpha1.IP
	IpamIPIds []*string `json:"ipamIpIds,omitempty" tf:"ipam_ip_ids,omitempty"`

	// References to IP in ipam to populate ipamIpIds.
	// +kubebuilder:validation:Optional
	IpamIPIdsRefs []v1.Reference `json:"ipamIpIdsRefs,omitempty" tf:"-"`

	// Selector for a list of IP in ipam to populate ipamIpIds.
	// +kubebuilder:validation:Optional
	IpamIPIdsSelector *v1.Selector `json:"ipamIpIdsSelector,omitempty" tf:"-"`

	// The ID of the private network attached to.
	// The private network ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/vpc/v1alpha1.PrivateNetwork
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// Reference to a PrivateNetwork in vpc to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDRef *v1.Reference `json:"privateNetworkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in vpc to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDSelector *v1.Selector `json:"privateNetworkIdSelector,omitempty" tf:"-"`

	// The ID of the server associated with.
	// The server ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/instance/v1alpha1.Server
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The tags associated with the private NIC.
	// The tags associated with the private-nic
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the server must be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNICObservation struct {

	// The ID of the private NIC.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPAM ip list, should be for internal use only
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network.
	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network
	IpamIPIds []*string `json:"ipamIpIds,omitempty" tf:"ipam_ip_ids,omitempty"`

	// The MAC address of the private NIC.
	// MAC address of the NIC
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The list of private IPv4 and IPv6 addresses associated with the resource.
	// List of private IPv4 and IPv6 addresses associated with the resource
	PrivateIps []PrivateIpsObservation `json:"privateIps,omitempty" tf:"private_ips,omitempty"`

	// The ID of the private network attached to.
	// The private network ID
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// The ID of the server associated with.
	// The server ID
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// The tags associated with the private NIC.
	// The tags associated with the private-nic
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the server must be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNICParameters struct {

	// IPAM ip list, should be for internal use only
	// +kubebuilder:validation:Optional
	IPIds []*string `json:"ipIds,omitempty" tf:"ip_ids,omitempty"`

	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network.
	// IPAM IDs of a pre-reserved IP addresses to assign to the Instance in the requested private network
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/ipam/v1alpha1.IP
	// +kubebuilder:validation:Optional
	IpamIPIds []*string `json:"ipamIpIds,omitempty" tf:"ipam_ip_ids,omitempty"`

	// References to IP in ipam to populate ipamIpIds.
	// +kubebuilder:validation:Optional
	IpamIPIdsRefs []v1.Reference `json:"ipamIpIdsRefs,omitempty" tf:"-"`

	// Selector for a list of IP in ipam to populate ipamIpIds.
	// +kubebuilder:validation:Optional
	IpamIPIdsSelector *v1.Selector `json:"ipamIpIdsSelector,omitempty" tf:"-"`

	// The ID of the private network attached to.
	// The private network ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/vpc/v1alpha1.PrivateNetwork
	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// Reference to a PrivateNetwork in vpc to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDRef *v1.Reference `json:"privateNetworkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in vpc to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDSelector *v1.Selector `json:"privateNetworkIdSelector,omitempty" tf:"-"`

	// The ID of the server associated with.
	// The server ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/instance/v1alpha1.Server
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// The tags associated with the private NIC.
	// The tags associated with the private-nic
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the server must be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// PrivateNICSpec defines the desired state of PrivateNIC
type PrivateNICSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateNICParameters `json:"forProvider"`
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
	InitProvider PrivateNICInitParameters `json:"initProvider,omitempty"`
}

// PrivateNICStatus defines the observed state of PrivateNIC.
type PrivateNICStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateNICObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateNIC is the Schema for the PrivateNICs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type PrivateNIC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateNICSpec   `json:"spec"`
	Status            PrivateNICStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateNICList contains a list of PrivateNICs
type PrivateNICList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateNIC `json:"items"`
}

// Repository type metadata.
var (
	PrivateNIC_Kind             = "PrivateNIC"
	PrivateNIC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateNIC_Kind}.String()
	PrivateNIC_KindAPIVersion   = PrivateNIC_Kind + "." + CRDGroupVersion.String()
	PrivateNIC_GroupVersionKind = CRDGroupVersion.WithKind(PrivateNIC_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateNIC{}, &PrivateNICList{})
}
