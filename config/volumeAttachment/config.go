package volumeAttachment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_volume_attachment", func(r *config.Resource) {
		r.ShortGroup = "volumeAttachment"

		r.References["server"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/server/v1alpha1.Server",
		}
		r.References["volume"] = config.Reference{
			Type: "github.com/transacid/provider-hcloud/apis/volume/v1alpha1.Volume",
		}
	})
}
