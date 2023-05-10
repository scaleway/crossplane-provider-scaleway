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

type IpObservation struct {

	// The date and time of the creation of the Flexible IP (Format ISO 8601)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the Flexible IP
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IPv4 address of the Flexible IP
	// The IPv4 address of the flexible IP
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The MAC address of the server associated with this flexible IP
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The organization of the Flexible IP
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The date and time of the last update of the Flexible IP (Format ISO 8601)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type IpParameters struct {

	// :  A description of the flexible IP.
	// Description of the flexible IP
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The project of the Flexible IP
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The reverse domain associated with this flexible IP.
	// The reverse DNS for this flexible IP
	// +kubebuilder:validation:Optional
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	// The ID of the associated server
	// The baremetal server associated with this flexible IP
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// :  A list of tags to apply to the flexible IP.
	// The tags associated with the flexible IP
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone of the Flexible IP
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// IpSpec defines the desired state of Ip
type IpSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpParameters `json:"forProvider"`
}

// IpStatus defines the observed state of Ip.
type IpStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ip is the Schema for the Ips API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Ip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpSpec   `json:"spec"`
	Status            IpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpList contains a list of Ips
type IpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ip `json:"items"`
}

// Repository type metadata.
var (
	Ip_Kind             = "Ip"
	Ip_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ip_Kind}.String()
	Ip_KindAPIVersion   = Ip_Kind + "." + CRDGroupVersion.String()
	Ip_GroupVersionKind = CRDGroupVersion.WithKind(Ip_Kind)
)

func init() {
	SchemeBuilder.Register(&Ip{}, &IpList{})
}
