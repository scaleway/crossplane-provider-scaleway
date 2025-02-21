package function

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "function"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_function", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Function"

		r.References["namespace_id"] = config.Reference{
			TerraformName: "scaleway_function_namespace",
		}
	})

	p.AddResourceConfigurator("scaleway_function_cron", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cron"

		r.References["function_id"] = config.Reference{
			TerraformName: "scaleway_function",
		}
	})

	p.AddResourceConfigurator("scaleway_function_domain", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["function_id"] = config.Reference{
			TerraformName: "scaleway_function",
		}
	})

	// Issues on using Namespace as Kind in Function
	// https://github.com/crossplane/terrajet/issues/234
	// https://github.com/kubernetes/kubernetes/pull/108382
	p.AddResourceConfigurator("scaleway_function_namespace", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "FunctionNamespace"
	})

	p.AddResourceConfigurator("scaleway_function_token", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Token"

		r.References["function_id"] = config.Reference{
			TerraformName: "scaleway_function",
		}
		r.References["namespace_id"] = config.Reference{
			TerraformName: "scaleway_function_namespace",
		}
	})
}
