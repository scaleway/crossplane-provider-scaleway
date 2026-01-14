package datawarehouse

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "datawarehouse"

// Configure adds configurations for scaleway_datawarehouse_user resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_datawarehouse_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "User"

		r.References["deployment_id"] = config.Reference{
			TerraformName: "scaleway_datawarehouse_deployment",
		}
	})

	p.AddResourceConfigurator("scaleway_datawarehouse_database", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Database"

		r.References["deployment_id"] = config.Reference{
			TerraformName: "scaleway_datawarehouse_deployment",
		}
	})

	p.AddResourceConfigurator("scaleway_datawarehouse_deployment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Deployment"
	})
}
