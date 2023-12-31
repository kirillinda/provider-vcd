/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/kirillinda/provider-vcd/apis/vcdnetworkroutedv2/v1alpha1"
	v1alpha1 "github.com/kirillinda/provider-vcd/apis/vm/v1alpha1"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this NetworkDHCPBinding.
func (mg *NetworkDHCPBinding) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DHCPV4Config); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DHCPV4Config[i3].Hostname),
			Extract:      resource.ExtractParamPath("computer_name", true),
			Reference:    mg.Spec.ForProvider.DHCPV4Config[i3].HostnameRef,
			Selector:     mg.Spec.ForProvider.DHCPV4Config[i3].HostnameSelector,
			To: reference.To{
				List:    &v1alpha1.VMList{},
				Managed: &v1alpha1.VM{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DHCPV4Config[i3].Hostname")
		}
		mg.Spec.ForProvider.DHCPV4Config[i3].Hostname = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DHCPV4Config[i3].HostnameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MacAddress),
		Extract:      resource.ExtractParamPath("network[0].mac", true),
		Reference:    mg.Spec.ForProvider.MacAddressRef,
		Selector:     mg.Spec.ForProvider.MacAddressSelector,
		To: reference.To{
			List:    &v1alpha1.VMList{},
			Managed: &v1alpha1.VM{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MacAddress")
	}
	mg.Spec.ForProvider.MacAddress = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MacAddressRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgNetworkID),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.ForProvider.OrgNetworkIDRef,
		Selector:     mg.Spec.ForProvider.OrgNetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.RoutedV2List{},
			Managed: &v1alpha11.RoutedV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgNetworkID")
	}
	mg.Spec.ForProvider.OrgNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrgNetworkIDRef = rsp.ResolvedReference

	return nil
}
