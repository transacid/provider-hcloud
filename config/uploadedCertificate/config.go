package uploadedCertificate

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_uploaded_certificate", func(r *config.Resource) {
		r.ShortGroup = "uploadedCertificate"

	})
}
