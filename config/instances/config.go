package instances

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "instances"

// Configure adds configurations for scaleway_instance_template resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_instance_template", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Template"

		r.References["filesystem_ids"] = config.Reference{
			TerraformName: "scaleway_file_filesystem",
		}
		r.References["placement_group_id"] = config.Reference{
			TerraformName: "scaleway_instance_placement_group",
		}
		r.References["private_networks"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
		r.References["security_group_id"] = config.Reference{
			TerraformName: "scaleway_instance_security_group",
		}
		r.References["windows_rdp_ssh_key_id"] = config.Reference{
			TerraformName: "scaleway_iam_ssh_key",
		}
	})
}
