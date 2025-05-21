/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/scaleway/crossplane-provider-scaleway/config/account"
	"github.com/scaleway/crossplane-provider-scaleway/config/applesilicon"
	"github.com/scaleway/crossplane-provider-scaleway/config/baremetal"
	"github.com/scaleway/crossplane-provider-scaleway/config/block"
	"github.com/scaleway/crossplane-provider-scaleway/config/cockpit"
	"github.com/scaleway/crossplane-provider-scaleway/config/container"
	"github.com/scaleway/crossplane-provider-scaleway/config/domain"
	"github.com/scaleway/crossplane-provider-scaleway/config/edgeservices"
	"github.com/scaleway/crossplane-provider-scaleway/config/flexibleip"
	"github.com/scaleway/crossplane-provider-scaleway/config/function"
	"github.com/scaleway/crossplane-provider-scaleway/config/iam"
	"github.com/scaleway/crossplane-provider-scaleway/config/inference"
	"github.com/scaleway/crossplane-provider-scaleway/config/instance"
	"github.com/scaleway/crossplane-provider-scaleway/config/iot"
	"github.com/scaleway/crossplane-provider-scaleway/config/ipam"
	"github.com/scaleway/crossplane-provider-scaleway/config/jobs"
	"github.com/scaleway/crossplane-provider-scaleway/config/k8s"
	"github.com/scaleway/crossplane-provider-scaleway/config/lb"
	"github.com/scaleway/crossplane-provider-scaleway/config/mongodb"
	"github.com/scaleway/crossplane-provider-scaleway/config/object"
	"github.com/scaleway/crossplane-provider-scaleway/config/rdb"
	"github.com/scaleway/crossplane-provider-scaleway/config/redis"
	"github.com/scaleway/crossplane-provider-scaleway/config/registry"
	"github.com/scaleway/crossplane-provider-scaleway/config/tem"
	"github.com/scaleway/crossplane-provider-scaleway/config/vpc"
)

const (
	resourcePrefix = "scaleway"
	modulePath     = "github.com/scaleway/crossplane-provider-scaleway"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		account.Configure,
		applesilicon.Configure,
		baremetal.Configure,
		block.Configure,
		cockpit.Configure,
		container.Configure,
		domain.Configure,
		flexibleip.Configure,
		function.Configure,
		iam.Configure,
		instance.Configure,
		iot.Configure,
		ipam.Configure,
		jobs.Configure,
		k8s.Configure,
		lb.Configure,
		object.Configure,
		rdb.Configure,
		redis.Configure,
		registry.Configure,
		tem.Configure,
		vpc.Configure,
		mongodb.Configure,
		edgeservices.Configure,
		inference.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
