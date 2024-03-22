package server

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_server", func(r *config.Resource) {
		r.ShortGroup = "server"

		r.References["firewall"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/firewall/v1alpha1.Firewall",
		}
		// r.References["network"] = config.Reference{
		// 	Type: "github.com/transacid/provider-hcloud/apis/network/v1alpha1.Network",
		// }
		r.References["placementGroup"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/placementGroup/v1alpha1.Group",
		}
		r.References["networksubnet"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/networksubnet/v1alpha1.Subnet",
		}
		r.References["primaryIp"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/primaryIp/v1alpha1.Ip",
		}
	})
}
