# Autoscaling resource identities

The provider exposes two different Terraform autoscaling resources:

| Terraform resource | Crossplane kind | Template reference |
| --- | --- | --- |
| `scaleway_autoscaling_instance_group` | `Group` | `autoscaling.Template` |
| `scaleway_autoscaling_group` | `AutoscalingGroup` | `instance.Template` |

Both kinds use `v1alpha1` in `autoscaling.scaleway.upbound.io` for cluster-scoped
resources and `autoscaling.scaleway.m.upbound.io` for namespaced resources.

## Existing resources

The published `Group` identity remains attached to
`scaleway_autoscaling_instance_group`. Its `capacity` and `loadBalancer` fields,
controller, external-name handling, and the references from `Policy` are
preserved. The existing `Group` CRDs are unchanged by regeneration, so existing
Kubernetes objects require no CRD migration or recreation.

The generated Go types `GroupInitParameters_2`, `GroupObservation_2`, and
`GroupParameters_2` lose the collision suffix `_2`. Go consumers using those
type names must use their unsuffixed equivalents; their serialized fields are
unchanged.

## Newly supported resource

`AutoscalingGroup` has separate CRDs, controllers, and examples. It exposes the
`scalingPolicy` and `loadBalancerConfiguration` fields of
`scaleway_autoscaling_group` without replacing the existing `Group` contract.

Changing an existing manifest from `Group` to `AutoscalingGroup` is not a
conversion: the two kinds manage different Terraform resource types and API
contracts. This change performs no object or Terraform-state migration between
them.
