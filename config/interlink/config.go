package interlink

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "interlink"

// Configure adds configurations for scaleway_interlink_routing_policy resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_interlink_routing_policy", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Policy"

	})

	p.AddResourceConfigurator("scaleway_interlink_link", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Link"
	})
}
