package messageq

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "messageq"

// Configure adds configurations for scaleway_messageq_deployment resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_messageq_deployment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Deployment"
	})

	p.AddResourceConfigurator("scaleway_messageq_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "User"

		r.References["deployment_id"] = config.Reference{
			TerraformName: "scaleway_messageq_deployment",
		}
	})
}
