package loadBalancerService

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer_service", func(r *config.Resource) {
		r.ShortGroup = "loadBalancerService"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["loadbalancer"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/loadbalancer/v1alpha1.Balancer",
		}
	})
}
