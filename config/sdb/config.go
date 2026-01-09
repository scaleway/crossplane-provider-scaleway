package sdb

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "sdb"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_sdb_sql_database", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "SQLDatabase"
	})
}
