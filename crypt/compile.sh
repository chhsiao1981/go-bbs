#!/bin/bash

clang -g -c bbscrypt.c -Os -W -Wall -Wunused -Wno-missing-field-initializers -pipe -I./include -Qunused-arguments -Wno-parentheses-equality  -fcolor-diagnostics -Wno-invalid-source-encoding
ar cq libbbscrypt.a bbscrypt.o
ranlib libbbscrypt.a
