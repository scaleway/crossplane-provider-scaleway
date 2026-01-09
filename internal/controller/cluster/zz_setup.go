/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	project "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/account/project"
	sshkey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/account/sshkey"
	server "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/applesilicon/server"
	group "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/autoscaling/group"
	policy "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/autoscaling/policy"
	template "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/autoscaling/template"
	serverbaremetal "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/baremetal/server"
	snapshot "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/block/snapshot"
	volume "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/block/volume"
	alertmanager "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/cockpit/alertmanager"
	cockpit "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/cockpit/cockpit"
	grafanauser "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/cockpit/grafanauser"
	source "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/cockpit/source"
	token "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/cockpit/token"
	container "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/container/container"
	containernamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/container/containernamespace"
	cron "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/container/cron"
	domain "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/container/domain"
	tokencontainer "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/container/token"
	database "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/datawarehouse/database"
	deployment "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/datawarehouse/deployment"
	user "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/datawarehouse/user"
	record "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/domain/record"
	registration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/domain/registration"
	zone "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/domain/zone"
	backendstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/backendstage"
	cachestage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/cachestage"
	dnsstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/dnsstage"
	headstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/headstage"
	pipeline "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/pipeline"
	plan "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/plan"
	routestage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/routestage"
	tlsstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/tlsstage"
	wafstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/edgeservices/wafstage"
	filesystem "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/file/filesystem"
	ip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/flexibleip/ip"
	cronfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/function/cron"
	domainfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/function/domain"
	function "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/function/function"
	functionnamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/function/functionnamespace"
	tokenfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/function/token"
	apikey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/apikey"
	application "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/application"
	groupiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/group"
	policyiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/policy"
	sshkeyiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/sshkey"
	useriam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iam/user"
	deploymentinference "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/inference/deployment"
	model "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/inference/model"
	image "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/image"
	ipinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/ip"
	placementgroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/placementgroup"
	privatenic "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/privatenic"
	securitygroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/securitygroup"
	securitygrouprule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/securitygrouprule"
	serverinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/server"
	snapshotinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/snapshot"
	userdata "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/userdata"
	volumeinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/instance/volume"
	device "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iot/device"
	hub "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iot/hub"
	network "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iot/network"
	route "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/iot/route"
	ipipam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/ipam/ip"
	ipreversedns "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/ipam/ipreversedns"
	definition "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/jobs/definition"
	acl "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/k8s/acl"
	cluster "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/k8s/cluster"
	pool "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/k8s/pool"
	key "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/keymanager/key"
	backend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/backend"
	certificate "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/certificate"
	frontend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/frontend"
	iplb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/ip"
	lb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/lb"
	privatenetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/privatenetwork"
	routelb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/lb/route"
	natsaccount "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/natsaccount"
	natscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/natscredentials"
	sns "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/sns"
	snscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/snscredentials"
	snstopic "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/snstopic"
	snstopicsubscription "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/snstopicsubscription"
	sqs "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/sqs"
	sqscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/sqscredentials"
	sqsqueue "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mnq/sqsqueue"
	instance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mongodb/instance"
	snapshotmongodb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mongodb/snapshot"
	usermongodb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/mongodb/user"
	aclobject "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/acl"
	bucket "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/bucket"
	lockconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/lockconfiguration"
	object "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/object"
	policyobject "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/policy"
	websiteconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/object/websiteconfiguration"
	providerconfig "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/providerconfig"
	aclrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/acl"
	databaserdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/database"
	databasebackup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/databasebackup"
	instancerdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/instance"
	privilege "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/privilege"
	readreplica "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/readreplica"
	snapshotrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/snapshot"
	userrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/rdb/user"
	clusterredis "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/redis/cluster"
	registrynamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/registry/registrynamespace"
	sqldatabase "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/sdb/sqldatabase"
	secret "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/secrets/secret"
	version "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/secrets/version"
	domaintem "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/tem/domain"
	list "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/tem/list"
	webhook "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/tem/webhook"
	aclvpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/acl"
	gatewaynetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/gatewaynetwork"
	privatenetworkvpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/privatenetwork"
	publicgateway "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/publicgateway"
	publicgatewaydhcp "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/publicgatewaydhcp"
	publicgatewayip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/publicgatewayip"
	publicgatewaypatrule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/publicgatewaypatrule"
	routevpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/route"
	vpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/cluster/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		sshkey.Setup,
		server.Setup,
		group.Setup,
		policy.Setup,
		template.Setup,
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
		database.Setup,
		deployment.Setup,
		user.Setup,
		record.Setup,
		registration.Setup,
		zone.Setup,
		backendstage.Setup,
		cachestage.Setup,
		dnsstage.Setup,
		headstage.Setup,
		pipeline.Setup,
		plan.Setup,
		routestage.Setup,
		tlsstage.Setup,
		wafstage.Setup,
		filesystem.Setup,
		ip.Setup,
		cronfunction.Setup,
		domainfunction.Setup,
		function.Setup,
		functionnamespace.Setup,
		tokenfunction.Setup,
		apikey.Setup,
		application.Setup,
		groupiam.Setup,
		policyiam.Setup,
		sshkeyiam.Setup,
		useriam.Setup,
		deploymentinference.Setup,
		model.Setup,
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
		acl.Setup,
		cluster.Setup,
		pool.Setup,
		key.Setup,
		backend.Setup,
		certificate.Setup,
		frontend.Setup,
		iplb.Setup,
		lb.Setup,
		privatenetwork.Setup,
		routelb.Setup,
		natsaccount.Setup,
		natscredentials.Setup,
		sns.Setup,
		snscredentials.Setup,
		snstopic.Setup,
		snstopicsubscription.Setup,
		sqs.Setup,
		sqscredentials.Setup,
		sqsqueue.Setup,
		instance.Setup,
		snapshotmongodb.Setup,
		usermongodb.Setup,
		aclobject.Setup,
		bucket.Setup,
		lockconfiguration.Setup,
		object.Setup,
		policyobject.Setup,
		websiteconfiguration.Setup,
		providerconfig.Setup,
		aclrdb.Setup,
		databaserdb.Setup,
		databasebackup.Setup,
		instancerdb.Setup,
		privilege.Setup,
		readreplica.Setup,
		snapshotrdb.Setup,
		userrdb.Setup,
		clusterredis.Setup,
		registrynamespace.Setup,
		sqldatabase.Setup,
		secret.Setup,
		version.Setup,
		domaintem.Setup,
		list.Setup,
		webhook.Setup,
		aclvpc.Setup,
		gatewaynetwork.Setup,
		privatenetworkvpc.Setup,
		publicgateway.Setup,
		publicgatewaydhcp.Setup,
		publicgatewayip.Setup,
		publicgatewaypatrule.Setup,
		routevpc.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.SetupGated,
		sshkey.SetupGated,
		server.SetupGated,
		group.SetupGated,
		policy.SetupGated,
		template.SetupGated,
		serverbaremetal.SetupGated,
		snapshot.SetupGated,
		volume.SetupGated,
		alertmanager.SetupGated,
		cockpit.SetupGated,
		grafanauser.SetupGated,
		source.SetupGated,
		token.SetupGated,
		container.SetupGated,
		containernamespace.SetupGated,
		cron.SetupGated,
		domain.SetupGated,
		tokencontainer.SetupGated,
		database.SetupGated,
		deployment.SetupGated,
		user.SetupGated,
		record.SetupGated,
		registration.SetupGated,
		zone.SetupGated,
		backendstage.SetupGated,
		cachestage.SetupGated,
		dnsstage.SetupGated,
		headstage.SetupGated,
		pipeline.SetupGated,
		plan.SetupGated,
		routestage.SetupGated,
		tlsstage.SetupGated,
		wafstage.SetupGated,
		filesystem.SetupGated,
		ip.SetupGated,
		cronfunction.SetupGated,
		domainfunction.SetupGated,
		function.SetupGated,
		functionnamespace.SetupGated,
		tokenfunction.SetupGated,
		apikey.SetupGated,
		application.SetupGated,
		groupiam.SetupGated,
		policyiam.SetupGated,
		sshkeyiam.SetupGated,
		useriam.SetupGated,
		deploymentinference.SetupGated,
		model.SetupGated,
		image.SetupGated,
		ipinstance.SetupGated,
		placementgroup.SetupGated,
		privatenic.SetupGated,
		securitygroup.SetupGated,
		securitygrouprule.SetupGated,
		serverinstance.SetupGated,
		snapshotinstance.SetupGated,
		userdata.SetupGated,
		volumeinstance.SetupGated,
		device.SetupGated,
		hub.SetupGated,
		network.SetupGated,
		route.SetupGated,
		ipipam.SetupGated,
		ipreversedns.SetupGated,
		definition.SetupGated,
		acl.SetupGated,
		cluster.SetupGated,
		pool.SetupGated,
		key.SetupGated,
		backend.SetupGated,
		certificate.SetupGated,
		frontend.SetupGated,
		iplb.SetupGated,
		lb.SetupGated,
		privatenetwork.SetupGated,
		routelb.SetupGated,
		natsaccount.SetupGated,
		natscredentials.SetupGated,
		sns.SetupGated,
		snscredentials.SetupGated,
		snstopic.SetupGated,
		snstopicsubscription.SetupGated,
		sqs.SetupGated,
		sqscredentials.SetupGated,
		sqsqueue.SetupGated,
		instance.SetupGated,
		snapshotmongodb.SetupGated,
		usermongodb.SetupGated,
		aclobject.SetupGated,
		bucket.SetupGated,
		lockconfiguration.SetupGated,
		object.SetupGated,
		policyobject.SetupGated,
		websiteconfiguration.SetupGated,
		providerconfig.SetupGated,
		aclrdb.SetupGated,
		databaserdb.SetupGated,
		databasebackup.SetupGated,
		instancerdb.SetupGated,
		privilege.SetupGated,
		readreplica.SetupGated,
		snapshotrdb.SetupGated,
		userrdb.SetupGated,
		clusterredis.SetupGated,
		registrynamespace.SetupGated,
		sqldatabase.SetupGated,
		secret.SetupGated,
		version.SetupGated,
		domaintem.SetupGated,
		list.SetupGated,
		webhook.SetupGated,
		aclvpc.SetupGated,
		gatewaynetwork.SetupGated,
		privatenetworkvpc.SetupGated,
		publicgateway.SetupGated,
		publicgatewaydhcp.SetupGated,
		publicgatewayip.SetupGated,
		publicgatewaypatrule.SetupGated,
		routevpc.SetupGated,
		vpc.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
