# Crossplane Provider Scaleway

`crossplane-provider-scaleway` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Scaleway](https://www.scaleway.com/).

Complete the following steps to: 
* Install Upbound Universal Crossplane (UXP) into your Kubernetes cluster.
* Install the `Provider` and apply a `ProviderConfig`.
* Create a *managed resource* in Scaleway with Kubernetes.

## Prerequisites
To perform the following steps, make sure you have:
* Your Scaleway [credentials](https://console.scaleway.com/project/credentials)
* A Kubernetes cluster with permissions to create pods and secrets
* A host with `kubectl` installed and configured to access the Kubernetes
  cluster

## Getting started

You can run each command individually or copy them to a local to avoid issues related to running commands in a terminal.

_Note:_ All commands use the current `kubeconfig` context and configuration.

### Install the Up command-line

Run the following command to download and install the Upbound `up` command-line.
_Note:_ To learn more about the Up command-line, refer to the [official dedicated documentation](https://docs.upbound.io/cli/).

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

### Install Upbound Universal Crossplane

Run the Up command-line `up uxp install` to install Upbound Universal Crossplane (UXP).
_Note:_ To learn more about Upbound Universal Crossplane (UXP), refer to the [official dedicated documentation](https://docs.upbound.io/uxp/).

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

### Install the provider

1. Install the provider into the Kubernetes cluster with a Kubernetes
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

2. Run `kubectl get providers` to verify the installed provider. The `INSTALLED` value should return as `True`. 

_Note:_  The procedure may take up to 5 minutes for `HEALTHY` to report true.

You should get an output similar to the following one, providing details about the provider.

```shell
$ kubectl get provider
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway   True        True   xpkg.upbound.io/scaleway/provider-scaleway:v0.1        11s
```

If there are any issue during the process of downloading and installing the provider, the `INSTALLED` field will return as empty. In that case, run `kubectl describe providers` to get more information.

```shell
$ kubectl get providers
NAME                INSTALLED   HEALTHY   PACKAGE                                             AGE
provider-scaleway                      xpkg.upbound.io/scaleway/provider-scaleway:v0.1      76s
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
      "region": "fr-par",
      "zone": "fr-par-1"
    }
```

2. If it does not already exist, run the following command to create a `crossplane-system` namespace:

```shell
$ kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -
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

1. Create a `ProviderConfig` Kubernetes configuration file to attach your Scaleway credentials to the previously installed provider.

Modify the values in the example according to your needs. Refer to the configuration reference information to understand the requested values.

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

2. Run `kubectl apply -f your-folder/` to apply this configuration with the secret.

3. Run `kubectl describe providerconfigs` to verify the `ProviderConfig`.

#### Configuration reference 

The `spec.secretRef` describes the parameters of the secret to use.
* `namespace` is the Kubernetes namespace the secret is in.
* `name` is the name of the Kubernetes `secret` object.
* `key` is the `Data` field from `kubectl describe secret`.

### Create a managed resource

1. Create a managed resource to see if the provider is properly functioning.

The following example creates a Scaleway Object Storage bucket.

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

2. Run `kubectl get buckets` to get details on the bucket's creation.

You should get an output similar to the following one, providing details about the bucket.

```shell
$ kubectl get buckets
NAME                           READY   SYNCED   EXTERNAL-NAME                     AGE
object-bucket                  True    True     fr-par/crossplane-object-bucket   9s
```

The bucket is successfully created when both the values for `READY` and `SYNCED` are `True`.

3. If there are any issue during the bucket creation process, the `READY` and/or `SYNCED` fields will return as empty. In that case, run `kubectl describe` to get more information.

### Delete the managed resource

1. Run `kubectl delete -f` (with the same `Bucket` file) to remove the managed resource.

2. Run `kubectl get buckets` to verify whether the bucket was properly removed.

You should get an output similar to this, providing details about the status of the bucket.

```shell
$ kubectl delete -f bucket.yml
bucket.object.scaleway.upbound.io "object-bucket" deleted

$ kubectl get buckets
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
