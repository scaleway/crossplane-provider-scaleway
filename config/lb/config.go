package lb

import "github.com/upbound/upjet/pkg/config"

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
			Type: "IP",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_acl", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ACL"

		r.References["frontend_id"] = config.Reference{
			Type: "Frontend",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_backend", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Backend"

		r.References["lb_id"] = config.Reference{
			Type: "LB",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_certificate", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Certificate"

		r.References["lb_id"] = config.Reference{
			Type: "LB",
		}
	})

	p.AddResourceConfigurator("scaleway_lb_frontend", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Frontend"

		r.References["lb_id"] = config.Reference{
			Type: "LB",
		}

		r.References["backend_id"] = config.Reference{
			Type: "Backend",
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
			Type: "Frontend",
		}

		r.References["backend_id"] = config.Reference{
			Type: "Backend",
		}
	})
}
