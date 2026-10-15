package partner

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "partner"

// Configure adds configurations for scaleway_partner_organization resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_partner_organization", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Organization"
	})
}
