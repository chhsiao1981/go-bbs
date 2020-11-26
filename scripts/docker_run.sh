#!/bin/bash

branch=`git branch|grep '^*'|sed 's/^\* //g'|sed -E 's/^\(HEAD detached at //g'|sed -E 's/\)$//g'`
project=`basename \`pwd\``

docker run --name ${project} -p 3456:3456 -p 8888:8888 -p 48763:48763 ${project}:${branch}
