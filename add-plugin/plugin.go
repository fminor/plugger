// add-plugin/plugin.go
package main

import (
	"github.com/fminor/plugger/operation"
)

type Addition struct{}

// Compile time check for Addition implements operation.Operation.
var _ operation.Operation = Addition{}

// Addition implements operation.Operation
func (p Addition) Compute(x float64, y float64) (r float64) {
	r = x + y
	return
}

var Plugin = Addition{}

func main() { /*empty because it does nothing*/ }
