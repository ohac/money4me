#!/bin/bash
set -e
go build
# cat a.json | go run main.go
cat a.json.gpg | gpg -d | go run main.go
