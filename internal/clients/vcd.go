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
	user                    = "user"
	password                = "password"
	auth_type               = "auth_type"
	org                     = "org"
	vdc                     = "vdc"
	url                     = "url"
	max_retry_timeout       = "max_retry_timeout"
	allow_unverified_ssl    = "allow_unverified_ssl"
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

		ps.Configuration = map[string]any{}
		if v, ok := creds[user]; ok {
			ps.Configuration[user] = v
		}
		if v, ok := creds[password]; ok {
			ps.Configuration[password] = v
		}
		if v, ok := creds[auth_type]; ok {
			ps.Configuration[auth_type] = v
		}
		if v, ok := creds[org]; ok {
			ps.Configuration[org] = v
		}
		if v, ok := creds[vdc]; ok {
			ps.Configuration[vdc] = v
		}
		if v, ok := creds[url]; ok {
			ps.Configuration[url] = v
		}
		if v, ok := creds[max_retry_timeout]; ok {
			ps.Configuration[max_retry_timeout] = v
		}
		if v, ok := creds[allow_unverified_ssl]; ok {
			ps.Configuration[allow_unverified_ssl] = v
		}

		return ps, nil
	}
}
