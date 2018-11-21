# Debug operator

A minimal repository to reproduce an edge case with k8s.io/code-generator + operator-sdk.

## Prerequisites

- Golang
- [operator-sdk](https://github.com/operator-framework/operator-sdk)

## Steps to reproduce

The skeleton for this package was generated using the `operator-sdk` by following the steps mentioned [here](https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md). Running the following commands will generate the clientset:

```
go get -u k8s.io/code-generator
cd $GOPATH/src/k8s.io/code-generator/
git checkout kubernetes-1.11.4
cd $GOPATH/src/github.com/indradhanush/debug-operator
CODEGEN_PKG=../../../k8s.io/code-generator hack/update-codegen.sh
```

## The issue

The generated clientset under at [pkg/client/clientset/versioned/typed/foo/v1alpha1/foo_client.go](pkg/client/clientset/versioned/typed/foo/v1alpha1/foo_client.go) does not contain any client code apart from the package headers.

## The workaround

This issue  has a fix and workaround mentioned here: https://github.com/operator-framework/operator-sdk/issues/678#issuecomment-435458003
