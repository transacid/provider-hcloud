package managedCertificate

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_managed_certificate", func(r *config.Resource) {
		r.ShortGroup = "managedCertificate"

	})
}
