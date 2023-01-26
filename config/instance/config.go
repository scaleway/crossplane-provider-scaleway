package instance

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "instance"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_instance_image", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Image"

		r.References["root_volume_id"] = config.Reference{
			Type: "Snapshot",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "IP"
	})

	p.AddResourceConfigurator("scaleway_instance_ip_reverse_dns", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "IPReverseDNS"

		r.References["ip_id"] = config.Reference{
			Type: "IP",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_placement_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PlacementGroup"
	})

	p.AddResourceConfigurator("scaleway_instance_private_nic", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNIC"

		r.References["server_id"] = config.Reference{
			Type: "Server",
		}
		/*		r.References["private_network_id"] = config.Reference{
				Type: "github.com/crossplane-contrib/provider-jet-scaleway/apis/vpc/v1alpha1.PrivateNetwork",
			}*/
	})

	p.AddResourceConfigurator("scaleway_instance_security_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SecurityGroup"
	})

	p.AddResourceConfigurator("scaleway_instance_security_group_rules", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SecurityGroupRule"

		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_server", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Server"

		r.References["ip_id"] = config.Reference{
			Type: "IP",
		}
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
		r.References["placement_group_id"] = config.Reference{
			Type: "PlacementGroup",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_snapshot", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Snapshot"

		r.References["volume_id"] = config.Reference{
			Type: "Volume",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_user_data", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "UserData"

		r.References["server_id"] = config.Reference{
			Type: "Server",
		}
	})

	p.AddResourceConfigurator("scaleway_instance_volume", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Volume"
	})
}
