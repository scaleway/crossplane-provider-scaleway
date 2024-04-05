// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/scaleway/provider-scaleway/apis/account/v1alpha1"
	v1alpha1applesilicon "github.com/scaleway/provider-scaleway/apis/applesilicon/v1alpha1"
	v1alpha1baremetal "github.com/scaleway/provider-scaleway/apis/baremetal/v1alpha1"
	v1alpha1cockpit "github.com/scaleway/provider-scaleway/apis/cockpit/v1alpha1"
	v1alpha1container "github.com/scaleway/provider-scaleway/apis/container/v1alpha1"
	v1alpha1domain "github.com/scaleway/provider-scaleway/apis/domain/v1alpha1"
	v1alpha1flexibleip "github.com/scaleway/provider-scaleway/apis/flexibleip/v1alpha1"
	v1alpha1function "github.com/scaleway/provider-scaleway/apis/function/v1alpha1"
	v1alpha1iam "github.com/scaleway/provider-scaleway/apis/iam/v1alpha1"
	v1alpha1instance "github.com/scaleway/provider-scaleway/apis/instance/v1alpha1"
	v1alpha1iot "github.com/scaleway/provider-scaleway/apis/iot/v1alpha1"
	v1alpha1k8s "github.com/scaleway/provider-scaleway/apis/k8s/v1alpha1"
	v1alpha1lb "github.com/scaleway/provider-scaleway/apis/lb/v1alpha1"
	v1alpha1object "github.com/scaleway/provider-scaleway/apis/object/v1alpha1"
	v1alpha1rdb "github.com/scaleway/provider-scaleway/apis/rdb/v1alpha1"
	v1alpha1redis "github.com/scaleway/provider-scaleway/apis/redis/v1alpha1"
	v1alpha1registry "github.com/scaleway/provider-scaleway/apis/registry/v1alpha1"
	v1alpha1tem "github.com/scaleway/provider-scaleway/apis/tem/v1alpha1"
	v1alpha1apis "github.com/scaleway/provider-scaleway/apis/v1alpha1"
	v1beta1 "github.com/scaleway/provider-scaleway/apis/v1beta1"
	v1alpha1vpc "github.com/scaleway/provider-scaleway/apis/vpc/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1applesilicon.SchemeBuilder.AddToScheme,
		v1alpha1baremetal.SchemeBuilder.AddToScheme,
		v1alpha1cockpit.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha1domain.SchemeBuilder.AddToScheme,
		v1alpha1flexibleip.SchemeBuilder.AddToScheme,
		v1alpha1function.SchemeBuilder.AddToScheme,
		v1alpha1iam.SchemeBuilder.AddToScheme,
		v1alpha1instance.SchemeBuilder.AddToScheme,
		v1alpha1iot.SchemeBuilder.AddToScheme,
		v1alpha1k8s.SchemeBuilder.AddToScheme,
		v1alpha1lb.SchemeBuilder.AddToScheme,
		v1alpha1object.SchemeBuilder.AddToScheme,
		v1alpha1rdb.SchemeBuilder.AddToScheme,
		v1alpha1redis.SchemeBuilder.AddToScheme,
		v1alpha1registry.SchemeBuilder.AddToScheme,
		v1alpha1tem.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1vpc.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
