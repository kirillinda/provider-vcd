package nsxtfirewall

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_firewall", func(r *config.Resource) {

		r.ShortGroup = "nsxtfirewall"
	})
}
