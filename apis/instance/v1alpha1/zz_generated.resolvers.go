/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/scaleway/crossplane-provider-scaleway/apis/ipam/v1alpha1"
	v1alpha11 "github.com/scaleway/crossplane-provider-scaleway/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Image.
func (mg *Image) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RootVolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RootVolumeIDRef,
		Selector:     mg.Spec.ForProvider.RootVolumeIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RootVolumeID")
	}
	mg.Spec.ForProvider.RootVolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RootVolumeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RootVolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RootVolumeIDRef,
		Selector:     mg.Spec.InitProvider.RootVolumeIDSelector,
		To: reference.To{
			List:    &SnapshotList{},
			Managed: &Snapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RootVolumeID")
	}
	mg.Spec.InitProvider.RootVolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RootVolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrivateNIC.
func (mg *PrivateNIC) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IpamIPIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.IpamIPIdsRefs,
		Selector:      mg.Spec.ForProvider.IpamIPIdsSelector,
		To: reference.To{
			List:    &v1alpha1.IPList{},
			Managed: &v1alpha1.IP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IpamIPIds")
	}
	mg.Spec.ForProvider.IpamIPIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IpamIPIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PrivateNetworkIDRef,
		Selector:     mg.Spec.ForProvider.PrivateNetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.PrivateNetworkList{},
			Managed: &v1alpha11.PrivateNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateNetworkID")
	}
	mg.Spec.ForProvider.PrivateNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.IpamIPIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.IpamIPIdsRefs,
		Selector:      mg.Spec.InitProvider.IpamIPIdsSelector,
		To: reference.To{
			List:    &v1alpha1.IPList{},
			Managed: &v1alpha1.IP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IpamIPIds")
	}
	mg.Spec.InitProvider.IpamIPIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.IpamIPIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PrivateNetworkIDRef,
		Selector:     mg.Spec.InitProvider.PrivateNetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.PrivateNetworkList{},
			Managed: &v1alpha11.PrivateNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrivateNetworkID")
	}
	mg.Spec.InitProvider.PrivateNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrivateNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroupRule.
func (mg *SecurityGroupRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IPIDRef,
		Selector:     mg.Spec.ForProvider.IPIDSelector,
		To: reference.To{
			List:    &IPList{},
			Managed: &IP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IPID")
	}
	mg.Spec.ForProvider.IPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IPIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlacementGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PlacementGroupIDRef,
		Selector:     mg.Spec.ForProvider.PlacementGroupIDSelector,
		To: reference.To{
			List:    &PlacementGroupList{},
			Managed: &PlacementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlacementGroupID")
	}
	mg.Spec.ForProvider.PlacementGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlacementGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupID")
	}
	mg.Spec.ForProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IPIDRef,
		Selector:     mg.Spec.InitProvider.IPIDSelector,
		To: reference.To{
			List:    &IPList{},
			Managed: &IP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IPID")
	}
	mg.Spec.InitProvider.IPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IPIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PlacementGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.PlacementGroupIDRef,
		Selector:     mg.Spec.InitProvider.PlacementGroupIDSelector,
		To: reference.To{
			List:    &PlacementGroupList{},
			Managed: &PlacementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlacementGroupID")
	}
	mg.Spec.InitProvider.PlacementGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlacementGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SecurityGroupID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.SecurityGroupIDRef,
		Selector:     mg.Spec.InitProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupID")
	}
	mg.Spec.InitProvider.SecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SecurityGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Snapshot.
func (mg *Snapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.VolumeIDRef,
		Selector:     mg.Spec.InitProvider.VolumeIDSelector,
		To: reference.To{
			List:    &VolumeList{},
			Managed: &Volume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VolumeID")
	}
	mg.Spec.InitProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserData.
func (mg *UserData) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServerIDRef,
		Selector:     mg.Spec.InitProvider.ServerIDSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServerID")
	}
	mg.Spec.InitProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}
