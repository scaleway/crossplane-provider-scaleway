package account

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "account"

// Configure adds configurations for account resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_account_project", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Project"
	})

	p.AddResourceConfigurator("scaleway_account_ssh_key", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SSHKey"
	})
}
