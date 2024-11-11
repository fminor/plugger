# plugger

A practical walkthrough of creating an application that uses plugins.

## Quick start

```sh
./clean.sh
./build.sh
./run.sh
```

### Expected results

```plain
root ➜ /workspaces/plugger (main) $ ./run.sh 
Addition:  1.5 + 0.5 = 2
Subtraction:  1.5 - 0.5 = 1
Division:  1.5 / 0.5 = 3
Multiplication:  1.5 * 0.5 = 0.75
Addition:  2 + 4 = 6
Subtraction:  2 - 4 = -2
Division:  2 / 4 = 0.5
Multiplication:  2 * 4 = 8
Addition:  10 + 6 = 16
Subtraction:  10 - 6 = 4
Division:  10 / 6 = 1.6666666666666667
Multiplication:  10 * 6 = 60
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
