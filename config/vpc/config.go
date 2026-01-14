package vpc

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "vpc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("scaleway_vpc", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "VPC"
	})

	p.AddResourceConfigurator("scaleway_vpc_private_network", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNetwork"

		r.References["vpc_id"] = config.Reference{
			TerraformName: "scaleway_vpc",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_public_gateway", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PublicGateway"

		r.References["ip_id"] = config.Reference{
			TerraformName: "scaleway_vpc_public_gateway_ip",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_gateway_network", func(r *config.Resource) {
		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "GatewayNetwork"

		// https://github.com/crossplane/upjet/v2/issues/197
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"static_address"},
		}

		r.References["gateway_id"] = config.Reference{
			TerraformName: "scaleway_vpc_public_gateway",
		}
		r.References["private_network_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
		r.References["dhcp_id"] = config.Reference{
			TerraformName: "scaleway_vpc_public_gateway_dhcp",
		}
		r.References["ipam_ip_id"] = config.Reference{
			TerraformName: "scaleway_ipam_ip",
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
			TerraformName: "scaleway_vpc_public_gateway",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Route"

		r.References["vpc_id"] = config.Reference{
			TerraformName: "scaleway_vpc",
		}
	})

	p.AddResourceConfigurator("scaleway_vpc_acl", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Acl"

		r.References["vpc_id"] = config.Reference{
			TerraformName: "scaleway_vpc",
		}
	})
}
