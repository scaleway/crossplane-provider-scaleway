package jobs

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "jobs"

// Configure adds configurations for scaleway_job_definition resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_job_definition", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Definition"
	})
}
