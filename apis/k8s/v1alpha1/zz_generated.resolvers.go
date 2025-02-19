/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/scaleway/crossplane-provider-scaleway/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PrivateNetworkIDRef,
		Selector:     mg.Spec.ForProvider.PrivateNetworkIDSelector,
		To: reference.To{
			List:    &v1alpha1.PrivateNetworkList{},
			Managed: &v1alpha1.PrivateNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateNetworkID")
	}
	mg.Spec.ForProvider.PrivateNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PrivateNetworkIDRef,
		Selector:     mg.Spec.InitProvider.PrivateNetworkIDSelector,
		To: reference.To{
			List:    &v1alpha1.PrivateNetworkList{},
			Managed: &v1alpha1.PrivateNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrivateNetworkID")
	}
	mg.Spec.InitProvider.PrivateNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrivateNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}
