package applesilicon

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "applesilicon"

// Configure configures the account resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_apple_silicon_server", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Server"
	})
}
