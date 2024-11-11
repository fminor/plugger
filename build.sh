#!/bin/bash

pushd add-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd