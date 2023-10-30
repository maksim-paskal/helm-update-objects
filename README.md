# Fix helm upgrading

## Problem

When you try to upgrade a old helm chart, you can get the following error:

```
current release manifest contains removed kubernetes api(s) for this kubernetes version and it is therefore unable to build the kubernetes objects for performing the diff. error from kubernetes: unable to recognize "": no matches for kind "Ingress" in version "networking.k8s.io/v1beta1"
```

## Posible solutions

1. Remove old release and install new one (if possible)
2. Update objects in the installed chart release with this command

```bash
go run github.com/maksim-paskal/helm-update-objects/cmd@latest \
--kubeconfig ~/.kube/config \
--namespace my-release-namespace \
--release-name my-release \
--dry-run=false
```

## Rules to update helm metadata

rules are defined in [pkg/config/config.go](./pkg/config/config.go)

