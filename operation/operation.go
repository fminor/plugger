// operation/operation.go
package operation

// Plugins should export a variable called "Plugin" which implements this interface
type Operation interface {
	Name() (n string)
	Compute(x float64, y float64) (r float64)
	Symbol() (t string)
}
