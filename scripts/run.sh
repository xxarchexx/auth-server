#!/bin/sh
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./cmd
./server