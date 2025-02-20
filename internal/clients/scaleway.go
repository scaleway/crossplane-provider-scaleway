/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/scaleway/crossplane-provider-scaleway/internal/version"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/scaleway/crossplane-provider-scaleway/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal scaleway credentials as JSON"

	keyAccessKey      = "access_key"
	keySecretKey      = "secret_key"
	keyProjectID      = "project_id"
	keyOrganizationID = "organization_id"
	keyRegion         = "region"
	keyZone           = "zone"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(tfversion, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: tfversion,
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

		scalewayCreds := map[string]string{}
		if err := json.Unmarshal(data, &scalewayCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{}
		for _, key := range []string{
			keyAccessKey,
			keySecretKey,
			keyProjectID,
			keyOrganizationID,
			keyRegion,
			keyZone,
		} {
			if scalewayCreds[key] != "" {
				ps.Configuration[key] = scalewayCreds[key]
			}
		}

		// Set the custom user agent
		userAgent := fmt.Sprintf("crossplane-provider-scaleway/%s", version.Version)
		err = os.Setenv("TF_APPEND_USER_AGENT", userAgent)
		if err != nil {
			return ps, errors.Wrap(err, "cannot set user agent")
		}

		return ps, nil
	}
}
