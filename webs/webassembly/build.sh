#!/bin/sh

set GOARCH=wasm&&set GOOS=js&&go build -o ./dist/hello.wasm hello.go