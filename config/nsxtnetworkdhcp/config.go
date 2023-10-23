package nsxtnetworkdhcp

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_network_dhcp", func(r *config.Resource) {

		r.ShortGroup = "nsxtnetworkdhcp"
		r.References["vcdnetworkroutedv2"] = config.Reference{
			Type: "github.com/myorg/provider-github/apis/vcdnetworkroutedv2/v1alpha1",
		}
	})
}
