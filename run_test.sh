#!/bin/bash

COVERAGE_DIR=./coverage

GOCACHE=off

rm -rf ${COVERAGE_DIR}
mkdir -p ${COVERAGE_DIR}

go test -coverprofile=${COVERAGE_DIR}/cover.out -covermode=count -cover  ./test/...
go tool cover -html=${COVERAGE_DIR}/cover.out -o ${COVERAGE_DIR}/index.html
