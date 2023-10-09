/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/kirillinda/provider-vcd/config/independentdisk"
	"github.com/kirillinda/provider-vcd/config/nsxtfirewall"
	"github.com/kirillinda/provider-vcd/config/nsxtnetworkdhcp"
	"github.com/kirillinda/provider-vcd/config/nsxtnetworkdhcpbinding"
	"github.com/kirillinda/provider-vcd/config/vcdnetworkroutedv2"
	"github.com/kirillinda/provider-vcd/config/vm"
	"github.com/kirillinda/provider-vcd/config/vminternaldisk"
)

const (
	resourcePrefix = "vcd"
	modulePath     = "github.com/kirillinda/provider-vcd"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		vcdnetworkroutedv2.Configure,
		vminternaldisk.Configure,
		vm.Configure,
		nsxtnetworkdhcpbinding.Configure,
		nsxtnetworkdhcp.Configure,
		nsxtfirewall.Configure,
		independentdisk.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
