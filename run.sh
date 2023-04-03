#!/bin/bash


export BOT_TOKEN=
export GROUP_ID=

go build -o ./build/main  ./cmd/main.go
./build/main
