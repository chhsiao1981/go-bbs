#!/bin/bash

if [ "$#" -lt "1" ]
then
    echo "usage: run-custom.sh [ini-filename]"
    exit 255
fi

tags=custom
ini_filename=$1

echo "to build: tags: ${tags} ini: %{ini_filename}"
go install -tags ${tags}
echo "to run go-bbs"
go-bbs -ini ${ini_filename}
