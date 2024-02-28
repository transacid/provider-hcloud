package loadBalancerTarget

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer_target", func(r *config.Resource) {
		r.ShortGroup = "loadBalancerTarget"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
		r.References["loadbalancer"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/loadbalancer/v1alpha1.Balancer",
		}
	})
}
