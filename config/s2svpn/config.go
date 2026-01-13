package s2svpn

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "s2svpn"

// Configure adds configurations for scaleway_s2s_vpn_routing_policy resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_s2s_vpn_routing_policy", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Policy"

	})

	p.AddResourceConfigurator("scaleway_s2s_vpn_connection", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Connection"

		r.References["bgp_config_ipv4.routing_policy_id"] = config.Reference{
			TerraformName: "scaleway_s2s_vpn_routing_policy",
		}
		r.References["customer_gateway_id"] = config.Reference{
			TerraformName: "scaleway_s2s_vpn_customer_gateway",
		}
		r.References["vpn_gateway_id"] = config.Reference{
			TerraformName: "scaleway_s2s_vpn_gateway",
		}
	})

	p.AddResourceConfigurator("scaleway_s2s_vpn_customer_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Gateway"

		r.References["ipv4_public"] = config.Reference{
			TerraformName: "scaleway_instance_ip",
		}
	})

	p.AddResourceConfigurator("scaleway_s2s_vpn_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Gateway"

		r.References["private_network_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
	})
}
