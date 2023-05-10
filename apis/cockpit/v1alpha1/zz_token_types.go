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

type ScopesObservation struct {
}

type ScopesParameters struct {

	// (Defaults to false) Query logs
	// Query logs
	// +kubebuilder:validation:Optional
	QueryLogs *bool `json:"queryLogs,omitempty" tf:"query_logs,omitempty"`

	// (Defaults to false) Query metrics
	// Query metrics
	// +kubebuilder:validation:Optional
	QueryMetrics *bool `json:"queryMetrics,omitempty" tf:"query_metrics,omitempty"`

	// (Defaults to false) Setup alerts
	// Setup alerts
	// +kubebuilder:validation:Optional
	SetupAlerts *bool `json:"setupAlerts,omitempty" tf:"setup_alerts,omitempty"`

	// (Defaults to false) Setup logs rules
	// Setup logs rules
	// +kubebuilder:validation:Optional
	SetupLogsRules *bool `json:"setupLogsRules,omitempty" tf:"setup_logs_rules,omitempty"`

	// (Defaults to false) Setup metrics rules
	// Setup metrics rules
	// +kubebuilder:validation:Optional
	SetupMetricsRules *bool `json:"setupMetricsRules,omitempty" tf:"setup_metrics_rules,omitempty"`

	// (Defaults to true) Write logs
	// Write logs
	// +kubebuilder:validation:Optional
	WriteLogs *bool `json:"writeLogs,omitempty" tf:"write_logs,omitempty"`

	// (Defaults to true) Write metrics
	// Write metrics
	// +kubebuilder:validation:Optional
	WriteMetrics *bool `json:"writeMetrics,omitempty" tf:"write_metrics,omitempty"`
}

type TokenObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TokenParameters struct {

	// The name of the token
	// The name of the token
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the cockpit is associated with.
	// The project_id you want to attach the resource to
	// +crossplane:generate:reference:type=github.com/scaleway/provider-scaleway/apis/account/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// Allowed scopes
	// Endpoints
	// +kubebuilder:validation:Optional
	Scopes []ScopesParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Token is the Schema for the Tokens API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TokenSpec   `json:"spec"`
	Status            TokenStatus `json:"status,omitempty"`
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
