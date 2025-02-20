package rdb

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "rdb"

// Configure adds configurations for rdb resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_rdb_acl", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ACL"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_database", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Database"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_database_backup", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "DatabaseBackup"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_instance", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Instance"

		r.References["private_network.pn_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_privilege", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Privilege"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_read_replica", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ReadReplica"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})

	p.AddResourceConfigurator("scaleway_rdb_user", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "User"

		r.References["instance_id"] = config.Reference{
			TerraformName: "scaleway_rdb_instance",
		}
	})
}
