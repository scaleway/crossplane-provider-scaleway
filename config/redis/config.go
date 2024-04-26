package redis

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "redis"

// Configure adds configurations for rdb resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_redis_cluster", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cluster"
	})
}
