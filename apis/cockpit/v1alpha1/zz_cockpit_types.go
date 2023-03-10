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

type CockpitObservation struct {

	// Endpoints
	// Endpoints
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CockpitParameters struct {

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
}

type EndpointsObservation struct {

	// The alertmanager URL
	AlertmanagerURL *string `json:"alertmanagerUrl,omitempty" tf:"alertmanager_url,omitempty"`

	// The grafana URL
	GrafanaURL *string `json:"grafanaUrl,omitempty" tf:"grafana_url,omitempty"`

	// The logs URL
	LogsURL *string `json:"logsUrl,omitempty" tf:"logs_url,omitempty"`

	// The metrics URL
	MetricsURL *string `json:"metricsUrl,omitempty" tf:"metrics_url,omitempty"`
}

type EndpointsParameters struct {
}

// CockpitSpec defines the desired state of Cockpit
type CockpitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CockpitParameters `json:"forProvider"`
}

// CockpitStatus defines the observed state of Cockpit.
type CockpitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CockpitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cockpit is the Schema for the Cockpits API. Manages Scaleway Cockpits.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Cockpit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CockpitSpec   `json:"spec"`
	Status            CockpitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CockpitList contains a list of Cockpits
type CockpitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cockpit `json:"items"`
}

// Repository type metadata.
var (
	Cockpit_Kind             = "Cockpit"
	Cockpit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cockpit_Kind}.String()
	Cockpit_KindAPIVersion   = Cockpit_Kind + "." + CRDGroupVersion.String()
	Cockpit_GroupVersionKind = CRDGroupVersion.WithKind(Cockpit_Kind)
)

func init() {
	SchemeBuilder.Register(&Cockpit{}, &CockpitList{})
}
