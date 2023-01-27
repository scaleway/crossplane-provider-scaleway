package baremetal

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "baremetal"

// Configure configures the elastic-metal resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_baremetal_server", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Server"

		r.References["ssh_key_ids"] = config.Reference{
			Type: "github.com/scaleway/provider-scaleway/apis/account/v1alpha1.SSHKey",
		}
	})
}
