#!/bin/bash
go build -o ./bin main.go
cp config.yaml ./bin/config.yaml

# zip bin
zip -r bin.zip bin