package rightsBundle

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_rights_bundle", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RightsBundle"
		r.Version = "v1alpha1"
	})
}