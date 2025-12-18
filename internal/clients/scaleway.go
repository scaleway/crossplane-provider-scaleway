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
		profile, err := resolveScwProfile(pc)
		if err != nil {
			return ps, err
		}
		fillConfigFromProfile(ps.Configuration, profile)

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

func resolveScwProfile(pc *v1beta1.ProviderConfig) (*scw.Profile, error) {
	useFile := shouldUseScwConfig(pc)

	fileProf, err := readScwConfigProfile(pc, useFile)
	if err != nil {
		return nil, err
	}

	envProf := scw.LoadEnvProfile()
	if fileProf == nil {
		return envProf, nil
	}
	return scw.MergeProfiles(fileProf, envProf), nil
}

func readScwConfigProfile(pc *v1beta1.ProviderConfig, useFile bool) (*scw.Profile, error) {
	if !useFile {
		return nil, nil
	}

	cfg, err := loadScwConfigFile(pc)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, nil
	}

	return selectScwProfile(pc, cfg)
}

func loadScwConfigFile(pc *v1beta1.ProviderConfig) (*scw.Config, error) {
	var cfg *scw.Config
	var err error

	if hasScwPath(pc) {
		cfg, err = scw.LoadConfigFromPath(*pc.Spec.Scw.Path)
	} else {
		cfg, err = scw.LoadConfig()
	}

	var notFound *scw.ConfigFileNotFoundError
	if err != nil && !errors.As(err, &notFound) {
		return nil, errors.Wrap(err, errLoadSCWConfig)
	}

	return cfg, nil
}

func selectScwProfile(pc *v1beta1.ProviderConfig, cfg *scw.Config) (*scw.Profile, error) {
	if cfg == nil {
		return nil, nil
	}

	if hasScwProfile(pc) {
		prof, err := cfg.GetProfile(*pc.Spec.Scw.Profile)
		if err != nil {
			return nil, errors.Wrap(err, errGetSCWProfile)
		}
		return prof, nil
	}

	prof, err := cfg.GetActiveProfile()
	if err != nil {
		return nil, errors.Wrap(err, errGetSCWProfile)
	}
	return prof, nil
}

func hasScwPath(pc *v1beta1.ProviderConfig) bool {
	return pc.Spec.Scw != nil && pc.Spec.Scw.Path != nil && *pc.Spec.Scw.Path != ""
}

func hasScwProfile(pc *v1beta1.ProviderConfig) bool {
	return pc.Spec.Scw != nil && pc.Spec.Scw.Profile != nil && *pc.Spec.Scw.Profile != ""
}

func shouldUseScwConfig(pc *v1beta1.ProviderConfig) bool {
	if pc.Spec.Scw != nil && pc.Spec.Scw.UseScwConfig != nil {
		return *pc.Spec.Scw.UseScwConfig
	}
	return true
}

// fillConfigFromProfile copies non-empty profile fields into cfg
func fillConfigFromProfile(cfg map[string]any, p *scw.Profile) {
	if cfg == nil || p == nil {
		return
	}

	assign := func(key string, val *string) {
		if val != nil && *val != "" {
			cfg[key] = *val
		}
	}

	assign(keyAccessKey, p.AccessKey)
	assign(keySecretKey, p.SecretKey)
	assign(keyProjectID, p.DefaultProjectID)
	assign(keyOrganizationID, p.DefaultOrganizationID)
	assign(keyRegion, p.DefaultRegion)
	assign(keyZone, p.DefaultZone)
	assign(keyAPIURL, p.APIURL)

	if p.Insecure != nil {
		cfg[keyInsecure] = *p.Insecure
	}
}
