package serverNetwork

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_server_network", func(r *config.Resource) {
		r.ShortGroup = "serverNetwork"

		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
		r.References["network"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/network/v1alpha1.Network",
		}
		r.References["networkSubnet"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/networkSubnet/v1alpha1.Subnet",
		}
	})
}
