package sshKey

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "sshKey"

	})
}
