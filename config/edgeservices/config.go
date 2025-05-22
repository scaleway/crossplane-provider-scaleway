package edgeservices

import "github.com/crossplane/upjet/pkg/config"

const shortGroup = "edgeservices"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_edge_services_head_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "HeadStage"

		r.References["head_stage_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_dns_stage",
		}
		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_plan", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Plan"

	})

	p.AddResourceConfigurator("scaleway_edge_services_backend_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "BackendStage"

		r.References["lb_backend_config.lb_config.frontend_id"] = config.Reference{
			TerraformName: "scaleway_lb_frontend",
		}
		r.References["lb_backend_config.lb_config.id"] = config.Reference{
			TerraformName: "scaleway_lb",
		}
		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_cache_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "CacheStage"

		r.References["backend_stage_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_backend_stage",
		}
		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
		r.References["purge.pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_tls_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "TLSStage"

		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_waf_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "WAFStage"

		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_pipeline", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Pipeline"

	})

	p.AddResourceConfigurator("scaleway_edge_services_dns_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "DNSStage"

		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
	})

	p.AddResourceConfigurator("scaleway_edge_services_route_stage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "RouteStage"

		r.References["pipeline_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_pipeline",
		}
		r.References["rule.backend_stage_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_backend_stage",
		}
		r.References["waf_stage_id"] = config.Reference{
			TerraformName: "scaleway_edge_services_waf_stage",
		}
	})
}
