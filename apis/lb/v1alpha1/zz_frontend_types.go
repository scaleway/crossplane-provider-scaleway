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

type ACLInitParameters struct {

	// Action to undertake when an ACL filter matches.
	// Action to undertake when an ACL filter matches
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// Description of the ACL
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ACL match rule. At least ip_subnet or http_filter and http_filter_value are required.
	// The ACL match rule
	Match []MatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	// The ACL name. If not provided it will be randomly generated.
	// The ACL name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ACLObservation struct {

	// Action to undertake when an ACL filter matches.
	// Action to undertake when an ACL filter matches
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// IsDate and time of ACL's creation (RFC 3339 format)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the ACL
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ACL match rule. At least ip_subnet or http_filter and http_filter_value are required.
	// The ACL match rule
	Match []MatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	// The ACL name. If not provided it will be randomly generated.
	// The ACL name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// IsDate and time of ACL's update (RFC 3339 format)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ACLParameters struct {

	// Action to undertake when an ACL filter matches.
	// Action to undertake when an ACL filter matches
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// Description of the ACL
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ACL match rule. At least ip_subnet or http_filter and http_filter_value are required.
	// The ACL match rule
	// +kubebuilder:validation:Optional
	Match []MatchParameters `json:"match" tf:"match,omitempty"`

	// The ACL name. If not provided it will be randomly generated.
	// The ACL name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ActionInitParameters struct {

	// Redirect parameters when using an ACL with redirect action.
	// Redirect parameters when using an ACL with `redirect` action
	Redirect []RedirectInitParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The action type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionObservation struct {

	// Redirect parameters when using an ACL with redirect action.
	// Redirect parameters when using an ACL with `redirect` action
	Redirect []RedirectObservation `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The action type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionParameters struct {

	// Redirect parameters when using an ACL with redirect action.
	// Redirect parameters when using an ACL with `redirect` action
	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The action type
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type FrontendInitParameters struct {

	// A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
	// ACL rules
	ACL []ACLInitParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ID of the Load Balancer backend this frontend is attached to.
	// The load-balancer backend ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Backend
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// Reference to a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDRef *v1.Reference `json:"backendIdRef,omitempty" tf:"-"`

	// Selector for a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDSelector *v1.Selector `json:"backendIdSelector,omitempty" tf:"-"`

	// List of certificate IDs that should be used by the frontend.
	// Collection of Certificate IDs related to the load balancer and domain
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Default: false) Activates HTTP/3 protocol.
	// Activates HTTP/3 protocol
	EnableHttp3 *bool `json:"enableHttp3,omitempty" tf:"enable_http3,omitempty"`

	// (Defaults to false) A boolean to specify whether to use lb_acl.
	// If external_acls is set to true, acl can not be set directly in the Load Balancer frontend.
	// This boolean determines if ACLs should be managed externally through the 'lb_acl' resource. If set to `true`, `acl` attribute cannot be set directly in the lb frontend
	ExternalAcls *bool `json:"externalAcls,omitempty" tf:"external_acls,omitempty"`

	// TCP port to listen to on the front side.
	// TCP port to listen on the front side
	InboundPort *float64 `json:"inboundPort,omitempty" tf:"inbound_port,omitempty"`

	// The ID of the Load Balancer this frontend is attached to.
	// The load-balancer ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.LB
	LBID *string `json:"lbId,omitempty" tf:"lb_id,omitempty"`

	// Reference to a LB in lb to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDRef *v1.Reference `json:"lbIdRef,omitempty" tf:"-"`

	// Selector for a LB in lb to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDSelector *v1.Selector `json:"lbIdSelector,omitempty" tf:"-"`

	// The name of the Load Balancer frontend.
	// The name of the frontend
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Maximum inactivity time on the client side. (e.g. 1s)
	// Set the maximum inactivity time on the client side
	TimeoutClient *string `json:"timeoutClient,omitempty" tf:"timeout_client,omitempty"`
}

type FrontendObservation struct {

	// A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
	// ACL rules
	ACL []ACLObservation `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ID of the Load Balancer backend this frontend is attached to.
	// The load-balancer backend ID
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// (Deprecated, use certificate_ids instead) First certificate ID used by the frontend.
	// Certificate ID
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// List of certificate IDs that should be used by the frontend.
	// Collection of Certificate IDs related to the load balancer and domain
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Default: false) Activates HTTP/3 protocol.
	// Activates HTTP/3 protocol
	EnableHttp3 *bool `json:"enableHttp3,omitempty" tf:"enable_http3,omitempty"`

	// (Defaults to false) A boolean to specify whether to use lb_acl.
	// If external_acls is set to true, acl can not be set directly in the Load Balancer frontend.
	// This boolean determines if ACLs should be managed externally through the 'lb_acl' resource. If set to `true`, `acl` attribute cannot be set directly in the lb frontend
	ExternalAcls *bool `json:"externalAcls,omitempty" tf:"external_acls,omitempty"`

	// The ID of the Load Balancer frontend.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TCP port to listen to on the front side.
	// TCP port to listen on the front side
	InboundPort *float64 `json:"inboundPort,omitempty" tf:"inbound_port,omitempty"`

	// The ID of the Load Balancer this frontend is attached to.
	// The load-balancer ID
	LBID *string `json:"lbId,omitempty" tf:"lb_id,omitempty"`

	// The name of the Load Balancer frontend.
	// The name of the frontend
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Maximum inactivity time on the client side. (e.g. 1s)
	// Set the maximum inactivity time on the client side
	TimeoutClient *string `json:"timeoutClient,omitempty" tf:"timeout_client,omitempty"`
}

