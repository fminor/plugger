// add-plugin/plugin.go
package main

import (
	"github.com/fminor/plugger/operation"
)

type Multiplication struct{}

// Compile time check for Multiplication implements operation.Operation.
var _ operation.Operation = Multiplication{}

// Multiplication implements operation.Operation
func (p Multiplication) Name() (n string) {
	n = "Multiplication"
	return
}

func (p Multiplication) Symbol() (t string) {
	t = "*"
	return
}

func (p Multiplication) Compute(x float64, y float64) (r float64) {
	r = x * y
	return
}

var Plugin = Multiplication{}

func main() { /*empty because it does nothing*/ }
