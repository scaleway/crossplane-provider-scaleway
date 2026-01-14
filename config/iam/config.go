package iam

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "iam"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_iam_api_key", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ApiKey"

		r.References["application_id"] = config.Reference{
			TerraformName: "scaleway_iam_application",
		}
	})

	p.AddResourceConfigurator("scaleway_iam_application", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Application"
	})

	p.AddResourceConfigurator("scaleway_iam_group", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Group"

		r.References["application_ids"] = config.Reference{
			TerraformName: "scaleway_iam_application",
		}
	})

	p.AddResourceConfigurator("scaleway_iam_policy", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Policy"

		r.References["application_id"] = config.Reference{
			TerraformName: "scaleway_iam_application",
		}
		r.References["group_id"] = config.Reference{
			TerraformName: "scaleway_iam_group",
		}
	})

	p.AddResourceConfigurator("scaleway_iam_ssh_key", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SSHKey"
	})

	p.AddResourceConfigurator("scaleway_iam_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "User"
	})
}
