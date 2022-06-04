#!/bin/bash
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -tracepath=. -o sensors_web