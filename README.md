# Hello-Operator

This example is created based on two tutorials:

- ['Hello, World' tutorial with Kubernetes Operators](https://developers.redhat.com/blog/2020/08/21/hello-world-tutorial-with-kubernetes-operators#set_up_your_environment)
	- This tutorial is outdated, we need to use other references to update this tutorial
- [Quickstart for Go-based Operators](https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/)
- [Another Interesting Interactive Tutorial](https://www.katacoda.com/openshift/courses/operatorframework/go-operator-podset)
- [Very Interesting Tutorial About a Pod Controller](https://kubernetes.io/blog/2021/06/21/writing-a-controller-for-pod-labels/)


## Deploy the Operator with Direct Deploy

### Check Versions
```
$ go version
go version go1.16.9 linux/amd64

$ operator-sdk version
operator-sdk version: "v1.14.0", commit: "78f08b4852faf344ad3ef457c54f86087aaa0a0a", kubernetes version: "1.21", go version: "go1.16.9", GOOS: "linux", GOARCH: "amd64"
```

### Create Directory

```
mkdir $GOPATH/src/operators
mkdir $GOPATH/src/operators/hello-operator
```

### Generate the example application code

```
operator-sdk init hello-operator
```

### Add a custom resource definition (and a controller)

```
operator-sdk create api --group example.com --version v1alpha1 --kind=Traveller
```

The above command asks whether to create CRD and then asks whether to create controller. 

### Build Image

```bash
make docker-build
```

### Push Image to Local

```bash
make docker-push-local
```

### Direct Deploy

```bash
$ make deploy-local

$ kubectl get crds
NAME                               CREATED AT
podgroups.scheduling.sigs.k8s.io   2021-11-01T14:58:55Z
travellers.example.com.my.domain   2021-11-01T15:14:15Z

```

### Create CR and Verify
```bash
$ hello-operator kubectl apply -f config/samples/example.com_v1alpha1_traveller.yaml
traveller.example.com.my.domain/traveller-sample created

$  kubectl get travellers
NAME               AGE
traveller-sample   6m9s
 
 
$ kubectl get travellers -o yaml
apiVersion: v1
items:
- apiVersion: example.com.my.domain/v1alpha1
  kind: Traveller
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"example.com.my.domain/v1alpha1","kind":"Traveller","metadata":{"annotations":{},"name":"traveller-sample","namespace":"default"},"spec":null}
    creationTimestamp: "2021-11-01T15:18:35Z"
    generation: 1
    name: traveller-sample
    namespace: default
    resourceVersion: "3272"
    uid: b0235f0e-bdd8-4380-a239-1c7a63bc67e0
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

## Reconcile Function

### Minimal Working Example

- Add a `Println` in Reconcile function
- `kubctl create -f sample.yaml`
-  check if the `Println` output shows up in the log `kubectl logs -n hello-operator-system hello-operator-controller-manager-6f8cdff894-mns9w manager` 


## Add Anonther Controller for Pod

### Add a Controller without Resource

```bash
operator-sdk create api --group example.com --version v1alpha1 --kind=Pod --controller=true --resource=false
```