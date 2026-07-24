package billing

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "billing"

// Configure adds configurations for scaleway_billing_budget_alert resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_billing_budget_alert", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Alert"

		r.References["budget_id"] = config.Reference{
			TerraformName: "scaleway_billing_budget",
		}
	})

	p.AddResourceConfigurator("scaleway_billing_budget", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Budget"

	})

	p.AddResourceConfigurator("scaleway_billing_budget_alert_notification", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Notification"

		r.References["budget_alert_id"] = config.Reference{
			TerraformName: "scaleway_billing_budget_alert",
		}
	})
}
