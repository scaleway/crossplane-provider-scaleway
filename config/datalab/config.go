package datalab

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "datalab"

// Configure adds configurations for scaleway_datalab resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_datalab", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Datalab"

		r.References["private_network_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
	})
}
