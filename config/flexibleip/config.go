package flexibleip

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "flexibleip"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_flexible_ip", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Ip"
	})
}
