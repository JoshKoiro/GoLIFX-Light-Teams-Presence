#!/bin/bash
go build -o ./bin/Teams-LIFX.exe main.go
cp config.yaml ./bin/config.yaml
touch ./bin/.env

# zip bin
zip -r bin.zip bin