package kafka

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "kafka"

// Configure adds configurations for scaleway_kafka_cluster resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_kafka_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cluster"

		r.References["private_network.pn_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
	})
}
