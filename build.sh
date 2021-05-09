#!/bin/bash
go build -o recursive4darwin main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o recursive4ubuntu main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o recursive4windows main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o recursive4raspi main.go
