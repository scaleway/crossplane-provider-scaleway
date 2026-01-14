package block

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "block"

// Configure adds configurations for scaleway_block_snapshot resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_block_snapshot", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Snapshot"

		r.References["volume_id"] = config.Reference{
			TerraformName: "scaleway_block_volume",
		}
	})

	p.AddResourceConfigurator("scaleway_block_volume", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Volume"

		r.References["snapshot_id"] = config.Reference{
			TerraformName: "scaleway_block_snapshot",
		}
	})
}
