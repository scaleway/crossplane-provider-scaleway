/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	project "github.com/scaleway/crossplane-provider-scaleway/internal/controller/account/project"
	sshkey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/account/sshkey"
	server "github.com/scaleway/crossplane-provider-scaleway/internal/controller/applesilicon/server"
	serverbaremetal "github.com/scaleway/crossplane-provider-scaleway/internal/controller/baremetal/server"
	snapshot "github.com/scaleway/crossplane-provider-scaleway/internal/controller/block/snapshot"
	volume "github.com/scaleway/crossplane-provider-scaleway/internal/controller/block/volume"
	alertmanager "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cockpit/alertmanager"
	cockpit "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cockpit/cockpit"
	grafanauser "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cockpit/grafanauser"
	source "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cockpit/source"
	token "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cockpit/token"
	container "github.com/scaleway/crossplane-provider-scaleway/internal/controller/container/container"
	containernamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/container/containernamespace"
	cron "github.com/scaleway/crossplane-provider-scaleway/internal/controller/container/cron"
	domain "github.com/scaleway/crossplane-provider-scaleway/internal/controller/container/domain"
	tokencontainer "github.com/scaleway/crossplane-provider-scaleway/internal/controller/container/token"
	record "github.com/scaleway/crossplane-provider-scaleway/internal/controller/domain/record"
	zone "github.com/scaleway/crossplane-provider-scaleway/internal/controller/domain/zone"
	ip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/flexibleip/ip"
	cronfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/function/cron"
	domainfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/function/domain"
	function "github.com/scaleway/crossplane-provider-scaleway/internal/controller/function/function"
	functionnamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/function/functionnamespace"
	tokenfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/function/token"
	apikey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/apikey"
	application "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/application"
	group "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/group"
	policy "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/policy"
	sshkeyiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/sshkey"
	user "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iam/user"
	image "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/image"
	ipinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/ip"
	placementgroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/placementgroup"
	privatenic "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/privatenic"
	securitygroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/securitygroup"
	securitygrouprule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/securitygrouprule"
	serverinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/server"
	snapshotinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/snapshot"
	userdata "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/userdata"
	volumeinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/instance/volume"
	device "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iot/device"
	hub "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iot/hub"
	network "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iot/network"
	route "github.com/scaleway/crossplane-provider-scaleway/internal/controller/iot/route"
	ipipam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/ipam/ip"
	ipreversedns "github.com/scaleway/crossplane-provider-scaleway/internal/controller/ipam/ipreversedns"
	definition "github.com/scaleway/crossplane-provider-scaleway/internal/controller/jobs/definition"
	cluster "github.com/scaleway/crossplane-provider-scaleway/internal/controller/k8s/cluster"
	pool "github.com/scaleway/crossplane-provider-scaleway/internal/controller/k8s/pool"
	backend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/backend"
	certificate "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/certificate"
	frontend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/frontend"
	iplb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/ip"
	lb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/lb"
	routelb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/lb/route"
	acl "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/acl"
	bucket "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/bucket"
	lockconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/lockconfiguration"
	object "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/object"
	policyobject "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/policy"
	websiteconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/object/websiteconfiguration"
	providerconfig "github.com/scaleway/crossplane-provider-scaleway/internal/controller/providerconfig"
	aclrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/acl"
	database "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/database"
	databasebackup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/databasebackup"
	instance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/instance"
	privilege "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/privilege"
	readreplica "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/readreplica"
	userrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/rdb/user"
	clusterredis "github.com/scaleway/crossplane-provider-scaleway/internal/controller/redis/cluster"
	registrynamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/registry/registrynamespace"
	domaintem "github.com/scaleway/crossplane-provider-scaleway/internal/controller/tem/domain"
	gatewaynetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/gatewaynetwork"
	privatenetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/privatenetwork"
	publicgateway "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/publicgateway"
	publicgatewaydhcp "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/publicgatewaydhcp"
	publicgatewayip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/publicgatewayip"
	publicgatewaypatrule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/publicgatewaypatrule"
	vpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		sshkey.Setup,
		server.Setup,
		serverbaremetal.Setup,
		snapshot.Setup,
		volume.Setup,
		alertmanager.Setup,
		cockpit.Setup,
		grafanauser.Setup,
		source.Setup,
		token.Setup,
		container.Setup,
		containernamespace.Setup,
		cron.Setup,
		domain.Setup,
		tokencontainer.Setup,
		record.Setup,
		zone.Setup,
		ip.Setup,
		cronfunction.Setup,
		domainfunction.Setup,
		function.Setup,
		functionnamespace.Setup,
		tokenfunction.Setup,
		apikey.Setup,
		application.Setup,
		group.Setup,
		policy.Setup,
		sshkeyiam.Setup,
		user.Setup,
		image.Setup,
		ipinstance.Setup,
		placementgroup.Setup,
		privatenic.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		serverinstance.Setup,
		snapshotinstance.Setup,
		userdata.Setup,
		volumeinstance.Setup,
		device.Setup,
		hub.Setup,
		network.Setup,
		route.Setup,
		ipipam.Setup,
		ipreversedns.Setup,
		definition.Setup,
		cluster.Setup,
		pool.Setup,
		backend.Setup,
		certificate.Setup,
		frontend.Setup,
		iplb.Setup,
		lb.Setup,
		routelb.Setup,
		acl.Setup,
		bucket.Setup,
		lockconfiguration.Setup,
		object.Setup,
		policyobject.Setup,
		websiteconfiguration.Setup,
		providerconfig.Setup,
		aclrdb.Setup,
		database.Setup,
		databasebackup.Setup,
		instance.Setup,
		privilege.Setup,
		readreplica.Setup,
		userrdb.Setup,
		clusterredis.Setup,
		registrynamespace.Setup,
		domaintem.Setup,
		gatewaynetwork.Setup,
		privatenetwork.Setup,
		publicgateway.Setup,
		publicgatewaydhcp.Setup,
		publicgatewayip.Setup,
		publicgatewaypatrule.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
