/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/kirillinda/provider-vcd/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal vcd credentials as JSON"
	vcd_user                = "vcd_user"
	vcd_password 			= "vcd_password"
	auth_type				= "auth_type"
	vcd_org					= "vcd_org"
	vcd_vdc					= "vcd_vdc"
	vcd_url					= "vcd_url"
	vcd_max_retry_timeout	= "vcd_max_retry_timeout"
	vcd_allow_unverified_ssl = "vcd_allow_unverified_ssl"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
			if v, ok := creds[vcd_user]; ok {
				ps.Configuration[vcd_user] = v
			  }
			if v, ok := creds[vcd_password]; ok {
				ps.Configuration[vcd_password] = v
			  }
			if v, ok := creds[auth_type]; ok {
				ps.Configuration[auth_type] = v
			  }
			if v, ok := creds[vcd_org]; ok {
				ps.Configuration[vcd_org] = v
			  }
			if v, ok := creds[vcd_vdc]; ok {
				ps.Configuration[vcd_vdc] = v
			  }
			if v, ok := creds[vcd_url]; ok {
				ps.Configuration[vcd_url] = v
			  }
			if v, ok := creds[vcd_max_retry_timeout]; ok {
				ps.Configuration[vcd_max_retry_timeout] = v
			  }
			if v, ok := creds[vcd_allow_unverified_ssl]; ok {
				ps.Configuration[vcd_allow_unverified_ssl] = v
			  }
		
		return ps, nil
	}
}