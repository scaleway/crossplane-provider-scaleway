package domain

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "domain"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_domain_record", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Record"

		r.References["dns_zone"] = config.Reference{
			TerraformName: "scaleway_domain_zone",
		}
	})

	p.AddResourceConfigurator("scaleway_domain_zone", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Zone"
	})

	p.AddResourceConfigurator("scaleway_domain_registration", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Registration"
	})
}
