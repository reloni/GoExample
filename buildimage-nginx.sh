#!/bin/bash
set -e

docker build -f dockerfile-nginx -t reloni/goexample-nginx:latest . && \
    docker push reloni/goexample-nginx:latest