/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/scaleway/provider-scaleway/config/account"
	"github.com/scaleway/provider-scaleway/config/applesilicon"
	"github.com/scaleway/provider-scaleway/config/baremetal"
	"github.com/scaleway/provider-scaleway/config/container"
	"github.com/scaleway/provider-scaleway/config/domain"
	"github.com/scaleway/provider-scaleway/config/flexibleip"
	"github.com/scaleway/provider-scaleway/config/function"
	"github.com/scaleway/provider-scaleway/config/iam"
	"github.com/scaleway/provider-scaleway/config/instance"
	"github.com/scaleway/provider-scaleway/config/iot"
	"github.com/scaleway/provider-scaleway/config/k8s"
	"github.com/scaleway/provider-scaleway/config/lb"
	"github.com/scaleway/provider-scaleway/config/mnq"
	"github.com/scaleway/provider-scaleway/config/object"
	"github.com/scaleway/provider-scaleway/config/rdb"
	"github.com/scaleway/provider-scaleway/config/redis"
	"github.com/scaleway/provider-scaleway/config/registry"
	"github.com/scaleway/provider-scaleway/config/tem"
	"github.com/scaleway/provider-scaleway/config/vpc"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "scaleway"
	modulePath     = "github.com/scaleway/provider-scaleway"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		account.Configure,
		applesilicon.Configure,
		baremetal.Configure,
		container.Configure,
		domain.Configure,
		flexibleip.Configure,
		function.Configure,
		iam.Configure,
		instance.Configure,
		iot.Configure,
		k8s.Configure,
		lb.Configure,
		mnq.Configure,
		object.Configure,
		rdb.Configure,
		redis.Configure,
		registry.Configure,
		tem.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
