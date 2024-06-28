package ipam

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "ipam"

// Configure adds configurations for scaleway_ipam_ip_reverse_dns resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_ipam_ip_reverse_dns", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "IPReverseDNS"

		r.References["ipam_ip_id"] = config.Reference{
			Type: "IP",
		}

	})

	p.AddResourceConfigurator("scaleway_ipam_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "IP"

		r.References["source.private_network_id"] = config.Reference{
			Type: "github.com/scaleway/provider-scaleway/apis/vpc/v1alpha1.PrivateNetwork",
		}
	})
}
