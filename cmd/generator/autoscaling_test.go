/*
Copyright 2026 The Crossplane Authors.
*/

package main

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	providerconfig "github.com/scaleway/crossplane-provider-scaleway/config"
)

func TestAutoscalingResourceIdentities(t *testing.T) {
	providers := []struct {
		name string
		load func() *ujconfig.Provider
	}{
		{name: "cluster", load: providerconfig.GetProvider},
		{name: "namespaced", load: providerconfig.GetProviderNamespaced},
	}
	for _, provider := range providers {
		t.Run(provider.name, func(t *testing.T) {
			p := provider.load()
			legacy := p.Resources["scaleway_autoscaling_instance_group"]
			current := p.Resources["scaleway_autoscaling_group"]
			if legacy == nil || current == nil {
				t.Fatal("both autoscaling resources must remain enabled")
			}
			if legacy.ShortGroup == current.ShortGroup && legacy.Kind == current.Kind {
				t.Fatalf("autoscaling resources collide at %s/%s", legacy.ShortGroup, legacy.Kind)
			}
			if legacy.ShortGroup != "autoscaling" || legacy.Kind != "Group" {
				t.Errorf("published instance group identity changed: %s/%s", legacy.ShortGroup, legacy.Kind)
			}
			if current.ShortGroup != "autoscaling" || current.Kind != "AutoscalingGroup" {
				t.Errorf("new group identity = %s/%s, want autoscaling/AutoscalingGroup", current.ShortGroup, current.Kind)
			}
			policy := p.Resources["scaleway_autoscaling_instance_policy"]
			if policy == nil || policy.References["instance_group_id"].TerraformName != "scaleway_autoscaling_instance_group" {
				t.Error("existing autoscaling policies must still reference the published instance group")
			}
		})
	}
}
