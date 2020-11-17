#!/bin/bash

ini_filename=00-config.template.ini

echo "to build"
cd go-bbs
go build
cd ..

echo "to run go-bbs"
./go-bbs/go-bbs -ini ${ini_filename}
