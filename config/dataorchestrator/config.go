package dataorchestrator

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "dataorchestrator"

// Configure adds configurations for scaleway_dataorchestrator_workflow resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_dataorchestrator_workflow", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Workflow"

	})
}
