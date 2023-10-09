package vcdnetworkroutedv2

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vm", func(r *config.Resource) {

		r.ShortGroup = "vm"
	})
}
