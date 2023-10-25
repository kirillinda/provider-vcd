package nsxtnetworkdhcpbinding

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_network_dhcp_binding", func(r *config.Resource) {

		r.ShortGroup = "nsxtnetworkdhcpbinding"
		r.References["mac_address"] = config.Reference{
			Type:      "github.com/kirillinda/provider-vcd/apis/vm/v1alpha1.VM",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("network.mac", true)`,
		}
		r.References["org_network_id"] = config.Reference{
			Type:      "github.com/kirillinda/provider-vcd/apis/vcdnetworkroutedv2/v1alpha1.RoutedV2",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id", true)`,
		}
		r.References["dhcp_v4_config.hostname"] = config.Reference{
			Type:      "github.com/kirillinda/provider-vcd/apis/vm/v1alpha1.VM",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("computer_name", true)`,
		}
	})
}
