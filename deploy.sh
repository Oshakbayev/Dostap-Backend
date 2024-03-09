#!/bin/bash

#cd /path/to/your/go/project
cd ../Dostap-Backend
#pkill bin
git pull origin main --rebase
go build -o bin ./cmd/main.go
./bin&
disown
