package inference

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "inference"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_inference_model", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Model"
	})

	p.AddResourceConfigurator("scaleway_inference_deployment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Deployment"

		r.References["model_id"] = config.Reference{
			TerraformName: "scaleway_inference_model",
		}
	})
}
