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

type ScopesInitParameters struct {

	// (Defaults to false) Permission to query logs.
	// Query logs
	QueryLogs *bool `json:"queryLogs,omitempty" tf:"query_logs,omitempty"`

	// (Defaults to false) Permission to query metrics.
	// Query metrics
	QueryMetrics *bool `json:"queryMetrics,omitempty" tf:"query_metrics,omitempty"`

	// (Defaults to false) Permission to query traces.
	// Query traces
	QueryTraces *bool `json:"queryTraces,omitempty" tf:"query_traces,omitempty"`

	// (Defaults to false) Permission to set up alerts.
	// Setup alerts
	SetupAlerts *bool `json:"setupAlerts,omitempty" tf:"setup_alerts,omitempty"`

	// (Defaults to false) Permission to set up logs rules.
	// Setup logs rules
	SetupLogsRules *bool `json:"setupLogsRules,omitempty" tf:"setup_logs_rules,omitempty"`

	// (Defaults to false) Permission to set up metrics rules.
	// Setup metrics rules
	SetupMetricsRules *bool `json:"setupMetricsRules,omitempty" tf:"setup_metrics_rules,omitempty"`

	// (Defaults to true) Permission to write logs.
	// Write logs
	WriteLogs *bool `json:"writeLogs,omitempty" tf:"write_logs,omitempty"`

	// (Defaults to true) Permission to write metrics.
	// Write metrics
	WriteMetrics *bool `json:"writeMetrics,omitempty" tf:"write_metrics,omitempty"`

	// (Defaults to false) Permission to write traces.
	// Write traces
	WriteTraces *bool `json:"writeTraces,omitempty" tf:"write_traces,omitempty"`
}

type ScopesObservation struct {

	// (Defaults to false) Permission to query logs.
	// Query logs
	QueryLogs *bool `json:"queryLogs,omitempty" tf:"query_logs,omitempty"`

	// (Defaults to false) Permission to query metrics.
	// Query metrics
	QueryMetrics *bool `json:"queryMetrics,omitempty" tf:"query_metrics,omitempty"`

	// (Defaults to false) Permission to query traces.
	// Query traces
	QueryTraces *bool `json:"queryTraces,omitempty" tf:"query_traces,omitempty"`

	// (Defaults to false) Permission to set up alerts.
	// Setup alerts
	SetupAlerts *bool `json:"setupAlerts,omitempty" tf:"setup_alerts,omitempty"`

	// (Defaults to false) Permission to set up logs rules.
	// Setup logs rules
	SetupLogsRules *bool `json:"setupLogsRules,omitempty" tf:"setup_logs_rules,omitempty"`

	// (Defaults to false) Permission to set up metrics rules.
	// Setup metrics rules
	SetupMetricsRules *bool `json:"setupMetricsRules,omitempty" tf:"setup_metrics_rules,omitempty"`

	// (Defaults to true) Permission to write logs.
	// Write logs
	WriteLogs *bool `json:"writeLogs,omitempty" tf:"write_logs,omitempty"`

	// (Defaults to true) Permission to write metrics.
	// Write metrics
	WriteMetrics *bool `json:"writeMetrics,omitempty" tf:"write_metrics,omitempty"`

	// (Defaults to false) Permission to write traces.
	// Write traces
	WriteTraces *bool `json:"writeTraces,omitempty" tf:"write_traces,omitempty"`
}

type ScopesParameters struct {

	// (Defaults to false) Permission to query logs.
	// Query logs
	// +kubebuilder:validation:Optional
	QueryLogs *bool `json:"queryLogs,omitempty" tf:"query_logs,omitempty"`

	// (Defaults to false) Permission to query metrics.
	// Query metrics
	// +kubebuilder:validation:Optional
	QueryMetrics *bool `json:"queryMetrics,omitempty" tf:"query_metrics,omitempty"`

	// (Defaults to false) Permission to query traces.
	// Query traces
	// +kubebuilder:validation:Optional
	QueryTraces *bool `json:"queryTraces,omitempty" tf:"query_traces,omitempty"`

	// (Defaults to false) Permission to set up alerts.
	// Setup alerts
	// +kubebuilder:validation:Optional
	SetupAlerts *bool `json:"setupAlerts,omitempty" tf:"setup_alerts,omitempty"`

	// (Defaults to false) Permission to set up logs rules.
	// Setup logs rules
	// +kubebuilder:validation:Optional
	SetupLogsRules *bool `json:"setupLogsRules,omitempty" tf:"setup_logs_rules,omitempty"`

	// (Defaults to false) Permission to set up metrics rules.
	// Setup metrics rules
	// +kubebuilder:validation:Optional
	SetupMetricsRules *bool `json:"setupMetricsRules,omitempty" tf:"setup_metrics_rules,omitempty"`

	// (Defaults to true) Permission to write logs.
	// Write logs
	// +kubebuilder:validation:Optional
	WriteLogs *bool `json:"writeLogs,omitempty" tf:"write_logs,omitempty"`

	// (Defaults to true) Permission to write metrics.
	// Write metrics
	// +kubebuilder:validation:Optional
	WriteMetrics *bool `json:"writeMetrics,omitempty" tf:"write_metrics,omitempty"`

	// (Defaults to false) Permission to write traces.
	// Write traces
	// +kubebuilder:validation:Optional
	WriteTraces *bool `json:"writeTraces,omitempty" tf:"write_traces,omitempty"`
}

type TokenInitParameters struct {

	// The name of the token.
	// The name of the token
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to the Project ID specified in the provider configuration) The ID of the Project the Cockpit is associated with.
	// The project_id you want to attach the resource to
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/account/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (Defaults to the region specified in the provider configuration) The region where the Cockpit token is located.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Scopes allowed, each with default values:
	// Endpoints
	Scopes []ScopesInitParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type TokenObservation struct {

	// The date and time of the creation of the Cockpit Token (Format ISO 8601)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the Cockpit token.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the token.
	// The name of the token
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to the Project ID specified in the provider configuration) The ID of the Project the Cockpit is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to the region specified in the provider configuration) The region where the Cockpit token is located.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Scopes allowed, each with default values:
	// Endpoints
	Scopes []ScopesObservation `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// The date and time of the last update of the Cockpit Token (Format ISO 8601)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type TokenParameters struct {

	// The name of the token.
	// The name of the token
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to the Project ID specified in the provider configuration) The ID of the Project the Cockpit is associated with.
	// The project_id you want to attach the resource to
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/account/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (Defaults to the region specified in the provider configuration) The region where the Cockpit token is located.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Scopes allowed, each with default values:
	// Endpoints
	// +kubebuilder:validation:Optional
	Scopes []ScopesParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
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
	InitProvider TokenInitParameters `json:"initProvider,omitempty"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Token is the Schema for the Tokens API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TokenSpec   `json:"spec"`
	Status TokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenList contains a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

// Repository type metadata.
var (
	Token_Kind             = "Token"
	Token_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Token_Kind}.String()
	Token_KindAPIVersion   = Token_Kind + "." + CRDGroupVersion.String()
	Token_GroupVersionKind = CRDGroupVersion.WithKind(Token_Kind)
)

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}
