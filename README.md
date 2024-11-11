# plugger

A practical walkthrough of creating an application that uses plugins.

## Quick start

```sh
./clean.sh
./build.sh
./run.sh
```

## References

* [Tutorial: Plugins with Go](https://medium.com/profusion-engineering/plugins-with-go-7ea1e7a280d3)
* [Ref: Go Workspaces](https://go.dev/doc/tutorial/workspaces)
* [Ref: plugin](https://pkg.go.dev/plugin)

## Design

Expected API: Operation

* calculate - Performs the calculation

API implementations:

* Addition
* Subtraction
* Multiplication
* Division

## Organization

* cmd/
* operation/
* add-plugin/
