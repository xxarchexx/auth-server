#!/bin/sh
go install github.com/go-delve/delve/cmd/dlv@v1.9.0
cd ./cmd && dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2