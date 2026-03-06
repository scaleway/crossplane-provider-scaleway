package opensearch

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "opensearch"

// Configure adds configurations for scaleway_opensearch_deployment resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_opensearch_deployment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Deployment"

		r.References["password"] = config.Reference{
			TerraformName: "var",
		}
	})
}
