/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"scaleway_account_project":                     config.NameAsIdentifier,
	"scaleway_account_ssh_key":                     config.NameAsIdentifier,
	"scaleway_apple_silicon_server":                config.NameAsIdentifier,
	"scaleway_baremetal_server":                    config.NameAsIdentifier,
	"scaleway_block_snapshot":                      config.NameAsIdentifier,
	"scaleway_block_volume":                        config.NameAsIdentifier,
	"scaleway_cockpit":                             config.NameAsIdentifier,
	"scaleway_cockpit_alert_manager":               config.NameAsIdentifier,
	"scaleway_cockpit_grafana_user":                config.NameAsIdentifier,
	"scaleway_cockpit_source":                      config.NameAsIdentifier,
	"scaleway_cockpit_token":                       config.NameAsIdentifier,
	"scaleway_container":                           config.NameAsIdentifier,
	"scaleway_container_cron":                      config.NameAsIdentifier,
	"scaleway_container_domain":                    config.NameAsIdentifier,
	"scaleway_container_namespace":                 config.NameAsIdentifier,
	"scaleway_container_token":                     config.NameAsIdentifier,
	"scaleway_domain_record":                       config.NameAsIdentifier,
	"scaleway_domain_registration":                 config.NameAsIdentifier,
	"scaleway_domain_zone":                         config.NameAsIdentifier,
	"scaleway_edge_services_backend_stage":         config.NameAsIdentifier,
	"scaleway_edge_services_cache_stage":           config.NameAsIdentifier,
	"scaleway_edge_services_dns_stage":             config.NameAsIdentifier,
	"scaleway_edge_services_head_stage":            config.NameAsIdentifier,
	"scaleway_edge_services_pipeline":              config.NameAsIdentifier,
	"scaleway_edge_services_plan":                  config.NameAsIdentifier,
	"scaleway_edge_services_route_stage":           config.NameAsIdentifier,
	"scaleway_edge_services_tls_stage":             config.NameAsIdentifier,
	"scaleway_edge_services_waf_stage":             config.NameAsIdentifier,
	"scaleway_flexible_ip":                         config.NameAsIdentifier,
	"scaleway_function":                            config.NameAsIdentifier,
	"scaleway_function_cron":                       config.NameAsIdentifier,
	"scaleway_function_domain":                     config.NameAsIdentifier,
	"scaleway_function_namespace":                  config.NameAsIdentifier,
	"scaleway_function_token":                      config.NameAsIdentifier,
	"scaleway_iam_api_key":                         config.NameAsIdentifier,
	"scaleway_iam_application":                     config.NameAsIdentifier,
	"scaleway_iam_group":                           config.NameAsIdentifier,
	"scaleway_iam_policy":                          config.NameAsIdentifier,
	"scaleway_iam_ssh_key":                         config.NameAsIdentifier,
	"scaleway_iam_user":                            config.NameAsIdentifier,
	"scaleway_inference_deployment":                config.NameAsIdentifier,
	"scaleway_inference_model":                     config.NameAsIdentifier,
	"scaleway_instance_image":                      config.NameAsIdentifier,
	"scaleway_instance_ip":                         config.NameAsIdentifier,
	"scaleway_instance_placement_group":            config.NameAsIdentifier,
	"scaleway_instance_private_nic":                config.NameAsIdentifier,
	"scaleway_instance_security_group":             config.NameAsIdentifier,
	"scaleway_instance_security_group_rules":       config.NameAsIdentifier,
	"scaleway_instance_server":                     config.NameAsIdentifier,
	"scaleway_instance_snapshot":                   config.NameAsIdentifier,
	"scaleway_instance_user_data":                  config.NameAsIdentifier,
	"scaleway_instance_volume":                     config.NameAsIdentifier,
	"scaleway_iot_hub":                             config.NameAsIdentifier,
	"scaleway_iot_device":                          config.NameAsIdentifier,
	"scaleway_iot_network":                         config.NameAsIdentifier,
	"scaleway_iot_route":                           config.NameAsIdentifier,
	"scaleway_ipam_ip":                             config.NameAsIdentifier,
	"scaleway_ipam_ip_reverse_dns":                 config.NameAsIdentifier,
	"scaleway_job_definition":                      config.NameAsIdentifier,
	"scaleway_k8s_acl":                             config.NameAsIdentifier,
	"scaleway_k8s_cluster":                         config.NameAsIdentifier,
	"scaleway_k8s_pool":                            config.NameAsIdentifier,
	"scaleway_lb":                                  config.NameAsIdentifier,
	"scaleway_lb_ip":                               config.NameAsIdentifier,
	"scaleway_lb_backend":                          config.NameAsIdentifier,
	"scaleway_lb_certificate":                      config.NameAsIdentifier,
	"scaleway_lb_frontend":                         config.NameAsIdentifier,
	"scaleway_lb_route":                            config.NameAsIdentifier,
	"scaleway_mnq_nats_account":                    config.NameAsIdentifier,
	"scaleway_mnq_nats_credentials":                config.NameAsIdentifier,
	"scaleway_mnq_sns":                             config.NameAsIdentifier,
	"scaleway_mnq_sns_credentials":                 config.NameAsIdentifier,
	"scaleway_mnq_sns_topic":                       config.NameAsIdentifier,
	"scaleway_mnq_sns_topic_subscription":          config.NameAsIdentifier,
	"scaleway_mnq_sqs":                             config.NameAsIdentifier,
	"scaleway_mnq_sqs_credentials":                 config.NameAsIdentifier,
	"scaleway_mnq_sqs_queue":                       config.NameAsIdentifier,
	"scaleway_mongodb_instance":                    config.NameAsIdentifier,
	"scaleway_mongodb_snapshot":                    config.NameAsIdentifier,
	"scaleway_object":                              config.NameAsIdentifier,
	"scaleway_object_bucket":                       config.NameAsIdentifier,
	"scaleway_object_bucket_acl":                   config.NameAsIdentifier,
	"scaleway_object_bucket_lock_configuration":    config.NameAsIdentifier,
	"scaleway_object_bucket_policy":                config.NameAsIdentifier,
	"scaleway_object_bucket_website_configuration": config.NameAsIdentifier,
	"scaleway_rdb_acl":                             config.NameAsIdentifier,
	"scaleway_rdb_database":                        config.NameAsIdentifier,
	"scaleway_rdb_database_backup":                 config.NameAsIdentifier,
	"scaleway_rdb_instance":                        config.NameAsIdentifier,
	"scaleway_rdb_privilege":                       config.NameAsIdentifier,
	"scaleway_rdb_snapshot":                        config.NameAsIdentifier,
	"scaleway_rdb_user":                            config.NameAsIdentifier,
	"scaleway_rdb_read_replica":                    config.NameAsIdentifier,
	"scaleway_redis_cluster":                       config.NameAsIdentifier,
	"scaleway_registry_namespace":                  config.NameAsIdentifier,
	"scaleway_sdb_sql_database":                    config.NameAsIdentifier,
	"scaleway_secret":                              config.NameAsIdentifier,
	"scaleway_secret_version":                      config.NameAsIdentifier,
	"scaleway_tem_blocked_list":                    config.NameAsIdentifier,
	"scaleway_tem_domain":                          config.NameAsIdentifier,
	"scaleway_tem_webhook":                         config.NameAsIdentifier,
	"scaleway_vpc":                                 config.NameAsIdentifier,
	"scaleway_vpc_acl":                             config.NameAsIdentifier,
	"scaleway_vpc_gateway_network":                 config.NameAsIdentifier,
	"scaleway_vpc_private_network":                 config.NameAsIdentifier,
	"scaleway_vpc_public_gateway":                  config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_dhcp":             config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_ip":               config.NameAsIdentifier,
	"scaleway_vpc_public_gateway_pat_rule":         config.NameAsIdentifier,
	"scaleway_key_manager_key":                     config.NameAsIdentifier,
	"scaleway_datawarehouse_user":                  config.NameAsIdentifier,
	"scaleway_datawarehouse_database":              config.NameAsIdentifier,
	"scaleway_datawarehouse_deployment":            config.NameAsIdentifier,
	"scaleway_file_filesystem":                     config.NameAsIdentifier,
	"scaleway_lb_private_network":                  config.NameAsIdentifier,
	"scaleway_autoscaling_instance_template":       config.NameAsIdentifier,
	"scaleway_autoscaling_instance_group":          config.NameAsIdentifier,
	"scaleway_mongodb_user":                        config.NameAsIdentifier,
	"scaleway_autoscaling_instance_policy":         config.NameAsIdentifier,
	"scaleway_vpc_route":                           config.NameAsIdentifier,
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
