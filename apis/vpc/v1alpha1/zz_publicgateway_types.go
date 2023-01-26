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

type PublicGatewayObservation struct {

	// The date and time of the creation of the public gateway
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The date and time of the last update of the public gateway
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type PublicGatewayParameters struct {

	// Enable SSH bastion on the gateway
	// +kubebuilder:validation:Optional
	BastionEnabled *bool `json:"bastionEnabled,omitempty" tf:"bastion_enabled,omitempty"`

	// Port of the SSH bastion
	// +kubebuilder:validation:Optional
	BastionPort *float64 `json:"bastionPort,omitempty" tf:"bastion_port,omitempty"`

	// Enable SMTP on the gateway
	// +kubebuilder:validation:Optional
	EnableSMTP *bool `json:"enableSmtp,omitempty" tf:"enable_smtp,omitempty"`

	// attach an existing IP to the gateway
	// +crossplane:generate:reference:type=PublicGatewayIP
	// +kubebuilder:validation:Optional
	IPID *string `json:"ipId,omitempty" tf:"ip_id,omitempty"`

	// Reference to a PublicGatewayIP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDRef *v1.Reference `json:"ipIdRef,omitempty" tf:"-"`

	// Selector for a PublicGatewayIP to populate ipId.
	// +kubebuilder:validation:Optional
	IPIDSelector *v1.Selector `json:"ipIdSelector,omitempty" tf:"-"`

	// name of the gateway
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The tags associated with public gateway
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// gateway type
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// override the gateway's default recursive DNS servers, if DNS features are enabled
	// +kubebuilder:validation:Optional
	UpstreamDNSServers []*string `json:"upstreamDnsServers,omitempty" tf:"upstream_dns_servers,omitempty"`

	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// PublicGatewaySpec defines the desired state of PublicGateway
type PublicGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicGatewayParameters `json:"forProvider"`
}

// PublicGatewayStatus defines the observed state of PublicGateway.
type PublicGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicGateway is the Schema for the PublicGateways API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type PublicGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicGatewaySpec   `json:"spec"`
	Status            PublicGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicGatewayList contains a list of PublicGateways
type PublicGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicGateway `json:"items"`
}

// Repository type metadata.
var (
	PublicGateway_Kind             = "PublicGateway"
	PublicGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicGateway_Kind}.String()
	PublicGateway_KindAPIVersion   = PublicGateway_Kind + "." + CRDGroupVersion.String()
	PublicGateway_GroupVersionKind = CRDGroupVersion.WithKind(PublicGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicGateway{}, &PublicGatewayList{})
}