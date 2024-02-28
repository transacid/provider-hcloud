/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/transacid/provider-hcloud/config/firewall"
	"github.com/transacid/provider-hcloud/config/firewallAttachment"
	"github.com/transacid/provider-hcloud/config/floatingIp"
	"github.com/transacid/provider-hcloud/config/floatingIpAssignment"
	"github.com/transacid/provider-hcloud/config/loadBalancer"
	"github.com/transacid/provider-hcloud/config/loadBalancerNetwork"
	"github.com/transacid/provider-hcloud/config/loadBalancerService"
	"github.com/transacid/provider-hcloud/config/loadBalancerTarget"
	"github.com/transacid/provider-hcloud/config/managedCertificate"
	"github.com/transacid/provider-hcloud/config/network"
	"github.com/transacid/provider-hcloud/config/networkRoute"
	"github.com/transacid/provider-hcloud/config/networkSubnet"
	"github.com/transacid/provider-hcloud/config/placementGroup"
	"github.com/transacid/provider-hcloud/config/primaryIp"
	"github.com/transacid/provider-hcloud/config/rDns"
	"github.com/transacid/provider-hcloud/config/server"
	"github.com/transacid/provider-hcloud/config/serverNetwork"
	"github.com/transacid/provider-hcloud/config/snapshot"
	"github.com/transacid/provider-hcloud/config/sshKey"
	"github.com/transacid/provider-hcloud/config/uploadedCertificate"
	"github.com/transacid/provider-hcloud/config/volume"
	"github.com/transacid/provider-hcloud/config/volumeAttachment"
)

const (
	resourcePrefix = "hcloud"
	modulePath     = "github.com/transacid/provider-hcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("hcloud.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		firewall.Configure,
		firewallAttachment.Configure,
		floatingIp.Configure,
		floatingIpAssignment.Configure,
		loadBalancer.Configure,
		loadBalancerNetwork.Configure,
		loadBalancerService.Configure,
		loadBalancerTarget.Configure,
		managedCertificate.Configure,
		network.Configure,
		networkRoute.Configure,
		networkSubnet.Configure,
		placementGroup.Configure,
		primaryIp.Configure,
		rDns.Configure,
		server.Configure,
		serverNetwork.Configure,
		snapshot.Configure,
		sshKey.Configure,
		uploadedCertificate.Configure,
		volume.Configure,
		volumeAttachment.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
