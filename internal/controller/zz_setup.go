/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	disk "github.com/kirillinda/provider-vcd/internal/controller/independent/disk"
	firewall "github.com/kirillinda/provider-vcd/internal/controller/nsxtfirewall/firewall"
	networkdhcp "github.com/kirillinda/provider-vcd/internal/controller/nsxtnetworkdhcp/networkdhcp"
	networkdhcpbinding "github.com/kirillinda/provider-vcd/internal/controller/nsxtnetworkdhcpbinding/networkdhcpbinding"
	providerconfig "github.com/kirillinda/provider-vcd/internal/controller/providerconfig"
	routedv2 "github.com/kirillinda/provider-vcd/internal/controller/vcdnetworkroutedv2/routedv2"
	vm "github.com/kirillinda/provider-vcd/internal/controller/vm/vm"
	internaldisk "github.com/kirillinda/provider-vcd/internal/controller/vminternaldisk/internaldisk"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		disk.Setup,
		firewall.Setup,
		networkdhcp.Setup,
		networkdhcpbinding.Setup,
		providerconfig.Setup,
		routedv2.Setup,
		vm.Setup,
		internaldisk.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
