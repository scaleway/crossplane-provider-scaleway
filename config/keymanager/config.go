package keymanager

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "keymanager"

// Configure adds configurations for scaleway_key_manager_key resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_key_manager_key", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Key"

	})
}
