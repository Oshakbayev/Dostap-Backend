#!/bin/bash

#cd /path/to/your/go/project
cd ../Dostap-Backend
#pkill -f "go run ./cmd/main.go"
git pull origin main --rebase
go run ./cmd/main.go
#go bin ./cmd/main.go
#pkill bin
#nohup ./bin &
