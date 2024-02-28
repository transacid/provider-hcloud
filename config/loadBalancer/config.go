package loadBalancer

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer", func(r *config.Resource) {
		r.ShortGroup = "loadBalancer"

	})
}
