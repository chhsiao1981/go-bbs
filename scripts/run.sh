#!/bin/bash

tags=${1:-default}

echo "to build: tags: ${tags}"
go build -tags ${tags} .
echo "to run main"
cd main
go run -tags ${tags} .
cd ..