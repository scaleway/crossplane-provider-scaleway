package lb

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "lb"

// Configure configures the lb resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_lb", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "LB"

		r.References["ip_id"] = config.Reference{
			TerraformName: "scaleway_lb_ip",
		}
		r.References["ipam_ids"] = config.Reference{
			TerraformName: "scaleway_ipam_ip",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_acl", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ACL"

		r.References["frontend_id"] = config.Reference{
			TerraformName: "scaleway_lb_frontend",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_backend", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Backend"

		r.References["lb_id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_certificate", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Certificate"

		r.References["lb_id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_frontend", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Frontend"

		r.References["lb_id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}

		r.References["backend_id"] = config.Reference{
			TerraformName: "scaleway_lb_backend",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_ip", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "IP"
	})

	p.AddResourceConfigurator("scaleway_lb_route", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Route"

		r.References["frontend_id"] = config.Reference{
			TerraformName: "scaleway_lb_frontend",
		}

		r.References["backend_id"] = config.Reference{
			TerraformName: "scaleway_lb_backend",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_private_network", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNetwork"

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
