#!/bin/bash

if [ "$#" -lt "1" ]
then
    echo "usage: run.sh"
fi

tags=default
ini_filename=00-config.template.ini

echo "to build: tags: ${tags}"
cd go-bbs
go build -tags ${tags}
cd ..

echo "to run go-bbs"
./go-bbs/go-bbs -ini ${ini_filename}
