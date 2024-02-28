package floatingIp

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_floating_ip", func(r *config.Resource) {
		r.ShortGroup = "floatingIp"

	})
}
