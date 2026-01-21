/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	project "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/account/project"
	sshkey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/account/sshkey"
	runner "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/applesilicon/runner"
	server "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/applesilicon/server"
	group "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/autoscaling/group"
	policy "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/autoscaling/policy"
	template "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/autoscaling/template"
	serverbaremetal "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/baremetal/server"
	snapshot "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/block/snapshot"
	volume "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/block/volume"
	alertmanager "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/cockpit/alertmanager"
	cockpit "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/cockpit/cockpit"
	grafanauser "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/cockpit/grafanauser"
	source "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/cockpit/source"
	token "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/cockpit/token"
	container "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/container/container"
	containernamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/container/containernamespace"
	cron "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/container/cron"
	domain "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/container/domain"
	tokencontainer "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/container/token"
	database "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/datawarehouse/database"
	deployment "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/datawarehouse/deployment"
	user "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/datawarehouse/user"
	record "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/domain/record"
	registration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/domain/registration"
	zone "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/domain/zone"
	backendstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/backendstage"
	cachestage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/cachestage"
	dnsstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/dnsstage"
	headstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/headstage"
	pipeline "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/pipeline"
	plan "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/plan"
	routestage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/routestage"
	tlsstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/tlsstage"
	wafstage "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/edgeservices/wafstage"
	filesystem "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/file/filesystem"
	ip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/flexibleip/ip"
	cronfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/function/cron"
	domainfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/function/domain"
	function "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/function/function"
	functionnamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/function/functionnamespace"
	tokenfunction "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/function/token"
	apikey "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/apikey"
	application "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/application"
	groupiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/group"
	policyiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/policy"
	sshkeyiam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/sshkey"
	useriam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iam/user"
	deploymentinference "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/inference/deployment"
	model "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/inference/model"
	image "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/image"
	ipinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/ip"
	placementgroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/placementgroup"
	privatenic "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/privatenic"
	securitygroup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/securitygroup"
	securitygrouprule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/securitygrouprule"
	serverinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/server"
	snapshotinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/snapshot"
	userdata "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/userdata"
	volumeinstance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/instance/volume"
	device "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iot/device"
	hub "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iot/hub"
	network "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iot/network"
	route "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/iot/route"
	ipipam "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/ipam/ip"
	ipreversedns "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/ipam/ipreversedns"
	definition "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/jobs/definition"
	acl "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/k8s/acl"
	cluster "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/k8s/cluster"
	pool "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/k8s/pool"
	key "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/keymanager/key"
	backend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/backend"
	certificate "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/certificate"
	frontend "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/frontend"
	iplb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/ip"
	lb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/lb"
	privatenetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/privatenetwork"
	routelb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/lb/route"
	natsaccount "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/natsaccount"
	natscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/natscredentials"
	sns "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/sns"
	snscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/snscredentials"
	snstopic "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/snstopic"
	snstopicsubscription "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/snstopicsubscription"
	sqs "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/sqs"
	sqscredentials "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/sqscredentials"
	sqsqueue "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mnq/sqsqueue"
	instance "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mongodb/instance"
	snapshotmongodb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mongodb/snapshot"
	usermongodb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/mongodb/user"
	aclobject "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/acl"
	bucket "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/bucket"
	lockconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/lockconfiguration"
	object "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/object"
	policyobject "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/policy"
	websiteconfiguration "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/object/websiteconfiguration"
	providerconfig "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/providerconfig"
	aclrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/acl"
	databaserdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/database"
	databasebackup "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/databasebackup"
	instancerdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/instance"
	privilege "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/privilege"
	readreplica "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/readreplica"
	snapshotrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/snapshot"
	userrdb "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/rdb/user"
	clusterredis "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/redis/cluster"
	registrynamespace "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/registry/registrynamespace"
	connection "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/s2svpn/connection"
	gateway "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/s2svpn/gateway"
	policys2svpn "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/s2svpn/policy"
	sqldatabase "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/sdb/sqldatabase"
	secret "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/secrets/secret"
	version "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/secrets/version"
	domaintem "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/tem/domain"
	list "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/tem/list"
	webhook "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/tem/webhook"
	aclvpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/acl"
	gatewaynetwork "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/gatewaynetwork"
	privatenetworkvpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/privatenetwork"
	publicgateway "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/publicgateway"
	publicgatewaydhcp "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/publicgatewaydhcp"
	publicgatewayip "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/publicgatewayip"
	publicgatewaypatrule "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/publicgatewaypatrule"
	routevpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/route"
	vpc "github.com/scaleway/crossplane-provider-scaleway/internal/controller/namespaced/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		sshkey.Setup,
		runner.Setup,
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
		connection.Setup,
		gateway.Setup,
		gateway.Setup,
		policys2svpn.Setup,
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
		runner.SetupGated,
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
		connection.SetupGated,
		gateway.SetupGated,
		gateway.SetupGated,
		policys2svpn.SetupGated,
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
