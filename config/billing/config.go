package billing

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "billing"

// Configure adds configurations for scaleway_billing_budget resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_billing_budget", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Budget"

	})
}
