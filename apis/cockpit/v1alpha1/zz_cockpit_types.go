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

type CockpitInitParameters struct {

	// (Deprecated) Name of the plan to use. Available plans are: free, premium, and custom.
	// ~> Important: The plan field is deprecated. Any modification or selection will have no effect.
	// [DEPRECATED] The plan field is deprecated. Any modification or selection will have no effect.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Defaults to the Project specified in the provider's configuration) The ID of the Project the Cockpit is associated with.
	// The project_id you want to attach the resource to
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/account/v1alpha1.Project
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in account to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`
}

type CockpitObservation struct {

	// (Deprecated) A list of endpoints related to Cockpit, each with specific URLs:
	// [DEPRECATED] Endpoints list. Please use 'scaleway_cockpit_source' instead.
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Deprecated) Name of the plan to use. Available plans are: free, premium, and custom.
	// ~> Important: The plan field is deprecated. Any modification or selection will have no effect.
	// [DEPRECATED] The plan field is deprecated. Any modification or selection will have no effect.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Deprecated) The ID of the current pricing plan.
	// [DEPRECATED] The plan ID of the cockpit. This field is no longer relevant.
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// (Defaults to the Project specified in the provider's configuration) The ID of the Project the Cockpit is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// [DEPRECATED] Push_url
	PushURL []PushURLObservation `json:"pushUrl,omitempty" tf:"push_url,omitempty"`
}

type CockpitParameters struct {

	// (Deprecated) Name of the plan to use. Available plans are: free, premium, and custom.
	// ~> Important: The plan field is deprecated. Any modification or selection will have no effect.
	// [DEPRECATED] The plan field is deprecated. Any modification or selection will have no effect.
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// (Defaults to the Project specified in the provider's configuration) The ID of the Project the Cockpit is associated with.
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
}

type EndpointsInitParameters struct {
}

type EndpointsObservation struct {

	// (Deprecated) URL for the Alert manager.
	AlertmanagerURL *string `json:"alertmanagerUrl,omitempty" tf:"alertmanager_url,omitempty"`

	// (Deprecated) URL for Grafana.
	GrafanaURL *string `json:"grafanaUrl,omitempty" tf:"grafana_url,omitempty"`

	// (Deprecated) URL for logs to retrieve in the Data sources tab of the Scaleway console.
	LogsURL *string `json:"logsUrl,omitempty" tf:"logs_url,omitempty"`

	// (Deprecated) URL for metrics to retrieve in the Data sources tab of the Scaleway console.
	MetricsURL *string `json:"metricsUrl,omitempty" tf:"metrics_url,omitempty"`

	// (Deprecated) URL for traces to retrieve in the Data sources tab of the Scaleway console.
	TracesURL *string `json:"tracesUrl,omitempty" tf:"traces_url,omitempty"`
}

type EndpointsParameters struct {
}

type PushURLInitParameters struct {
}

type PushURLObservation struct {

	// (Deprecated) URL for logs to retrieve in the Data sources tab of the Scaleway console.
	PushLogsURL *string `json:"pushLogsUrl,omitempty" tf:"push_logs_url,omitempty"`

	// (Deprecated) URL for metrics to retrieve in the Data sources tab of the Scaleway console.
	PushMetricsURL *string `json:"pushMetricsUrl,omitempty" tf:"push_metrics_url,omitempty"`
}

type PushURLParameters struct {
}

// CockpitSpec defines the desired state of Cockpit
type CockpitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CockpitParameters `json:"forProvider"`
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
	InitProvider CockpitInitParameters `json:"initProvider,omitempty"`
}

// CockpitStatus defines the observed state of Cockpit.
type CockpitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CockpitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cockpit is the Schema for the Cockpits API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
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
