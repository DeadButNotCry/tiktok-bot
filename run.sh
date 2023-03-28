#!/bin/bash


export BOT_TOKEN=

go build -o ./build/main  ./cmd/main.go
./build/main
