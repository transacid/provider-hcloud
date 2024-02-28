package volume

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_volume", func(r *config.Resource) {
		r.ShortGroup = "volume"

		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
	})
}
