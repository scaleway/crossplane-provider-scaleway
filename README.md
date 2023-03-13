# Crossplane Provider Scaleway

`crossplane-provider-scaleway` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Scaleway](https://www.scaleway.com/).

To install and use this provider:
* Install Upbound Universal Crossplane (UXP) into your Kubernetes cluster.
* Install the `Provider` and apply a `ProviderConfig`.
* Create a *managed resource* in Scaleway with Kubernetes.

## Prerequisites
This quickstart requires:
* a Kubernetes cluster with permissions to create pods and secrets
* a host with `kubectl` installed and configured to access the Kubernetes
  cluster
* your Scaleway [credentials](https://console.scaleway.com/project/credentials)

## Getting started

Run each command individually or copy to a local to prevent issues
running the commands in the terminal.

_Note:_ all commands use the current `kubeconfig` context and configuration.

### Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

### Install Upbound Universal Crossplane
Install Upbound Universal Crossplane (UXP) with the Up command-line `up uxp
install` command.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Find more information in the [Upbound UXP
documentation](https://docs.upbound.io/uxp/).

### Install the provider

Install the provider into the Kubernetes cluster with a Kubernetes
configuration file.

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-scaleway
spec:
  package: xpkg.upbound.io/scaleway/provider-scaleway:v0.1
EOF
```

Verify the provider installed with `kubectl get providers`.

The `INSTALLED` value should be `True`. It may take up to 5 minutes for
`HEALTHY` to report true.
```shell
$ kubectl get provider
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway   True        True   xpkg.upbound.io/scaleway/provider-scaleway:v0.1        11s
```

If there are issues downloading and installing the provider the `INSTALLED`
field is empty.

```shell
$ kubectl get providers
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway                      xpkg.upbound.io/scaleway/provider-scaleway:v0.1      76s
```

Use `kubectl describe providers` for more information.

### Create a Kubernetes secret resource for Scaleway
The provider requires credentials to create and manage Scaleway resources.

In a new folder, create a `secret.yaml` file and adapt it to your credentials.

Create `crossplane-system` namespace if not exists
```shell
$ kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -
```

Create a `secret.yaml` file

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
      "region": "fr-par",
      "zone": "fr-par-1"
    }
```

#### Configuration reference

| Provider Argument | Description                                                                                                                                                      |
|-------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `access_key`      | [Scaleway access key](https://console.scaleway.com/iam/api-keys)                                                                                          |
| `secret_key`      | [Scaleway secret key](https://console.scaleway.com/iam/api-keys)                                                                                          |
| `project_id`      | The [project ID](https://console.scaleway.com/project/settings) that will be used as default value for project-scoped resources.                                 |
| `region`          | The [region](https://developers.scaleway.com/en/quickstart/#region-and-zone)  that will be used as default value for all resources. (`fr-par` if none specified) |
| `zone`            | The [zone](https://developers.scaleway.com/en/quickstart/#region-and-zone) that will be used as default value for all resources. (`fr-par-1` if none specified)  |

### Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach your Scaleway
credentials to the installed provider.

```yaml
apiVersion: scaleway.upbound.io/v1beta1
kind: ProviderConfig
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

The `spec.secretRef` describes the parameters of the secret to use.
* `namespace` is the Kubernetes namespace the secret is in.
* `name` is the name of the Kubernetes `secret` object.
* `key` is the `Data` field from `kubectl describe secret`.

Apply this configuration with the secret `kubectl apply -f your-folder/`.

You can verify the `ProviderConfig` with `kubectl describe providerconfigs`.

### Create a managed resource
Create a managed resource to verify the provider is functioning.

This example creates a Scaleway Object Storage Bucket.

```yaml
apiVersion: object.scaleway.upbound.io/v1alpha1
kind: Bucket
metadata:
  name: object-bucket
spec:
  forProvider:
    name: crossplane-object-bucket
  providerConfigRef:
    name: default
```

Use `kubectl get buckets` to verify bucket creation.

```shell
$ kubectl get buckets
NAME                           READY   SYNCED   EXTERNAL-NAME                     AGE
object-bucket                  True    True     fr-par/crossplane-object-bucket   9s
```

Upbound created the bucket when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to
understand why.

### Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket`
object file. Verify removal of the bucket with `kubectl get buckets`

```shell
$ kubectl delete -f bucket.yml
bucket.object.scaleway.upbound.io "object-bucket" deleted

$ kubectl get buckets
No resources found
```

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/scaleway/crossplane-provider-scaleway/issues).
