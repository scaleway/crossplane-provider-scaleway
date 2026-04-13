package objectstorage

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "objectstorage"

// Configure adds configurations for scaleway_object_bucket_server_side_encryption_configuration resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_object_bucket_server_side_encryption_configuration", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Configuration"

		r.References["bucket"] = config.Reference{
			TerraformName: "scaleway_object_bucket",
		}
	})
}
