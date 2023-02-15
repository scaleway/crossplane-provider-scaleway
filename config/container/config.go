package container

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "container"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_container", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Container"

		r.References["namespace_id"] = config.Reference{
			Type: "ContainerNamespace",
		}
	})

	p.AddResourceConfigurator("scaleway_container_cron", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cron"

		r.References["container_id"] = config.Reference{
			Type: "Container",
		}
	})

	p.AddResourceConfigurator("scaleway_container_domain", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Domain"

		r.References["container_id"] = config.Reference{
			Type: "Container",
		}
	})

	// Issues on using Namespace as Kind in Container
	// https://github.com/crossplane/terrajet/issues/234
	// https://github.com/kubernetes/kubernetes/pull/108382
	p.AddResourceConfigurator("scaleway_container_namespace", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ContainerNamespace"
	})

	p.AddResourceConfigurator("scaleway_container_token", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Token"

		r.References["container_id"] = config.Reference{
			Type: "Container",
		}
		r.References["namespace_id"] = config.Reference{
			Type: "ContainerNamespace",
		}
	})
}
