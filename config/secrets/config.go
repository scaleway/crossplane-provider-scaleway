package secrets

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "secrets"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_secret_version", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Version"

		r.References["secret_id"] = config.Reference{
			TerraformName: "scaleway_secret",
		}
	})

	p.AddResourceConfigurator("scaleway_secret", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Secret"
	})
}
