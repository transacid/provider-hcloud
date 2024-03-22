package floatingIpAssignment

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_floating_ip_assignment", func(r *config.Resource) {
		r.ShortGroup = "floatingIpAssignment"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["floatingIp"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/floatingIp/v1alpha1.IP",
		}
		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
	})
}
