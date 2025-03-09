// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attachment "github.com/transacid/provider-hcloud/internal/controller/firewall/attachment"
	firewall "github.com/transacid/provider-hcloud/internal/controller/firewall/firewall"
	ip "github.com/transacid/provider-hcloud/internal/controller/floatingip/ip"
	ipassignment "github.com/transacid/provider-hcloud/internal/controller/floatingip/ipassignment"
	balancer "github.com/transacid/provider-hcloud/internal/controller/loadbalancer/balancer"
	balancernetwork "github.com/transacid/provider-hcloud/internal/controller/loadbalancer/balancernetwork"
	balancerservice "github.com/transacid/provider-hcloud/internal/controller/loadbalancer/balancerservice"
	balancertarget "github.com/transacid/provider-hcloud/internal/controller/loadbalancer/balancertarget"
	certificate "github.com/transacid/provider-hcloud/internal/controller/managed/certificate"
	network "github.com/transacid/provider-hcloud/internal/controller/network/network"
	route "github.com/transacid/provider-hcloud/internal/controller/network/route"
	subnet "github.com/transacid/provider-hcloud/internal/controller/network/subnet"
	group "github.com/transacid/provider-hcloud/internal/controller/placementgroup/group"
	ipprimaryip "github.com/transacid/provider-hcloud/internal/controller/primaryip/ip"
	providerconfig "github.com/transacid/provider-hcloud/internal/controller/providerconfig"
	rdns "github.com/transacid/provider-hcloud/internal/controller/rdns/rdns"
	networkserver "github.com/transacid/provider-hcloud/internal/controller/server/network"
	server "github.com/transacid/provider-hcloud/internal/controller/server/server"
	snapshot "github.com/transacid/provider-hcloud/internal/controller/snapshot/snapshot"
	key "github.com/transacid/provider-hcloud/internal/controller/sshkey/key"
	certificateuploaded "github.com/transacid/provider-hcloud/internal/controller/uploaded/certificate"
	attachmentvolume "github.com/transacid/provider-hcloud/internal/controller/volume/attachment"
	volume "github.com/transacid/provider-hcloud/internal/controller/volume/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		firewall.Setup,
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
		networkserver.Setup,
		server.Setup,
		snapshot.Setup,
		key.Setup,
		certificateuploaded.Setup,
		attachmentvolume.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
