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
	"github.com/scaleway/scaleway-sdk-go/scw"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/scaleway/crossplane-provider-scaleway/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errUnmarshalCredentials = "cannot unmarshal scaleway credentials as JSON"
	errLoadSCWConfig        = "cannot load SCW config file"
	errGetSCWProfile        = "cannot get SCW profile from config"

	keyAccessKey      = "access_key"
	keySecretKey      = "secret_key"
	keyProjectID      = "project_id"
	keyOrganizationID = "organization_id"
	keyRegion         = "region"
	keyZone           = "zone"
	keyAPIURL         = "api_url"
	keyInsecure       = "insecure"
)

func mapProfileToConfig(cfg map[string]any, p *scw.Profile) {
	if p == nil {
		return
	}

	if p.AccessKey != nil && *p.AccessKey != "" {
		cfg[keyAccessKey] = *p.AccessKey
	}
	if p.SecretKey != nil && *p.SecretKey != "" {
		cfg[keySecretKey] = *p.SecretKey
	}
	if p.DefaultProjectID != nil && *p.DefaultProjectID != "" {
		cfg[keyProjectID] = *p.DefaultProjectID
	}
	if p.DefaultOrganizationID != nil && *p.DefaultOrganizationID != "" {
		cfg[keyOrganizationID] = *p.DefaultOrganizationID
	}
	if p.DefaultRegion != nil && *p.DefaultRegion != "" {
		cfg[keyRegion] = *p.DefaultRegion
	}
	if p.DefaultZone != nil && *p.DefaultZone != "" {
		cfg[keyZone] = *p.DefaultZone
	}
	if p.APIURL != nil && *p.APIURL != "" {
		cfg[keyAPIURL] = *p.APIURL
	}
	if p.Insecure != nil {
		cfg[keyInsecure] = *p.Insecure
	}
}

// overlayCredentials overlays non-empty Secret values over existing config
func overlayCredentials(cfg map[string]any, vals map[string]string) {
	for _, k := range []string{
		keyAccessKey, keySecretKey, keyProjectID, keyOrganizationID,
		keyRegion, keyZone, keyAPIURL, keyInsecure,
	} {
		if v, ok := vals[k]; ok && v != "" {
			cfg[k] = v
		}
	}
}

// loadSCWProfile loads and merges Scaleway config + Env
func loadSCWProfile(pc *v1beta1.ProviderConfig) (*scw.Profile, error) {
	useFile := true
	if pc.Spec.Scw != nil && pc.Spec.Scw.UseScwConfig != nil {
		useFile = *pc.Spec.Scw.UseScwConfig
	}

	var fileProf *scw.Profile
	var err error

	if useFile {
		var cfg *scw.Config
		if pc.Spec.Scw != nil && pc.Spec.Scw.Path != nil && *pc.Spec.Scw.Path != "" {
			cfg, err = scw.LoadConfigFromPath(*pc.Spec.Scw.Path)
		} else {
			cfg, err = scw.LoadConfig()
		}

		var notFound *scw.ConfigFileNotFoundError
		if err != nil && !errors.As(err, &notFound) {
			return nil, errors.Wrap(err, errLoadSCWConfig)
		}

		if cfg != nil {
			if pc.Spec.Scw != nil && pc.Spec.Scw.Profile != nil && *pc.Spec.Scw.Profile != "" {
				fileProf, err = cfg.GetProfile(*pc.Spec.Scw.Profile)
			} else {
				fileProf, err = cfg.GetActiveProfile()
			}
			if err != nil && !errors.As(err, &notFound) {
				return nil, errors.Wrap(err, errGetSCWProfile)
			}
		}
	}

	envProf := scw.LoadEnvProfile()
	return scw.MergeProfiles(fileProf, envProf), nil
}

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
			Configuration: map[string]any{},
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

		// Load Scaleway config file + Env (Env > File)
		profile, err := loadSCWProfile(pc)
		if err != nil {
			return ps, err
		}
		mapProfileToConfig(ps.Configuration, profile)

		// Overlay Secret (if any). Secret > Env > File
		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err == nil && len(data) > 0 {
			vals := map[string]string{}
			if err := json.Unmarshal(data, &vals); err != nil {
				return ps, errors.Wrap(err, errUnmarshalCredentials)
			}
			overlayCredentials(ps.Configuration, vals)
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
