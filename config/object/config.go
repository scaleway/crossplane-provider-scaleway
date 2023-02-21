package object

import (
	"github.com/upbound/upjet/pkg/config"
)

const shortGroup = "object"

// Configure configures the account resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_object", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Object"
	})

	p.AddResourceConfigurator("scaleway_object_bucket", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Bucket"
	})

	p.AddResourceConfigurator("scaleway_object_bucket_acl", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "ACL"
	})

	p.AddResourceConfigurator("scaleway_object_bucket_lock_configuration", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "LockConfiguration"
	})

	p.AddResourceConfigurator("scaleway_object_bucket_policy", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Policy"
	})

	p.AddResourceConfigurator("scaleway_object_bucket_website_configuration", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "WebsiteConfiguration"
	})
}
