#!/bin/bash
set -e

docker build -f dockerfile.base -t reloni/goexample-base . && \
    docker build -f dockerfile.debug -t reloni/goexampledebug .