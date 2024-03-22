package firewallAttachment

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_firewall_attachment", func(r *config.Resource) {
		r.ShortGroup = "firewallAttachment"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["firewall"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/firewall/v1alpha1.Firewall",
		}
		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
	})
}
