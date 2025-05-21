package tem

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "tem"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_tem_domain", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Domain"
	})

	p.AddResourceConfigurator("scaleway_tem_webhook", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Webhook"

		r.References["domain_id"] = config.Reference{
			TerraformName: "scaleway_tem_domain",
		}
		r.References["sns_arn"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_topic",
		}
	})

	p.AddResourceConfigurator("scaleway_tem_blocked_list", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "List"

		r.References["domain_id"] = config.Reference{
			TerraformName: "scaleway_tem_domain",
		}
	})
}
