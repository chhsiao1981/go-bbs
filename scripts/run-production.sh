#!/bin/bash

if [ "$#" -lt "1" ]
then
    echo "usage: run-production.sh [ini-filename]"
    exit 255
fi

tags=production
ini_filename=$1

echo "to build: tags: ${tags} ini: %{ini_filename}"
cd go-bbs
go install -tags ${tags}
cd ..
echo "to run go-bbs"
go-bbs -ini ${ini_filename}
