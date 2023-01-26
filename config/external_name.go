/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"scaleway_account_ssh_key":               config.NameAsIdentifier,
	"scaleway_apple_silicon_server":          config.NameAsIdentifier,
	"scaleway_baremetal_server":              config.NameAsIdentifier,
	"scaleway_flexible_ip":                   config.NameAsIdentifier,
	"scaleway_instance_image":                config.NameAsIdentifier,
	"scaleway_instance_ip":                   config.NameAsIdentifier,
	"scaleway_instance_placement_group":      config.NameAsIdentifier,
	"scaleway_instance_security_group":       config.NameAsIdentifier,
	"scaleway_instance_security_group_rules": config.NameAsIdentifier,
	"scaleway_instance_server":               config.NameAsIdentifier,
	"scaleway_instance_snapshot":             config.NameAsIdentifier,
	"scaleway_instance_user_data":            config.NameAsIdentifier,
	"scaleway_instance_volume":               config.NameAsIdentifier,
	"scaleway_iot_hub":                       config.NameAsIdentifier,
	"scaleway_iot_device":                    config.NameAsIdentifier,
	"scaleway_iot_network":                   config.NameAsIdentifier,
	"scaleway_iot_route":                     config.NameAsIdentifier,
	"scaleway_k8s_cluster":                   config.NameAsIdentifier,
	"scaleway_k8s_pool":                      config.NameAsIdentifier,
	"scaleway_lb":                            config.NameAsIdentifier,
	"scaleway_lb_ip":                         config.NameAsIdentifier,
	"scaleway_lb_backend":                    config.NameAsIdentifier,
	"scaleway_lb_certificate":                config.NameAsIdentifier,
	"scaleway_lb_frontend":                   config.NameAsIdentifier,
	"scaleway_lb_route":                      config.NameAsIdentifier,
	"scaleway_rdb_acl":                       config.NameAsIdentifier,
	"scaleway_rdb_database":                  config.NameAsIdentifier,
	"scaleway_rdb_database_backup":           config.NameAsIdentifier,
	"scaleway_rdb_instance":                  config.NameAsIdentifier,
	"scaleway_rdb_privilege":                 config.NameAsIdentifier,
	"scaleway_rdb_user":                      config.NameAsIdentifier,
	"scaleway_rdb_read_replica":              config.NameAsIdentifier,
	"scaleway_registry_namespace":            config.NameAsIdentifier,
	"scaleway_vpc_gateway_network":           config.NameAsIdentifier,
	"scaleway_vpc_private_network":           config.NameAsIdentifier,
	"scaleway_vpc_public_gateway":            config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_dhcp":       config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_ip":         config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_pat_rule":   config.NameAsIdentifier,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
