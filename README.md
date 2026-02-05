# Crossplane Provider Scaleway

`crossplane-provider-scaleway` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Scaleway](https://www.scaleway.com/).

This provider supports **Crossplane v2** and exposes both cluster-scoped and namespaced managed resources.

Complete the following steps to:
* Install Crossplane v2 into your Kubernetes cluster.
* Install the `Provider` and apply a `ClusterProviderConfig` (or namespaced `ProviderConfig`) for managed resources.
* Create a *managed resource* in Scaleway with Kubernetes.

## Prerequisites

To perform the following steps, make sure you have:
* **Crossplane v2** installed in your cluster
* Your Scaleway [credentials](https://console.scaleway.com/project/credentials)
* A Kubernetes cluster with permissions to create pods and secrets
* A host with `kubectl` installed and configured to access the Kubernetes cluster

## Getting started

You can run each command individually or copy them to a local file to avoid issues related to running commands in a terminal.

_Note:_ All commands use the current `kubeconfig` context and configuration.

### Install Crossplane

Install Crossplane v2 using the official Helm chart as described in [Install Crossplane (v2.0)](https://docs.crossplane.io/v2.0/get-started/install/):

```shell
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update
helm install crossplane \
  --namespace crossplane-system \
  --create-namespace \
  crossplane-stable/crossplane
```

Verify with `kubectl get pods -n crossplane-system`.

### Install the provider

1. Install the provider into the Kubernetes cluster with a Kubernetes configuration file.

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-scaleway
spec:
  package: xpkg.upbound.io/scaleway/provider-scaleway:v0.6.0
EOF
```

2. Run `kubectl get providers` to verify the installed provider. The `INSTALLED` value should return as `True`. 

_Note:_  The procedure may take up to 5 minutes for `HEALTHY` to report true.

You should get an output similar to the following one, providing details about the provider.

```shell
$ kubectl get provider
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway   True        True   xpkg.upbound.io/scaleway/provider-scaleway:v0.6.0        11s
```

If there are any issues during the process of downloading and installing the provider, the `INSTALLED` field will return as empty. In that case, run `kubectl describe providers` to get more information.

```shell
$ kubectl get providers
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway                      xpkg.upbound.io/scaleway/provider-scaleway:v0.6.0      76s
```

### Create a Kubernetes secret resource for Scaleway

The provider requires credentials to create and manage Scaleway resources.

1. In a new folder, create a `secret.yaml` file. 

Modify the values in the example according to your needs, using the information in the Configuration reference table to help.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "access_key": "SCWXXXXXXXXXXXXXXXXX",
      "secret_key": "11111111-1111-1111-1111-111111111111",
      "project_id": "11111111-1111-1111-1111-111111111111",
      "organization_id": "11111111-1111-1111-1111-111111111111",
      "region": "fr-par",
      "zone": "fr-par-1"
    }
```

#### Configuration reference

| Provider Argument | Description                                                                                                                                                      |
|-------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `access_key`      | [Scaleway access key](https://console.scaleway.com/iam/api-keys)                                                                                                 |
| `secret_key`      | [Scaleway secret key](https://console.scaleway.com/iam/api-keys)                                                                                                 |
| `project_id`      | The [project ID](https://console.scaleway.com/project/settings) that will be used as default value for project-scoped resources.                                 |
| `organization_id` | The [organization ID](https://console.scaleway.com/organization/settings) that will be used as default value for organization-scoped resources.                  |
| `region`          | The [region](https://developers.scaleway.com/en/quickstart/#region-and-zone)  that will be used as default value for all resources. (`fr-par` if none specified) |
| `zone`            | The [zone](https://developers.scaleway.com/en/quickstart/#region-and-zone) that will be used as default value for all resources. (`fr-par-1` if none specified)  |
| `api_url`         | The URL of the API                                                                                                                                               |

### Create a ProviderConfig

Managed resources in this provider use the **namespaced** API group (`*.scaleway.m.upbound.io`). You can attach credentials in either of these ways:

- **ClusterProviderConfig**: one cluster-wide config, reference it from any namespace.
- **ProviderConfig**: per-namespace config, create one in each namespace where you create resources.

The following example creates a **ClusterProviderConfig** so the provider can use your Scaleway credentials from any namespace.

1. Create a `ClusterProviderConfig` Kubernetes configuration file to attach your Scaleway credentials to the previously installed provider.

Modify the values in the example according to your needs. Refer to the configuration reference information to understand the requested values.

```yaml
apiVersion: scaleway.m.upbound.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```

2. Run `kubectl apply -f your-folder/` to apply this configuration with the secret.

3. Run `kubectl describe clusterproviderconfigs.scaleway.m.upbound.io` to verify the `ClusterProviderConfig`.

For per-namespace credentials, create a `ProviderConfig` (`apiVersion: scaleway.m.upbound.io/v1beta1`) in the same namespace as your resources and add `metadata.namespace` to it. Legacy cluster-scoped resources (`*.scaleway.upbound.io`) use a cluster-scoped `ProviderConfig` from `scaleway.upbound.io/v1beta1`.

#### Configuration reference

The `spec.secretRef` describes the parameters of the secret to use.
* `namespace` is the Kubernetes namespace the secret is in.
* `name` is the name of the Kubernetes `secret` object.
* `key` is the `Data` field from `kubectl describe secret`.

### SCW config support

This provider can read the standard SCW config file (`~/.config/scw/config.yaml`) and environment variables. Precedence is:

1. ProviderConfig / ClusterProviderConfig credentials
2. Environment variables (`SCW_*`)
3. SCW config file

You can control behavior in `spec.scw` of your ProviderConfig or ClusterProviderConfig:

```yaml
spec:
  scw:
    useScwConfig: true
    # path: /home/me/.config/scw/config.yaml
    # profile: myProfile
```
### Create a managed resource

1. Create a managed resource to see if the provider is properly functioning.

The following example creates a Scaleway Object Storage bucket using the namespaced API. The resource lives in a namespace and references the `ClusterProviderConfig` named `default`. To write connection details (e.g. credentials) to a Kubernetes Secret, use `spec.writeConnectionSecretToRef` on the managed resource.

```yaml
apiVersion: object.scaleway.m.upbound.io/v1alpha1
kind: Bucket
metadata:
  name: object-bucket
  namespace: crossplane-system
spec:
  forProvider:
    name: crossplane-object-bucket
  providerConfigRef:
    name: default
```

2. Run `kubectl get buckets -n crossplane-system` to get details on the bucket's creation.

You should get an output similar to the following one, providing details about the bucket.

```shell
$ kubectl get buckets -n crossplane-system
NAME                           READY   SYNCED   EXTERNAL-NAME                     AGE
object-bucket                  True    True     fr-par/crossplane-object-bucket   9s
```

The bucket is successfully created when both the values for `READY` and `SYNCED` are `True`.

3. If there are any issues during the bucket creation process, the `READY` and/or `SYNCED` fields will return as empty. In that case, run `kubectl describe bucket -n crossplane-system object-bucket` to get more information.

### Delete the managed resource

1. Run `kubectl delete -f` (with the same `Bucket` file) to remove the managed resource.

2. Run `kubectl get buckets -n crossplane-system` to verify whether the bucket was properly removed.

You should get an output similar to this, providing details about the status of the bucket.

```shell
$ kubectl delete -f bucket.yml
bucket.object.scaleway.m.upbound.io "object-bucket" deleted

$ kubectl get buckets -n crossplane-system
No resources found
```

## Developing

- To run code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

- To run against a Kubernetes cluster:

```console
make run
```

- To build, push, and install:

```console
make all
```

- To build binary:

```console
make build
```

## Report a Bug

To file bugs, suggest improvements, or request new features, please open an [issue](https://github.com/scaleway/crossplane-provider-scaleway/issues).
