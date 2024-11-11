// add-plugin/plugin.go
package main

import (
	"github.com/fminor/plugger/operation"
)

type Subtraction struct{}

// Compile time check for Subtraction implements operation.Operation.
var _ operation.Operation = Subtraction{}

// Subtraction implements operation.Operation
func (p Subtraction) Name() (n string) {
	n = "Subtraction"
	return
}

func (p Subtraction) Symbol() (t string) {
	t = "-"
	return
}

func (p Subtraction) Compute(x float64, y float64) (r float64) {
	r = x - y
	return
}

var Plugin = Subtraction{}

func main() { /*empty because it does nothing*/ }
