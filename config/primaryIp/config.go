package primaryIp

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_primary_ip", func(r *config.Resource) {
		r.ShortGroup = "primaryIp"

	})
}
