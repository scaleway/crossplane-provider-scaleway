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

type GeoIPInitParameters struct {

	// The list of matches. (Can be more than one).
	// The list of matches
	Matches []MatchesInitParameters `json:"matches,omitempty" tf:"matches,omitempty"`
}

type GeoIPObservation struct {

	// The list of matches. (Can be more than one).
	// The list of matches
	Matches []MatchesObservation `json:"matches,omitempty" tf:"matches,omitempty"`
}

type GeoIPParameters struct {

	// The list of matches. (Can be more than one).
	// The list of matches
	// +kubebuilder:validation:Optional
	Matches []MatchesParameters `json:"matches" tf:"matches,omitempty"`
}

type HTTPServiceInitParameters struct {

	// List of IPs to check.
	// IPs to check
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// Text to search.
	// Text to search
	MustContain *string `json:"mustContain,omitempty" tf:"must_contain,omitempty"`

	// Strategy to return an IP from the IPs list. Can be random, hashed, or all.
	// Strategy to return an IP from the IPs list
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// URL to match the must_contain text to validate an IP.
	// URL to match the must_contain text to validate an IP
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// User-agent used when checking the URL.
	// User-agent used when checking the URL
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type HTTPServiceObservation struct {

	// List of IPs to check.
	// IPs to check
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// Text to search.
	// Text to search
	MustContain *string `json:"mustContain,omitempty" tf:"must_contain,omitempty"`

	// Strategy to return an IP from the IPs list. Can be random, hashed, or all.
	// Strategy to return an IP from the IPs list
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// URL to match the must_contain text to validate an IP.
	// URL to match the must_contain text to validate an IP
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// User-agent used when checking the URL.
	// User-agent used when checking the URL
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type HTTPServiceParameters struct {

	// List of IPs to check.
	// IPs to check
	// +kubebuilder:validation:Optional
	Ips []*string `json:"ips" tf:"ips,omitempty"`

	// Text to search.
	// Text to search
	// +kubebuilder:validation:Optional
	MustContain *string `json:"mustContain" tf:"must_contain,omitempty"`

	// Strategy to return an IP from the IPs list. Can be random, hashed, or all.
	// Strategy to return an IP from the IPs list
	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`

	// URL to match the must_contain text to validate an IP.
	// URL to match the must_contain text to validate an IP
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`

	// User-agent used when checking the URL.
	// User-agent used when checking the URL
	// +kubebuilder:validation:Optional
	UserAgent *string `json:"userAgent,omitempty" tf:"user_agent,omitempty"`
}

type MatchesInitParameters struct {

	// List of continents (eg: EU for Europe, NA for North America, AS for Asia, etc.). Check the list of all continent codes.
	// List of continents (eg: EU for Europe, NA for North America, AS for Asia...). List of all continents code: https://api.scaleway.com/domain-private/v2beta1/continents
	Continents []*string `json:"continents,omitempty" tf:"continents,omitempty"`

	// List of countries (eg: FR for France, US for the United States, GB for Great Britain, etc.). Check the list of all country codes.
	// List of countries (eg: FR for France, US for the United States, GB for Great Britain...). List of all countries code: https://api.scaleway.com/domain-private/v2beta1/countries
	Countries []*string `json:"countries,omitempty" tf:"countries,omitempty"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the match result
	Data *string `json:"data,omitempty" tf:"data,omitempty"`
}

type MatchesObservation struct {

	// List of continents (eg: EU for Europe, NA for North America, AS for Asia, etc.). Check the list of all continent codes.
	// List of continents (eg: EU for Europe, NA for North America, AS for Asia...). List of all continents code: https://api.scaleway.com/domain-private/v2beta1/continents
	Continents []*string `json:"continents,omitempty" tf:"continents,omitempty"`

	// List of countries (eg: FR for France, US for the United States, GB for Great Britain, etc.). Check the list of all country codes.
	// List of countries (eg: FR for France, US for the United States, GB for Great Britain...). List of all countries code: https://api.scaleway.com/domain-private/v2beta1/countries
	Countries []*string `json:"countries,omitempty" tf:"countries,omitempty"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the match result
	Data *string `json:"data,omitempty" tf:"data,omitempty"`
}

type MatchesParameters struct {

	// List of continents (eg: EU for Europe, NA for North America, AS for Asia, etc.). Check the list of all continent codes.
	// List of continents (eg: EU for Europe, NA for North America, AS for Asia...). List of all continents code: https://api.scaleway.com/domain-private/v2beta1/continents
	// +kubebuilder:validation:Optional
	Continents []*string `json:"continents,omitempty" tf:"continents,omitempty"`

	// List of countries (eg: FR for France, US for the United States, GB for Great Britain, etc.). Check the list of all country codes.
	// List of countries (eg: FR for France, US for the United States, GB for Great Britain...). List of all countries code: https://api.scaleway.com/domain-private/v2beta1/countries
	// +kubebuilder:validation:Optional
	Countries []*string `json:"countries,omitempty" tf:"countries,omitempty"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the match result
	// +kubebuilder:validation:Optional
	Data *string `json:"data" tf:"data,omitempty"`
}

type RecordInitParameters struct {

	// The DNS zone of the domain. If the domain has no DNS zone, one will be automatically created.
	// The zone you want to add the record in
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/domain/v1alpha1.Zone
	DNSZone *string `json:"dnsZone,omitempty" tf:"dns_zone,omitempty"`

	// Reference to a Zone in domain to populate dnsZone.
	// +kubebuilder:validation:Optional
	DNSZoneRef *v1.Reference `json:"dnsZoneRef,omitempty" tf:"-"`

	// Selector for a Zone in domain to populate dnsZone.
	// +kubebuilder:validation:Optional
	DNSZoneSelector *v1.Selector `json:"dnsZoneSelector,omitempty" tf:"-"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the record
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The Geo IP provides DNS resolution based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. Check the documentation for more information.
	// Return record based on client localisation
	GeoIP []GeoIPInitParameters `json:"geoIp,omitempty" tf:"geo_ip,omitempty"`

	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs, by excluding the ones not responding to the given string to check. Check the documentation for more information.
	// Return record based on client localisation
	HTTPService []HTTPServiceInitParameters `json:"httpService,omitempty" tf:"http_service,omitempty"`

	// When destroying a resource, if only NS records remain and this is set to false, the zone will be deleted. Note that each zone not deleted will be billed.
	// When destroy a resource record, if a zone have only NS, delete the zone
	KeepEmptyZone *bool `json:"keepEmptyZone,omitempty" tf:"keep_empty_zone,omitempty"`

	// The name of the record (can be an empty string for a root record).
	// The name of the record
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the record (mostly used with an MX record).
	// The priority of the record
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the record.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Time To Live of the record in seconds.
	// The ttl of the record
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of the record (A, AAAA, MX, CNAME, DNAME, ALIAS, NS, PTR, SRV, TXT, TLSA, or CAA).
	// The type of the record
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The answer to a DNS request is based on the client’s (resolver) subnet. (Can be more than 1) Check the documentation for more information.
	// Return record based on client subnet
	View []ViewInitParameters `json:"view,omitempty" tf:"view,omitempty"`

	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with their related IP compared to the others in the list. (Can be more than 1) Check the documentation for more information.
	// Return record based on weight
	Weighted []WeightedInitParameters `json:"weighted,omitempty" tf:"weighted,omitempty"`
}

type RecordObservation struct {

	// The DNS zone of the domain. If the domain has no DNS zone, one will be automatically created.
	// The zone you want to add the record in
	DNSZone *string `json:"dnsZone,omitempty" tf:"dns_zone,omitempty"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the record
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The FQDN of the record.
	// The FQDN of the record
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Geo IP provides DNS resolution based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. Check the documentation for more information.
	// Return record based on client localisation
	GeoIP []GeoIPObservation `json:"geoIp,omitempty" tf:"geo_ip,omitempty"`

	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs, by excluding the ones not responding to the given string to check. Check the documentation for more information.
	// Return record based on client localisation
	HTTPService []HTTPServiceObservation `json:"httpService,omitempty" tf:"http_service,omitempty"`

	// The ID of the record.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When destroying a resource, if only NS records remain and this is set to false, the zone will be deleted. Note that each zone not deleted will be billed.
	// When destroy a resource record, if a zone have only NS, delete the zone
	KeepEmptyZone *bool `json:"keepEmptyZone,omitempty" tf:"keep_empty_zone,omitempty"`

	// The name of the record (can be an empty string for a root record).
	// The name of the record
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the record (mostly used with an MX record).
	// The priority of the record
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the record.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Does the DNS zone is the root zone or not
	RootZone *bool `json:"rootZone,omitempty" tf:"root_zone,omitempty"`

	// Time To Live of the record in seconds.
	// The ttl of the record
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of the record (A, AAAA, MX, CNAME, DNAME, ALIAS, NS, PTR, SRV, TXT, TLSA, or CAA).
	// The type of the record
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The answer to a DNS request is based on the client’s (resolver) subnet. (Can be more than 1) Check the documentation for more information.
	// Return record based on client subnet
	View []ViewObservation `json:"view,omitempty" tf:"view,omitempty"`

	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with their related IP compared to the others in the list. (Can be more than 1) Check the documentation for more information.
	// Return record based on weight
	Weighted []WeightedObservation `json:"weighted,omitempty" tf:"weighted,omitempty"`
}

type RecordParameters struct {

	// The DNS zone of the domain. If the domain has no DNS zone, one will be automatically created.
	// The zone you want to add the record in
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/domain/v1alpha1.Zone
	// +kubebuilder:validation:Optional
	DNSZone *string `json:"dnsZone,omitempty" tf:"dns_zone,omitempty"`

	// Reference to a Zone in domain to populate dnsZone.
	// +kubebuilder:validation:Optional
	DNSZoneRef *v1.Reference `json:"dnsZoneRef,omitempty" tf:"-"`

	// Selector for a Zone in domain to populate dnsZone.
	// +kubebuilder:validation:Optional
	DNSZoneSelector *v1.Selector `json:"dnsZoneSelector,omitempty" tf:"-"`

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the record
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The Geo IP provides DNS resolution based on the user’s geographical location. You can define a default IP that resolves if no Geo IP rule matches, and specify IPs for each geographical zone. Check the documentation for more information.
	// Return record based on client localisation
	// +kubebuilder:validation:Optional
	GeoIP []GeoIPParameters `json:"geoIp,omitempty" tf:"geo_ip,omitempty"`

	// The DNS service checks the provided URL on the configured IPs and resolves the request to one of the IPs, by excluding the ones not responding to the given string to check. Check the documentation for more information.
	// Return record based on client localisation
	// +kubebuilder:validation:Optional
	HTTPService []HTTPServiceParameters `json:"httpService,omitempty" tf:"http_service,omitempty"`

	// When destroying a resource, if only NS records remain and this is set to false, the zone will be deleted. Note that each zone not deleted will be billed.
	// When destroy a resource record, if a zone have only NS, delete the zone
	// +kubebuilder:validation:Optional
	KeepEmptyZone *bool `json:"keepEmptyZone,omitempty" tf:"keep_empty_zone,omitempty"`

	// The name of the record (can be an empty string for a root record).
	// The name of the record
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the record (mostly used with an MX record).
	// The priority of the record
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the record.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Time To Live of the record in seconds.
	// The ttl of the record
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of the record (A, AAAA, MX, CNAME, DNAME, ALIAS, NS, PTR, SRV, TXT, TLSA, or CAA).
	// The type of the record
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The answer to a DNS request is based on the client’s (resolver) subnet. (Can be more than 1) Check the documentation for more information.
	// Return record based on client subnet
	// +kubebuilder:validation:Optional
	View []ViewParameters `json:"view,omitempty" tf:"view,omitempty"`

	// You provide a list of IPs with their corresponding weights. These weights are used to proportionally direct requests to each IP. Depending on the weight of a record more or fewer requests are answered with their related IP compared to the others in the list. (Can be more than 1) Check the documentation for more information.
	// Return record based on weight
	// +kubebuilder:validation:Optional
	Weighted []WeightedParameters `json:"weighted,omitempty" tf:"weighted,omitempty"`
}

type ViewInitParameters struct {

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the view record
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The subnet of the view.
	// The subnet of the view
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type ViewObservation struct {

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the view record
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The subnet of the view.
	// The subnet of the view
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type ViewParameters struct {

	// The content of the record (an IPv4 for an A record, a string for a TXT record, etc.).
	// The data of the view record
	// +kubebuilder:validation:Optional
	Data *string `json:"data" tf:"data,omitempty"`

	// The subnet of the view.
	// The subnet of the view
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet" tf:"subnet,omitempty"`
}

type WeightedInitParameters struct {

	// The weighted IP.
	// The weighted IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The weight of the IP as an integer UInt32.
	// The weight of the IP
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WeightedObservation struct {

	// The weighted IP.
	// The weighted IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The weight of the IP as an integer UInt32.
	// The weight of the IP
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WeightedParameters struct {

	// The weighted IP.
	// The weighted IP
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`

	// The weight of the IP as an integer UInt32.
	// The weight of the IP
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
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
	InitProvider RecordInitParameters `json:"initProvider,omitempty"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Record is the Schema for the Records API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.data) || (has(self.initProvider) && has(self.initProvider.data))",message="spec.forProvider.data is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   RecordSpec   `json:"spec"`
	Status RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
