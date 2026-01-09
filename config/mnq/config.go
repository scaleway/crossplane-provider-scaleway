package mnq

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "mnq"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_mnq_sns_topic_subscription", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SNSTopicSubscription"

		r.References["access_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("access_key",false)`,
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
		r.References["secret_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("secret_key",false)`,
		}
		r.References["topic_id"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_topic",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sns", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SNS"

		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sqs_credentials", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SQSCredentials"

		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_nats_account", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "NATSAccount"
	})

	p.AddResourceConfigurator("scaleway_mnq_nats_credentials", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "NATSCredentials"

		r.References["account_id"] = config.Reference{
			TerraformName: "scaleway_mnq_nats_account",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sns_topic", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SNSTopic"

		r.References["access_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("access_key",false)`,
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_mnq_sns",
		}
		r.References["secret_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sns_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("secret_key",false)`,
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sqs_queue", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SQSQueue"

		r.References["access_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sqs_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("access_key",false)`,
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
		r.References["secret_key"] = config.Reference{
			TerraformName: "scaleway_mnq_sqs_credentials",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("secret_key",false)`,
		}
		r.References["sqs_endpoint"] = config.Reference{
			TerraformName: "scaleway_mnq_sqs",
			Extractor:     `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("endpoint",true)`,
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sqs", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SQS"

		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
	})

	p.AddResourceConfigurator("scaleway_mnq_sns_credentials", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SNSCredentials"

		r.References["project_id"] = config.Reference{
			TerraformName: "scaleway_account_project",
		}
	})
}
