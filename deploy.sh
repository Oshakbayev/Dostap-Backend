#!/bin/bash

#cd /path/to/your/go/project
git pull origin main --rebase
go build -o bin ./cmd/main.go
pkill bin
nohup ./bin &
