package vm

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vm", func(r *config.Resource) {

		r.ShortGroup = "vm"
		r.References["network.name"] = config.Reference{
			Type: "github.com/kirillinda/provider-vcd/apis/vcdnetworkroutedv2/v1alpha1.RoutedV2",
		}
		r.References["disk.name"] = config.Reference{
			Type: "github.com/kirillinda/provider-vcd/apis/independent/v1alpha1.Disk",
		}
	})
}
