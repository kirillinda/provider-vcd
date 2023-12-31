/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/kirillinda/provider-vcd/apis/independent/v1alpha1"
	v1alpha11 "github.com/kirillinda/provider-vcd/apis/vcdnetworkroutedv2/v1alpha1"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VM.
func (mg *VM) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Disk); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Disk[i3].Name),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.Disk[i3].NameRef,
			Selector:     mg.Spec.ForProvider.Disk[i3].NameSelector,
			To: reference.To{
				List:    &v1alpha1.DiskList{},
				Managed: &v1alpha1.Disk{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Disk[i3].Name")
		}
		mg.Spec.ForProvider.Disk[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Disk[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Network); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network[i3].Name),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.Network[i3].NameRef,
			Selector:     mg.Spec.ForProvider.Network[i3].NameSelector,
			To: reference.To{
				List:    &v1alpha11.RoutedV2List{},
				Managed: &v1alpha11.RoutedV2{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Network[i3].Name")
		}
		mg.Spec.ForProvider.Network[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Network[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}
