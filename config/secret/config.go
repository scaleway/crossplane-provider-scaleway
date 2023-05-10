package secret

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "secret"

// Configure adds configurations for account resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_secret", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Secret"
	})

	p.AddResourceConfigurator("scaleway_secret_version", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Version"

		r.References["secret_id"] = config.Reference{
			Type: "Secret",
		}
	})
}
