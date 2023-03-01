package cockpit

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "cockpit"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_cockpit", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cockpit"

		r.References["project_id"] = config.Reference{
			Type: "github.com/scaleway/provider-scaleway/apis/account/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("scaleway_cockpit_token", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Token"

		r.References["project_id"] = config.Reference{
			Type: "github.com/scaleway/provider-scaleway/apis/account/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("scaleway_cockpit_grafana_user", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "GrafanaUser"

		r.References["project_id"] = config.Reference{
			Type: "github.com/scaleway/provider-scaleway/apis/account/v1alpha1.Project",
		}
	})
}