type FrontendParameters struct {

	// A list of ACL rules to apply to the Load Balancer frontend.  Defined below.
	// ACL rules
	// +kubebuilder:validation:Optional
	ACL []ACLParameters `json:"acl,omitempty" tf:"acl,omitempty"`

	// The ID of the Load Balancer backend this frontend is attached to.
	// The load-balancer backend ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.Backend
	// +kubebuilder:validation:Optional
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// Reference to a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDRef *v1.Reference `json:"backendIdRef,omitempty" tf:"-"`

	// Selector for a Backend in lb to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDSelector *v1.Selector `json:"backendIdSelector,omitempty" tf:"-"`

	// List of certificate IDs that should be used by the frontend.
	// Collection of Certificate IDs related to the load balancer and domain
	// +kubebuilder:validation:Optional
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// (Default: false) Activates HTTP/3 protocol.
	// Activates HTTP/3 protocol
	// +kubebuilder:validation:Optional
	EnableHttp3 *bool `json:"enableHttp3,omitempty" tf:"enable_http3,omitempty"`

	// (Defaults to false) A boolean to specify whether to use lb_acl.
	// If external_acls is set to true, acl can not be set directly in the Load Balancer frontend.
	// This boolean determines if ACLs should be managed externally through the 'lb_acl' resource. If set to `true`, `acl` attribute cannot be set directly in the lb frontend
	// +kubebuilder:validation:Optional
	ExternalAcls *bool `json:"externalAcls,omitempty" tf:"external_acls,omitempty"`

	// TCP port to listen to on the front side.
	// TCP port to listen on the front side
	// +kubebuilder:validation:Optional
	InboundPort *float64 `json:"inboundPort,omitempty" tf:"inbound_port,omitempty"`

	// The ID of the Load Balancer this frontend is attached to.
	// The load-balancer ID
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/lb/v1alpha1.LB
	// +kubebuilder:validation:Optional
	LBID *string `json:"lbId,omitempty" tf:"lb_id,omitempty"`

	// Reference to a LB in lb to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDRef *v1.Reference `json:"lbIdRef,omitempty" tf:"-"`

	// Selector for a LB in lb to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDSelector *v1.Selector `json:"lbIdSelector,omitempty" tf:"-"`

	// The name of the Load Balancer frontend.
	// The name of the frontend
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Maximum inactivity time on the client side. (e.g. 1s)
	// Set the maximum inactivity time on the client side
	// +kubebuilder:validation:Optional
	TimeoutClient *string `json:"timeoutClient,omitempty" tf:"timeout_client,omitempty"`
}

type MatchInitParameters struct {

	// The HTTP filter to match. This filter is supported only if your backend protocol has an HTTP forward protocol.
	// It extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part).
	// Possible values are: acl_http_filter_none, path_begin, path_end, http_header_match or regex.
	// The HTTP filter to match
	HTTPFilter *string `json:"httpFilter,omitempty" tf:"http_filter,omitempty"`

	// If you have http_filter at http_header_match, you can use this field to filter on the HTTP header's value.
	// You can use this field with http_header_match acl type to set the header name to filter
	HTTPFilterOption *string `json:"httpFilterOption,omitempty" tf:"http_filter_option,omitempty"`

	// A list of possible values to match for the given HTTP filter.
	// Keep in mind that in the case of http_header_match the HTTP header field name is case insensitive.
	// A list of possible values to match for the given HTTP filter
	HTTPFilterValue []*string `json:"httpFilterValue,omitempty" tf:"http_filter_value,omitempty"`

	// A list of IPs, or CIDR v4/v6 addresses of the session client, to match.
	// A list of IPs or CIDR v4/v6 addresses of the client of the session to match
	IPSubnet []*string `json:"ipSubnet,omitempty" tf:"ip_subnet,omitempty"`

	// If set to true, the condition will be of type "unless".
	// If set to true, the condition will be of type "unless"
	Invert *bool `json:"invert,omitempty" tf:"invert,omitempty"`
}

