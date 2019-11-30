#!/bin/bash
set -e

docker build -f Dockerfile.base -t reloni/goexample-base . && \
    docker build -f Dockerfile.debug -t reloni/goexampledebug .