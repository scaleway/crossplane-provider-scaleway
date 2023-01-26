/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	project "github.com/scaleway/provider-scaleway/internal/controller/account/project"
	sshkey "github.com/scaleway/provider-scaleway/internal/controller/account/sshkey"
	server "github.com/scaleway/provider-scaleway/internal/controller/applesilicon/server"
	serverbaremetal "github.com/scaleway/provider-scaleway/internal/controller/baremetal/server"
	ip "github.com/scaleway/provider-scaleway/internal/controller/flexibleip/ip"
	image "github.com/scaleway/provider-scaleway/internal/controller/instance/image"
	ipinstance "github.com/scaleway/provider-scaleway/internal/controller/instance/ip"
	placementgroup "github.com/scaleway/provider-scaleway/internal/controller/instance/placementgroup"
	securitygroup "github.com/scaleway/provider-scaleway/internal/controller/instance/securitygroup"
	securitygrouprule "github.com/scaleway/provider-scaleway/internal/controller/instance/securitygrouprule"
	serverinstance "github.com/scaleway/provider-scaleway/internal/controller/instance/server"
	snapshot "github.com/scaleway/provider-scaleway/internal/controller/instance/snapshot"
	userdata "github.com/scaleway/provider-scaleway/internal/controller/instance/userdata"
	volume "github.com/scaleway/provider-scaleway/internal/controller/instance/volume"
	device "github.com/scaleway/provider-scaleway/internal/controller/iot/device"
	hub "github.com/scaleway/provider-scaleway/internal/controller/iot/hub"
	network "github.com/scaleway/provider-scaleway/internal/controller/iot/network"
	route "github.com/scaleway/provider-scaleway/internal/controller/iot/route"
	cluster "github.com/scaleway/provider-scaleway/internal/controller/k8s/cluster"
	pool "github.com/scaleway/provider-scaleway/internal/controller/k8s/pool"
	backend "github.com/scaleway/provider-scaleway/internal/controller/lb/backend"
	certificate "github.com/scaleway/provider-scaleway/internal/controller/lb/certificate"
	frontend "github.com/scaleway/provider-scaleway/internal/controller/lb/frontend"
	iplb "github.com/scaleway/provider-scaleway/internal/controller/lb/ip"
	lb "github.com/scaleway/provider-scaleway/internal/controller/lb/lb"
	routelb "github.com/scaleway/provider-scaleway/internal/controller/lb/route"
	providerconfig "github.com/scaleway/provider-scaleway/internal/controller/providerconfig"
	acl "github.com/scaleway/provider-scaleway/internal/controller/rdb/acl"
	database "github.com/scaleway/provider-scaleway/internal/controller/rdb/database"
	databasebackup "github.com/scaleway/provider-scaleway/internal/controller/rdb/databasebackup"
	instance "github.com/scaleway/provider-scaleway/internal/controller/rdb/instance"
	privilege "github.com/scaleway/provider-scaleway/internal/controller/rdb/privilege"
	readreplica "github.com/scaleway/provider-scaleway/internal/controller/rdb/readreplica"
	user "github.com/scaleway/provider-scaleway/internal/controller/rdb/user"
	clusterredis "github.com/scaleway/provider-scaleway/internal/controller/redis/cluster"
	namespace "github.com/scaleway/provider-scaleway/internal/controller/registry/namespace"
	gatewaynetwork "github.com/scaleway/provider-scaleway/internal/controller/vpc/gatewaynetwork"
	privatenetwork "github.com/scaleway/provider-scaleway/internal/controller/vpc/privatenetwork"
	publicgateway "github.com/scaleway/provider-scaleway/internal/controller/vpc/publicgateway"
	publicgatewaydhcp "github.com/scaleway/provider-scaleway/internal/controller/vpc/publicgatewaydhcp"
	publicgatewayip "github.com/scaleway/provider-scaleway/internal/controller/vpc/publicgatewayip"
	publicgatewaypatrule "github.com/scaleway/provider-scaleway/internal/controller/vpc/publicgatewaypatrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		sshkey.Setup,
		server.Setup,
		serverbaremetal.Setup,
		ip.Setup,
		image.Setup,
		ipinstance.Setup,
		placementgroup.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serverinstance.Setup,
		snapshot.Setup,
		userdata.Setup,
		volume.Setup,
		device.Setup,
		hub.Setup,
		network.Setup,
		route.Setup,
		cluster.Setup,
		pool.Setup,
		backend.Setup,
		certificate.Setup,
		frontend.Setup,
		iplb.Setup,
		lb.Setup,
		routelb.Setup,
		providerconfig.Setup,
		acl.Setup,
		database.Setup,
		databasebackup.Setup,
		instance.Setup,
		privilege.Setup,
		readreplica.Setup,
		user.Setup,
		clusterredis.Setup,
		namespace.Setup,
		gatewaynetwork.Setup,
		privatenetwork.Setup,
		publicgateway.Setup,
		publicgatewaydhcp.Setup,
		publicgatewayip.Setup,
		publicgatewaypatrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