type MatchObservation struct {

	// The HTTP filter to match. This filter is supported only if your backend protocol has an HTTP forward protocol.
	// It extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part).
	// Possible values are: acl_http_filter_none, path_begin, path_end, http_header_match or regex.
	// The HTTP filter to match
	HTTPFilter *string `json:"httpFilter,omitempty" tf:"http_filter,omitempty"`

	// If you have http_filter at http_header_match, you can use this field to filter on the HTTP header's value.
	// You can use this field with http_header_match acl type to set the header name to filter
	HTTPFilterOption *string `json:"httpFilterOption,omitempty" tf:"http_filter_option,omitempty"`

	// A list of possible values to match for the given HTTP filter.
	// Keep in mind that in the case of http_header_match the HTTP header field name is case insensitive.
	// A list of possible values to match for the given HTTP filter
	HTTPFilterValue []*string `json:"httpFilterValue,omitempty" tf:"http_filter_value,omitempty"`

	// A list of IPs, or CIDR v4/v6 addresses of the session client, to match.
	// A list of IPs or CIDR v4/v6 addresses of the client of the session to match
	IPSubnet []*string `json:"ipSubnet,omitempty" tf:"ip_subnet,omitempty"`

	// If set to true, the condition will be of type "unless".
	// If set to true, the condition will be of type "unless"
	Invert *bool `json:"invert,omitempty" tf:"invert,omitempty"`
}

type MatchParameters struct {

	// The HTTP filter to match. This filter is supported only if your backend protocol has an HTTP forward protocol.
	// It extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part).
	// Possible values are: acl_http_filter_none, path_begin, path_end, http_header_match or regex.
	// The HTTP filter to match
	// +kubebuilder:validation:Optional
	HTTPFilter *string `json:"httpFilter,omitempty" tf:"http_filter,omitempty"`

	// If you have http_filter at http_header_match, you can use this field to filter on the HTTP header's value.
	// You can use this field with http_header_match acl type to set the header name to filter
	// +kubebuilder:validation:Optional
	HTTPFilterOption *string `json:"httpFilterOption,omitempty" tf:"http_filter_option,omitempty"`

	// A list of possible values to match for the given HTTP filter.
	// Keep in mind that in the case of http_header_match the HTTP header field name is case insensitive.
	// A list of possible values to match for the given HTTP filter
	// +kubebuilder:validation:Optional
	HTTPFilterValue []*string `json:"httpFilterValue,omitempty" tf:"http_filter_value,omitempty"`

	// A list of IPs, or CIDR v4/v6 addresses of the session client, to match.
	// A list of IPs or CIDR v4/v6 addresses of the client of the session to match
	// +kubebuilder:validation:Optional
	IPSubnet []*string `json:"ipSubnet,omitempty" tf:"ip_subnet,omitempty"`

	// If set to true, the condition will be of type "unless".
	// If set to true, the condition will be of type "unless"
	// +kubebuilder:validation:Optional
	Invert *bool `json:"invert,omitempty" tf:"invert,omitempty"`
}

type RedirectInitParameters struct {

	// The HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308.
	// The HTTP redirect code to use
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// A URL can be used in case of a location redirect (e.g. https://scaleway.com will redirect to this same URL). A scheme name (e.g. https, http, ftp, git) will replace the request's original scheme.
	// An URL can be used in case of a location redirect
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The redirect type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedirectObservation struct {

	// The HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308.
	// The HTTP redirect code to use
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// A URL can be used in case of a location redirect (e.g. https://scaleway.com will redirect to this same URL). A scheme name (e.g. https, http, ftp, git) will replace the request's original scheme.
	// An URL can be used in case of a location redirect
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The redirect type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RedirectParameters struct {

	// The HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308.
	// The HTTP redirect code to use
	// +kubebuilder:validation:Optional
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// A URL can be used in case of a location redirect (e.g. https://scaleway.com will redirect to this same URL). A scheme name (e.g. https, http, ftp, git) will replace the request's original scheme.
	// An URL can be used in case of a location redirect
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The action type. Possible values are: allow or deny or redirect.
	// The redirect type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// FrontendSpec defines the desired state of Frontend
type FrontendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontendParameters `json:"forProvider"`
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
	InitProvider FrontendInitParameters `json:"initProvider,omitempty"`
}

// FrontendStatus defines the observed state of Frontend.
type FrontendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Frontend is the Schema for the Frontends API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Frontend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.inboundPort) || (has(self.initProvider) && has(self.initProvider.inboundPort))",message="spec.forProvider.inboundPort is a required parameter"
	Spec   FrontendSpec   `json:"spec"`
	Status FrontendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontendList contains a list of Frontends
type FrontendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Frontend `json:"items"`
}

// Repository type metadata.
var (
	Frontend_Kind             = "Frontend"
	Frontend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Frontend_Kind}.String()
	Frontend_KindAPIVersion   = Frontend_Kind + "." + CRDGroupVersion.String()
	Frontend_GroupVersionKind = CRDGroupVersion.WithKind(Frontend_Kind)
)

func init() {
	SchemeBuilder.Register(&Frontend{}, &FrontendList{})
}
