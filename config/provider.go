/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
)

const (
	resourcePrefix = "hcloud"
	modulePath     = "github.com/transacid/provider-hcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// "null_resource": config.IdentifierFromProvider,
	"hcloud_uploaded_certificate":   ujconfig.IdentifierFromProvider,
	"hcloud_managed_certificate":    ujconfig.IdentifierFromProvider,
	"hcloud_firewall":               ujconfig.IdentifierFromProvider,
	"hcloud_firewall_attachment":    ujconfig.IdentifierFromProvider,
	"hcloud_floating_ip":            ujconfig.IdentifierFromProvider,
	"hcloud_floating_ip_assignment": ujconfig.IdentifierFromProvider,
	"hcloud_load_balancer":          ujconfig.IdentifierFromProvider,
	"hcloud_load_balancer_network":  ujconfig.IdentifierFromProvider,
	"hcloud_load_balancer_service":  ujconfig.IdentifierFromProvider,
	"hcloud_load_balancer_target":   ujconfig.IdentifierFromProvider,
	"hcloud_network":                ujconfig.IdentifierFromProvider,
	"hcloud_network_route":          ujconfig.IdentifierFromProvider,
	"hcloud_network_subnet":         ujconfig.IdentifierFromProvider,
	"hcloud_placement_group":        ujconfig.IdentifierFromProvider,
	"hcloud_primary_ip":             ujconfig.IdentifierFromProvider,
	"hcloud_rdns":                   ujconfig.IdentifierFromProvider,
	"hcloud_server":                 ujconfig.IdentifierFromProvider,
	"hcloud_server_network":         ujconfig.IdentifierFromProvider,
	"hcloud_snapshot":               ujconfig.IdentifierFromProvider,
	"hcloud_ssh_key":                ujconfig.IdentifierFromProvider,
	"hcloud_volume":                 ujconfig.IdentifierFromProvider,
	"hcloud_volume_attachment":      ujconfig.IdentifierFromProvider,
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("hcloud.crossplane.io"),
		ujconfig.WithShortName("hc"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))
	pc.AddResourceConfigurator("hcloud_firewall", func(r *ujconfig.Resource) {
		r.ShortGroup = "firewall"
		r.References["apply_to.server"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_uploaded_certificate", func(r *ujconfig.Resource) {})
	pc.AddResourceConfigurator("hcloud_managed_certificate", func(r *ujconfig.Resource) {})
	pc.AddResourceConfigurator("hcloud_firewall_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = "firewall"
		r.References["firefall_id"] = ujconfig.Reference{
			TerraformName: "hcloud_firewall",
		}
		r.References["server_ids"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_floating_ip", func(r *ujconfig.Resource) {
		r.ShortGroup = "floatingip"
		r.References["server_ids"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_floating_ip_assignment", func(r *ujconfig.Resource) {
		r.ShortGroup = "floatingip"
		r.References["server_ids"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
		r.References["floating_ip_id"] = ujconfig.Reference{
			TerraformName: "hcloud_floating_ip",
		}
	})
	pc.AddResourceConfigurator("hcloud_load_balancer", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
	})
	pc.AddResourceConfigurator("hcloud_load_balancer_network", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["network_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network",
		}
		r.References["load_balancer_id"] = ujconfig.Reference{
			TerraformName: "hcloud_load_balancer",
		}
	})
	pc.AddResourceConfigurator("hcloud_load_balancer_service", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["load_balancer_id"] = ujconfig.Reference{
			TerraformName: "hcloud_load_balancer",
		}
	})
	pc.AddResourceConfigurator("hcloud_load_balancer_target", func(r *ujconfig.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["server_ids"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_network", func(r *ujconfig.Resource) {
		r.ShortGroup = "network"
	})
	pc.AddResourceConfigurator("hcloud_network_route", func(r *ujconfig.Resource) {
		r.ShortGroup = "network"
		r.References["network_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network",
		}
	})
	pc.AddResourceConfigurator("hcloud_network_subnet", func(r *ujconfig.Resource) {
		r.ShortGroup = "network"
		r.References["network_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network",
		}
	})
	pc.AddResourceConfigurator("hcloud_placement_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "placementgroup"
	})
	pc.AddResourceConfigurator("hcloud_primary_ip", func(r *ujconfig.Resource) {
		r.ShortGroup = "primaryip"
	})
	pc.AddResourceConfigurator("hcloud_rdns", func(r *ujconfig.Resource) {
		r.ShortGroup = "rdns"
	})
	pc.AddResourceConfigurator("hcloud_server", func(r *ujconfig.Resource) {
		r.ShortGroup = "server"
		r.References["firwall_ids"] = ujconfig.Reference{
			TerraformName: "hcloud_firewall",
		}
		r.References["network.network_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network",
		}
		r.References["placement_group_id"] = ujconfig.Reference{
			TerraformName: "hcloud_placement_group",
		}
	})
	pc.AddResourceConfigurator("hcloud_server_network", func(r *ujconfig.Resource) {
		r.ShortGroup = "server"
		r.References["server_id"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
		r.References["network_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network",
		}
		r.References["subnet_id"] = ujconfig.Reference{
			TerraformName: "hcloud_network_subnet",
		}
	})
	pc.AddResourceConfigurator("hcloud_snapshot", func(r *ujconfig.Resource) {
		r.ShortGroup = "snapshot"
		r.References["server_id"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_ssh_key", func(r *ujconfig.Resource) {
		r.ShortGroup = "sshkey"
	})
	pc.AddResourceConfigurator("hcloud_volume", func(r *ujconfig.Resource) {
		r.ShortGroup = "volume"
		r.References["server_id"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
	})
	pc.AddResourceConfigurator("hcloud_volume_attachment", func(r *ujconfig.Resource) {
		r.ShortGroup = "volume"
		r.References["server_id"] = ujconfig.Reference{
			TerraformName: "hcloud_server",
		}
		r.References["volume_id"] = ujconfig.Reference{
			TerraformName: "hcloud_volume",
		}
	})

	pc.ConfigureResources()
	return pc
}
