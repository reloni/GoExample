#!/bin/bash
set -e

docker build -f Dockerfile.base.debug -t reloni/goexample-base . && \
    docker build -f Dockerfile.debug -t reloni/goexampledebug .