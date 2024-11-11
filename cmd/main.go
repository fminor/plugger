// cmd/main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"plugin"
	"strconv"

	"github.com/fminor/plugger/operation"
)

var pluginPathList []string

func LoadConfig() {
	f, err := os.ReadFile("config.json")
	if err != nil {
		// NOTE: in real cases, deal with this error
		panic(err)
	}
	json.Unmarshal(f, &pluginPathList)
}

var pluginList []*operation.Operation

func LoadPlugins() {
	// Allocate a list for storing all our plugins
	pluginList = make([]*operation.Operation, 0, len(pluginPathList))
	for _, p := range pluginPathList {
		// We use plugin.Open to load the plugin by path
		plg, err := plugin.Open(p)
		if err != nil {
			// NOTE: in real cases, deal with this error
			panic(err)
		}

		// Search for variable named "Plugin"
		v, err := plg.Lookup("Plugin")
		if err != nil {
			// NOTE: in real cases, deal with this error
			panic(err)
		}

		// Cast symbol to protocol type
		castV, ok := v.(operation.Operation)
		if !ok {
			// NOTE: in real cases, deal with this error
			panic("Could not cast plugin")
		}

		pluginList = append(pluginList, &castV)
	}
}

// Let's throw this here so it loads the plugins as soon as we import this module
func init() {
	LoadConfig()
	LoadPlugins()
}

func PerformComputations(x float64, y float64) {
	for _, plg := range pluginList {
		r := (*plg).Compute(x, y)
		fmt.Println("Operation: ", plg, "X: ", x, " Y: ", y, "=", r)
	}
}

func main() {
	x := os.Args[1]
	y := os.Args[2]

	xFloat, _ := strconv.ParseFloat(x, 64)
	yFloat, _ := strconv.ParseFloat(y, 64)

	PerformComputations(xFloat, yFloat)
}
