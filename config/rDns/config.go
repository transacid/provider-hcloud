package rDns

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_rdns", func(r *config.Resource) {
		r.ShortGroup = "rDns"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
		r.References["loadBalancer"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/loadBalancer/v1alpha1.Balancer",
		}
		r.References["floatingIp"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/floatingIp/v1alpha1.Ip",
		}
		r.References["primaryIp"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/primaryIp/v1alpha1.Ip",
		}
	})
}
