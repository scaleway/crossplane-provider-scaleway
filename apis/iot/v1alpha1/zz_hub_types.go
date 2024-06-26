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

type HubInitParameters struct {

	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning *bool `json:"deviceAutoProvisioning,omitempty" tf:"device_auto_provisioning,omitempty"`

	// Whether to enable the hub events or not
	DisableEvents *bool `json:"disableEvents,omitempty" tf:"disable_events,omitempty"`

	// Wether the IoT Hub instance should be enabled or not.
	// Whether to enable the hub or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Topic prefix for the hub events
	EventsTopicPrefix *string `json:"eventsTopicPrefix,omitempty" tf:"events_topic_prefix,omitempty"`

	// Custom user provided certificate authority
	HubCA *string `json:"hubCa,omitempty" tf:"hub_ca,omitempty"`

	// Challenge certificate for the user provided hub CA
	HubCAChallenge *string `json:"hubCaChallenge,omitempty" tf:"hub_ca_challenge,omitempty"`

	// The name of the IoT Hub instance you want to create (e.g. my-hub).
	// The name of the hub
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Product plan to create the hub, see documentation for available product plans (e.g. plan_shared)
	// The product plan of the hub
	ProductPlan *string `json:"productPlan,omitempty" tf:"product_plan,omitempty"`

	// (Defaults to provider project_id) The ID of the project the IoT Hub Instance is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region in which the Database Instance should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type HubObservation struct {

	// The current number of connected devices in the Hub.
	// The current number of connected devices in the Hub
	ConnectedDeviceCount *float64 `json:"connectedDeviceCount,omitempty" tf:"connected_device_count,omitempty"`

	// The date and time the Hub was created.
	// The date and time of the creation of the IoT Hub
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Wether to enable the device auto provisioning or not
	DeviceAutoProvisioning *bool `json:"deviceAutoProvisioning,omitempty" tf:"device_auto_provisioning,omitempty"`

	// The number of registered devices in the Hub.
	// The number of registered devices in the Hub
	DeviceCount *float64 `json:"deviceCount,omitempty" tf:"device_count,omitempty"`

	// Whether to enable the hub events or not
	DisableEvents *bool `json:"disableEvents,omitempty" tf:"disable_events,omitempty"`

	// Wether the IoT Hub instance should be enabled or not.
	// Whether to enable the hub or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The MQTT network endpoint to connect MQTT devices to.
	// The endpoint to connect the devices to
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Topic prefix for the hub events
	EventsTopicPrefix *string `json:"eventsTopicPrefix,omitempty" tf:"events_topic_prefix,omitempty"`

	// Custom user provided certificate authority
	HubCA *string `json:"hubCa,omitempty" tf:"hub_ca,omitempty"`

	// Challenge certificate for the user provided hub CA
	HubCAChallenge *string `json:"hubCaChallenge,omitempty" tf:"hub_ca_challenge,omitempty"`

	// The ID of the Hub.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The MQTT certificat content
	// The MQTT certificat content
	MqttCA *string `json:"mqttCa,omitempty" tf:"mqtt_ca,omitempty"`

	// The MQTT ca url
	// The url of the MQTT ca
	MqttCAURL *string `json:"mqttCaUrl,omitempty" tf:"mqtt_ca_url,omitempty"`

	// The name of the IoT Hub instance you want to create (e.g. my-hub).
	// The name of the hub
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Hub.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Product plan to create the hub, see documentation for available product plans (e.g. plan_shared)
	// The product plan of the hub
	ProductPlan *string `json:"productPlan,omitempty" tf:"product_plan,omitempty"`

	// (Defaults to provider project_id) The ID of the project the IoT Hub Instance is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region in which the Database Instance should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The current status of the Hub.
	// The status of the hub
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The date and time the Hub resource was updated.
	// The date and time of the last update of the IoT Hub
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type HubParameters struct {

	// Wether to enable the device auto provisioning or not
	// +kubebuilder:validation:Optional
	DeviceAutoProvisioning *bool `json:"deviceAutoProvisioning,omitempty" tf:"device_auto_provisioning,omitempty"`

	// Whether to enable the hub events or not
	// +kubebuilder:validation:Optional
	DisableEvents *bool `json:"disableEvents,omitempty" tf:"disable_events,omitempty"`

	// Wether the IoT Hub instance should be enabled or not.
	// Whether to enable the hub or not
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Topic prefix for the hub events
	// +kubebuilder:validation:Optional
	EventsTopicPrefix *string `json:"eventsTopicPrefix,omitempty" tf:"events_topic_prefix,omitempty"`

	// Custom user provided certificate authority
	// +kubebuilder:validation:Optional
	HubCA *string `json:"hubCa,omitempty" tf:"hub_ca,omitempty"`

	// Challenge certificate for the user provided hub CA
	// +kubebuilder:validation:Optional
	HubCAChallenge *string `json:"hubCaChallenge,omitempty" tf:"hub_ca_challenge,omitempty"`

	// The name of the IoT Hub instance you want to create (e.g. my-hub).
	// The name of the hub
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Product plan to create the hub, see documentation for available product plans (e.g. plan_shared)
	// The product plan of the hub
	// +kubebuilder:validation:Optional
	ProductPlan *string `json:"productPlan,omitempty" tf:"product_plan,omitempty"`

	// (Defaults to provider project_id) The ID of the project the IoT Hub Instance is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region in which the Database Instance should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// HubSpec defines the desired state of Hub
type HubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubParameters `json:"forProvider"`
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
	InitProvider HubInitParameters `json:"initProvider,omitempty"`
}

// HubStatus defines the observed state of Hub.
type HubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Hub is the Schema for the Hubs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Hub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.productPlan) || (has(self.initProvider) && has(self.initProvider.productPlan))",message="spec.forProvider.productPlan is a required parameter"
	Spec   HubSpec   `json:"spec"`
	Status HubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubList contains a list of Hubs
type HubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hub `json:"items"`
}

// Repository type metadata.
var (
	Hub_Kind             = "Hub"
	Hub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Hub_Kind}.String()
	Hub_KindAPIVersion   = Hub_Kind + "." + CRDGroupVersion.String()
	Hub_GroupVersionKind = CRDGroupVersion.WithKind(Hub_Kind)
)

func init() {
	SchemeBuilder.Register(&Hub{}, &HubList{})
}
