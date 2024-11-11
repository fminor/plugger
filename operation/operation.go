// operation/operation.go
package operation

// Plugins should export a variable called "Plugin" which implements this interface
type Operation interface {
	Compute(x float64, y float64) (r float64)
}
