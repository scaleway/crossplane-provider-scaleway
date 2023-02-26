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

type IpsObservation struct {

	// The address of the IP.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The id of the option to enable. Use this endpoint to find the available options IDs.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The reverse of the IP.
	Reverse *string `json:"reverse,omitempty" tf:"reverse,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IpsParameters struct {
}

type OptionsObservation struct {

	// The name of the server.
	// name of the option
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type OptionsParameters struct {

	// The auto expiration date for compatible options
	// Auto expire the option after this date
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// The id of the option to enable. Use this endpoint to find the available options IDs.
	// IDs of the options
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type PrivateNetworkObservation struct {

	// The date and time of the creation of the private network.
	// The date and time of the creation of the private network
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The private network status.
	// The private network status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The date and time of the last update of the private network.
	// The date and time of the last update of the private network
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// The VLAN ID associated to the private network.
	// The VLAN ID associated to the private network
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

type PrivateNetworkParameters struct {

	// The id of the option to enable. Use this endpoint to find the available options IDs.
	// The private network ID
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type ServerObservation struct {

	// The domain of the server.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The id of the option to enable. Use this endpoint to find the available options IDs.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (List of) The IPs of the server.
	Ips []IpsObservation `json:"ips,omitempty" tf:"ips,omitempty"`

	// The ID of the offer.
	// ID of the server offer
	OfferID *string `json:"offerId,omitempty" tf:"offer_id,omitempty"`

	// The options to enable on the server.
	// ~> The options block supports:
	// The options to enable on server
	// +kubebuilder:validation:Optional
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`

	// The organization ID the server is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// The ID of the os.
	// The base image ID of the server
	OsID *string `json:"osId,omitempty" tf:"os_id,omitempty"`

	// The private networks to attach to the server. For more information, see the documentation
	// The private networks to attach to the server
	// +kubebuilder:validation:Optional
	PrivateNetwork []PrivateNetworkObservation `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`
}

type ServerParameters struct {

	// A description for the server.
	// Some description to associate to the server, max 255 characters
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The hostname of the server.
	// Hostname of the server
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The name of the server.
	// Name of the server
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The offer name or UUID of the baremetal server.
	// Use this endpoint to find the right offer.
	// ID or name of the server offer
	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// The options to enable on the server.
	// ~> The options block supports:
	// The options to enable on server
	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// The UUID of the os to install on the server.
	// Use this endpoint to find the right OS ID.
	// ~> Important: Updates to os will reinstall the server.
	// The base image of the server
	// +kubebuilder:validation:Required
	Os *string `json:"os" tf:"os,omitempty"`

	// Password used for the installation. May be required depending on used os.
	// Password used for the installation.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The private networks to attach to the server. For more information, see the documentation
	// The private networks to attach to the server
	// +kubebuilder:validation:Optional
	PrivateNetwork []PrivateNetworkParameters `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the server is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// If True, this boolean allows to reinstall the server on install config changes.
	// ~> Important: Updates to ssh_key_ids, user, password, service_user or service_password will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
	// If True, this boolean allows to reinstall the server on SSH key IDs, user or password changes
	// +kubebuilder:validation:Optional
	ReinstallOnConfigChanges *bool `json:"reinstallOnConfigChanges,omitempty" tf:"reinstall_on_config_changes,omitempty"`

	// List of SSH keys allowed to connect to the server.
	// Array of SSH key IDs allowed to SSH to the server
	//
	// **NOTE** : If you are attempting to update your SSH key IDs, it will induce the reinstall of your server.
	// If this behaviour is wanted, please set 'reinstall_on_ssh_key_changes' argument to true.
	// +crossplane:generate:reference:type=github.com/scaleway/provider-scaleway/apis/account/v1alpha1.SSHKey
	// +kubebuilder:validation:Optional
	SSHKeyIds []*string `json:"sshKeyIds,omitempty" tf:"ssh_key_ids,omitempty"`

	// References to SSHKey in account to populate sshKeyIds.
	// +kubebuilder:validation:Optional
	SSHKeyIdsRefs []v1.Reference `json:"sshKeyIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SSHKey in account to populate sshKeyIds.
	// +kubebuilder:validation:Optional
	SSHKeyIdsSelector *v1.Selector `json:"sshKeyIdsSelector,omitempty" tf:"-"`

	// Password used for the service to install. May be required depending on used os.
	// Password used for the service to install.
	// +kubebuilder:validation:Optional
	ServicePasswordSecretRef *v1.SecretKeySelector `json:"servicePasswordSecretRef,omitempty" tf:"-"`

	// User used for the service to install.
	// User used for the service to install.
	// +kubebuilder:validation:Optional
	ServiceUser *string `json:"serviceUser,omitempty" tf:"service_user,omitempty"`

	// The tags associated with the server.
	// Array of tags to associate with the server
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// User used for the installation.
	// User used for the installation.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// (Defaults to provider zone) The zone in which the server should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Manages Scaleway Compute Baremetal servers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}