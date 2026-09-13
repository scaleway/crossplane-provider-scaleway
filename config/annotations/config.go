package annotations

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "annotations"

// Configure adds configurations for scaleway_annotations_binding resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_annotations_binding", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Binding"

		r.References["srn"] = config.Reference{
			TerraformName: "scaleway_key_manager_key",
		}
		r.References["value_id"] = config.Reference{
			TerraformName: "scaleway_annotations_value",
		}
	})

	p.AddResourceConfigurator("scaleway_annotations_value", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Value"

		r.References["key_id"] = config.Reference{
			TerraformName: "scaleway_annotations_key",
		}
	})

	p.AddResourceConfigurator("scaleway_annotations_key", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Key"

	})
}
