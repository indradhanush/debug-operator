package controller

import (
	"github.com/indradhanush/k8s-codegen-test/debug-operator/pkg/controller/debug"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, debug.Add)
}
