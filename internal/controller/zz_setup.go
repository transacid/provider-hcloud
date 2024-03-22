/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	firewall "github.com/transacid/provider-hcloud/internal/controller/firewall/firewall"
	attachment "github.com/transacid/provider-hcloud/internal/controller/firewallattachment/attachment"
	ip "github.com/transacid/provider-hcloud/internal/controller/floatingip/ip"
	ipassignment "github.com/transacid/provider-hcloud/internal/controller/floatingipassignment/ipassignment"
	balancer "github.com/transacid/provider-hcloud/internal/controller/loadbalancer/balancer"
	balancernetwork "github.com/transacid/provider-hcloud/internal/controller/loadbalancernetwork/balancernetwork"
	balancerservice "github.com/transacid/provider-hcloud/internal/controller/loadbalancerservice/balancerservice"
	balancertarget "github.com/transacid/provider-hcloud/internal/controller/loadbalancertarget/balancertarget"
	certificate "github.com/transacid/provider-hcloud/internal/controller/managedcertificate/certificate"
	network "github.com/transacid/provider-hcloud/internal/controller/network/network"
	route "github.com/transacid/provider-hcloud/internal/controller/networkroute/route"
	subnet "github.com/transacid/provider-hcloud/internal/controller/networksubnet/subnet"
	group "github.com/transacid/provider-hcloud/internal/controller/placementgroup/group"
	ipprimaryip "github.com/transacid/provider-hcloud/internal/controller/primaryip/ip"
	providerconfig "github.com/transacid/provider-hcloud/internal/controller/providerconfig"
	rdns "github.com/transacid/provider-hcloud/internal/controller/rdns/rdns"
	server "github.com/transacid/provider-hcloud/internal/controller/server/server"
	networkservernetwork "github.com/transacid/provider-hcloud/internal/controller/servernetwork/network"
	snapshot "github.com/transacid/provider-hcloud/internal/controller/snapshot/snapshot"
	key "github.com/transacid/provider-hcloud/internal/controller/sshkey/key"
	certificateuploadedcertificate "github.com/transacid/provider-hcloud/internal/controller/uploadedcertificate/certificate"
	volume "github.com/transacid/provider-hcloud/internal/controller/volume/volume"
	attachmentvolumeattachment "github.com/transacid/provider-hcloud/internal/controller/volumeattachment/attachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.Setup,
		attachment.Setup,
		ip.Setup,
		ipassignment.Setup,
		balancer.Setup,
		balancernetwork.Setup,
		balancerservice.Setup,
		balancertarget.Setup,
		certificate.Setup,
		network.Setup,
		route.Setup,
		subnet.Setup,
		group.Setup,
		ipprimaryip.Setup,
		providerconfig.Setup,
		rdns.Setup,
		server.Setup,
		networkservernetwork.Setup,
		snapshot.Setup,
		key.Setup,
		certificateuploadedcertificate.Setup,
		volume.Setup,
		attachmentvolumeattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
