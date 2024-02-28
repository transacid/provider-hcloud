package networkSubnet

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "networkSubnet"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["network"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/network/v1alpha1.Network",
		}
	})
}
