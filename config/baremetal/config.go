package baremetal

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "baremetal"

// Configure configures the elastic-metal resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_baremetal_server", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Server"

		r.References["ssh_key_ids"] = config.Reference{
			Type: "github.com/scaleway/crossplane-provider-scaleway/apis/account/v1alpha1.SSHKey",
		}
		r.References["ipam_ip_ids"] = config.Reference{
			TerraformName: "scaleway_ipam_ip",
		}
	})
}
