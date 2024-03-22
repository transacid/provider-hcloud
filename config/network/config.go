package network

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network", func(r *config.Resource) {
		r.ShortGroup = "network"

	})
}
