// add-plugin/plugin.go
package main

import (
	"github.com/fminor/plugger/operation"
)

type Division struct{}

// Compile time check for Division implements operation.Operation.
var _ operation.Operation = Division{}

// Division implements operation.Operation
func (p Division) Name() (n string) {
	n = "Division"
	return
}

func (p Division) Symbol() (t string) {
	t = "/"
	return
}

func (p Division) Compute(x float64, y float64) (r float64) {
	r = x / y
	return
}

var Plugin = Division{}

func main() { /*empty because it does nothing*/ }
