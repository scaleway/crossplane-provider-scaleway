package file

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "file"

// Configure adds configurations for scaleway_file_filesystem resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_file_filesystem", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Filesystem"

	})
}
