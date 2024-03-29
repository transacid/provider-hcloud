/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"null_resource":                 config.IdentifierFromProvider,
	"hcloud_firewall":               config.IdentifierFromProvider,
	"hcloud_firewall_attachment":    config.IdentifierFromProvider,
	"hcloud_floating_ip":            config.IdentifierFromProvider,
	"hcloud_floating_ip_assignment": config.IdentifierFromProvider,
	"hcloud_load_balancer":          config.IdentifierFromProvider,
	"hcloud_load_balancer_network":  config.IdentifierFromProvider,
	"hcloud_load_balancer_service":  config.IdentifierFromProvider,
	"hcloud_load_balancer_target":   config.IdentifierFromProvider,
	"hcloud_managed_certificate":    config.IdentifierFromProvider,
	"hcloud_network":                config.IdentifierFromProvider,
	"hcloud_network_route":          config.IdentifierFromProvider,
	"hcloud_network_subnet":         config.IdentifierFromProvider,
	"hcloud_placement_group":        config.IdentifierFromProvider,
	"hcloud_primary_ip":             config.IdentifierFromProvider,
	"hcloud_rdns":                   config.IdentifierFromProvider,
	"hcloud_server":                 config.IdentifierFromProvider,
	"hcloud_server_network":         config.IdentifierFromProvider,
	"hcloud_snapshot":               config.IdentifierFromProvider,
	"hcloud_ssh_key":                config.IdentifierFromProvider,
	"hcloud_uploaded_certificate":   config.IdentifierFromProvider,
	"hcloud_volume":                 config.IdentifierFromProvider,
	"hcloud_volume_attachment":      config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
