package autoscaling

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "autoscaling"

// Configure adds configurations for scaleway_autoscaling_instance_template resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_autoscaling_instance_template", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Template"

		r.References["private_network_ids"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}

		r.References["volumes.from_snapshot.snapshot_id"] = config.Reference{
			TerraformName: "scaleway_block_snapshot",
		}

	})

	p.AddResourceConfigurator("scaleway_autoscaling_instance_group", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Group"

		r.References["load_balancer.backend_ids"] = config.Reference{
			TerraformName: "scaleway_lb_backend",
		}
		r.References["load_balancer.id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}
		r.References["load_balancer.private_network_id"] = config.Reference{
			TerraformName: "scaleway_vpc_private_network",
		}
		r.References["template_id"] = config.Reference{
			TerraformName: "scaleway_autoscaling_instance_template",
		}
	})

	p.AddResourceConfigurator("scaleway_autoscaling_instance_policy", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Policy"

		r.References["instance_group_id"] = config.Reference{
			TerraformName: "scaleway_autoscaling_instance_group",
		}
	})
}
