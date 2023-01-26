package vpc

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("scaleway_vpc_private_network", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNetwork"
	})

	p.AddResourceConfigurator("scaleway_vpc_public_gateway", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PublicGateway"

		r.References["ip_id"] = config.Reference{
			Type: "PublicGatewayIP",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_gateway_network", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "GatewayNetwork"

		r.References["gateway_id"] = config.Reference{
			Type: "PublicGateway",
		}
		r.References["private_network_id"] = config.Reference{
			Type: "PrivateNetwork",
		}
		r.References["dhcp_id"] = config.Reference{
			Type: "PublicGatewayDHCP",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_public_gateway_dhcp", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PublicGatewayDHCP"
	})

	p.AddResourceConfigurator("scaleway_vpc_public_gateway_ip", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PublicGatewayIP"
	})

	p.AddResourceConfigurator("scaleway_vpc_public_gateway_pat_rule", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PublicGatewayPATRule"

		r.References["gateway_id"] = config.Reference{
			Type: "PublicGateway",
		}
	})
}
