package iot

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "iot"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_iot_hub", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Hub"
	})

	p.AddResourceConfigurator("scaleway_iot_device", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Device"

		r.References["hub_id"] = config.Reference{
			TerraformName: "scaleway_iot_hub",
		}
	})

	p.AddResourceConfigurator("scaleway_iot_network", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Network"

		r.References["hub_id"] = config.Reference{
			TerraformName: "scaleway_iot_hub",
		}
	})

	p.AddResourceConfigurator("scaleway_iot_route", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Route"

		r.References["hub_id"] = config.Reference{
			TerraformName: "scaleway_iot_hub",
		}
	})
}
