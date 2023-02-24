package mnq

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "mnq"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_mnq_credential", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Credential"

		r.References["namespace_id"] = config.Reference{
			Type: "MNQNamespace",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_namespace", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "MNQNamespace"
	})
}
