package mailbox

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "mailbox"

// Configure adds configurations for scaleway_mailbox_domain resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_mailbox_domain", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
	})

	p.AddResourceConfigurator("scaleway_mailbox", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Mailbox"

		r.References["domain_id"] = config.Reference{
			TerraformName: "scaleway_mailbox_domain",
		}
	})
}
