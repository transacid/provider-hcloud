package placementGroup

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_placement_group", func(r *config.Resource) {
		r.ShortGroup = "placementGroup"

	})
}
