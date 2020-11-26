#!/bin/bash

docker run --name go-bbs -p 3456:3456 -p 8888:8888 -p 48763:48763 go-bbs:dockerfile
