# Using Code Generator on Top of Operator-SDK

- [reference](https://www.fatalerrors.org/a/writing-crd-by-mixing-kubeuilder-and-code-generator.html)

## Setup 

```bash
operator-sdk edit --multigroup
```

### Install Code Generator

#### Correct

check `go.mod` in `hello-operator`,  find `k8s.io/client-go` version

```bash
K8S_VERSION=v0.22.1
go get k8s.io/code-generator@$K8S_VERSION
go mod vendor
```
```bash
chmod +x ./vendor/k8s.io/code-generator/generate-groups.sh
```


#### Alternative

```bash

cd ~/opt

git clone git@github.com:kubernetes/code-generator.git

chmod +x ./code-generator/generate-groups.sh
```

