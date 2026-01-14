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

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"github.com/scaleway/crossplane-provider-scaleway/internal/version"
	"github.com/scaleway/scaleway-sdk-go/scw"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/scaleway/crossplane-provider-scaleway/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/scaleway/crossplane-provider-scaleway/apis/namespaced/v1beta1"
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

		// Resolve ProviderConfig based on the managed resource type
		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot resolve provider config")
		}

		// Load Scaleway config file + Env (Env > File)
		profile, err := resolveScwProfile(pcSpec)
		if err != nil {
			return ps, err
		}
		fillConfigFromProfile(ps.Configuration, profile)

		// Overlay Secret (if any). Secret > Env > File
		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
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

// resolveProviderConfig resolves the ProviderConfigSpec based on the managed resource type
func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

// resolveLegacy handles cluster-scoped MRs with legacy ProviderConfigReferencer
func resolveLegacy(ctx context.Context, crClient client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(crClient, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toNamespacedPCSpec(pc)
}

func toNamespacedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}

	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)

	return &mSpec, err
}

// resolveModern handles namespaced MRs with TypedProviderConfigReferencer
func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}

	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return nil, errors.New("resolved type is not a client.Object")
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}

	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		// For namespaced ProviderConfig, set secret namespace to MR's namespace
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}

	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return &pcSpec, nil
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

func resolveScwProfile(spec *namespacedv1beta1.ProviderConfigSpec) (*scw.Profile, error) {
	useFile := shouldUseScwConfig(spec)

	fileProf, err := readScwConfigProfile(spec, useFile)
	if err != nil {
		return nil, err
	}

	envProf := scw.LoadEnvProfile()
	if fileProf == nil {
		return envProf, nil
	}
	return scw.MergeProfiles(fileProf, envProf), nil
}

func readScwConfigProfile(spec *namespacedv1beta1.ProviderConfigSpec, useFile bool) (*scw.Profile, error) {
	if !useFile {
		return nil, nil
	}

	cfg, err := loadScwConfigFile(spec)
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, nil
	}

	return selectScwProfile(spec, cfg)
}

func loadScwConfigFile(spec *namespacedv1beta1.ProviderConfigSpec) (*scw.Config, error) {
	var cfg *scw.Config
	var err error

	if hasScwPath(spec) {
		cfg, err = scw.LoadConfigFromPath(*spec.Scw.Path)
	} else {
		cfg, err = scw.LoadConfig()
	}

	var notFound *scw.ConfigFileNotFoundError
	if err != nil && !errors.As(err, &notFound) {
		return nil, errors.Wrap(err, errLoadSCWConfig)
	}

	return cfg, nil
}

func selectScwProfile(spec *namespacedv1beta1.ProviderConfigSpec, cfg *scw.Config) (*scw.Profile, error) {
	if cfg == nil {
		return nil, nil
	}

	if hasScwProfile(spec) {
		prof, err := cfg.GetProfile(*spec.Scw.Profile)
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

func hasScwPath(spec *namespacedv1beta1.ProviderConfigSpec) bool {
	return spec.Scw != nil && spec.Scw.Path != nil && *spec.Scw.Path != ""
}

func hasScwProfile(spec *namespacedv1beta1.ProviderConfigSpec) bool {
	return spec.Scw != nil && spec.Scw.Profile != nil && *spec.Scw.Profile != ""
}

func shouldUseScwConfig(spec *namespacedv1beta1.ProviderConfigSpec) bool {
	if spec.Scw != nil && spec.Scw.UseScwConfig != nil {
		return *spec.Scw.UseScwConfig
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
