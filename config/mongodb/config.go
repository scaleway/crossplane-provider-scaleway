package mongodb

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "mongodb"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_mongodb_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Instance"
	})

	p.AddResourceConfigurator("scaleway_mongodb_snapshot", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Snapshot"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_mongodb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_mongodb_user", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "User"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_mongodb_instance",
		}
	})
}
