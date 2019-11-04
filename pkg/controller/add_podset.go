package controller

import (
	"github.com/l0rd/operator-hello-world/pkg/controller/podset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, podset.Add)
}
