package loadbalancers

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "loadbalancers"

// Configure adds configurations for scaleway_lb_private_network resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_lb_private_network", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Network"

		r.References["ipam_ip_ids"] = config.Reference{
			TerraformName: "scaleway_ipam_ip",
		}

		r.References["lb_id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}

		r.References["private_network_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}

	})
}
