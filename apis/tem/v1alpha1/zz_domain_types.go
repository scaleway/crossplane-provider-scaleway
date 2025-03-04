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

type DomainInitParameters struct {

	// Accept the Scaleway Terms of Service
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// Enable automatic configuration options for the domain
	Autoconfig *bool `json:"autoconfig,omitempty" tf:"autoconfig,omitempty"`

	// The domain name used when sending emails
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DomainObservation struct {

	// Accept the Scaleway Terms of Service
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// Enable automatic configuration options for the domain
	Autoconfig *bool `json:"autoconfig,omitempty" tf:"autoconfig,omitempty"`

	// Date and time of domain's creation (RFC 3339 format)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// DKIM public key, as should be recorded in the DNS zone
	DKIMConfig *string `json:"dkimConfig,omitempty" tf:"dkim_config,omitempty"`

	// DMARC record for the domain, as should be recorded in the DNS zone
	DmarcConfig *string `json:"dmarcConfig,omitempty" tf:"dmarc_config,omitempty"`

	// DMARC name for the domain, as should be recorded in the DNS zone
	DmarcName *string `json:"dmarcName,omitempty" tf:"dmarc_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Error message if the last check failed
	LastError *string `json:"lastError,omitempty" tf:"last_error,omitempty"`

	// Date and time the domain was last found to be valid (RFC 3339 format)
	LastValidAt *string `json:"lastValidAt,omitempty" tf:"last_valid_at,omitempty"`

	// The Scaleway's blackhole MX server to use
	MxBlackhole *string `json:"mxBlackhole,omitempty" tf:"mx_blackhole,omitempty"`

	// The domain name used when sending emails
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Date and time of the next scheduled check (RFC 3339 format)
	NextCheckAt *string `json:"nextCheckAt,omitempty" tf:"next_check_at,omitempty"`

	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The domain's reputation
	Reputation []ReputationObservation `json:"reputation,omitempty" tf:"reputation,omitempty"`

	// Date and time of the revocation of the domain (RFC 3339 format)
	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	// SMTP host to use to send emails
	SMTPHost *string `json:"smtpHost,omitempty" tf:"smtp_host,omitempty"`

	// SMTP port to use to send emails over TLS. (Port 587)
	SMTPPort *float64 `json:"smtpPort,omitempty" tf:"smtp_port,omitempty"`

	// SMTP port to use to send emails over TLS. (Port 2587)
	SMTPPortAlternative *float64 `json:"smtpPortAlternative,omitempty" tf:"smtp_port_alternative,omitempty"`

	// SMTP port to use to send emails. (Port 25)
	SMTPPortUnsecure *float64 `json:"smtpPortUnsecure,omitempty" tf:"smtp_port_unsecure,omitempty"`

	// SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission
	SmtpsAuthUser *string `json:"smtpsAuthUser,omitempty" tf:"smtps_auth_user,omitempty"`

	// SMTPS port to use to send emails over TLS Wrapper. (Port 465)
	SmtpsPort *float64 `json:"smtpsPort,omitempty" tf:"smtps_port,omitempty"`

	// SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
	SmtpsPortAlternative *float64 `json:"smtpsPortAlternative,omitempty" tf:"smtps_port_alternative,omitempty"`

	// Snippet of the SPF record that should be registered in the DNS zone
	SpfConfig *string `json:"spfConfig,omitempty" tf:"spf_config,omitempty"`

	// Status of the domain
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainParameters struct {

	// Accept the Scaleway Terms of Service
	// +kubebuilder:validation:Optional
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// Enable automatic configuration options for the domain
	// +kubebuilder:validation:Optional
	Autoconfig *bool `json:"autoconfig,omitempty" tf:"autoconfig,omitempty"`

	// The domain name used when sending emails
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). Specifies the region where the domain is registered. If not specified, it defaults to the provider's region.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ReputationInitParameters struct {
}

type ReputationObservation struct {
	PreviousScore *float64 `json:"previousScore,omitempty" tf:"previous_score,omitempty"`

	PreviousScoredAt *string `json:"previousScoredAt,omitempty" tf:"previous_scored_at,omitempty"`

	Score *float64 `json:"score,omitempty" tf:"score,omitempty"`

	ScoredAt *string `json:"scoredAt,omitempty" tf:"scored_at,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReputationParameters struct {
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
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
	InitProvider DomainInitParameters `json:"initProvider,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Domain is the Schema for the Domains API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acceptTos) || (has(self.initProvider) && has(self.initProvider.acceptTos))",message="spec.forProvider.acceptTos is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DomainSpec   `json:"spec"`
	Status DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
