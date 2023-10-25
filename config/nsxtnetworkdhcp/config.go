package nsxtnetworkdhcp

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_network_dhcp", func(r *config.Resource) {

		r.ShortGroup = "nsxtnetworkdhcp"
		r.References["org_network_id"] = config.Reference{
			Type:      "github.com/kirillinda/provider-vcd/apis/vcdnetworkroutedv2/v1alpha1.RoutedV2",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id", true)`,
		}
	})
}
