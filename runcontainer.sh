#!/bin/bash
set -e

NAME=gotest

docker container rm -f $(docker container ps -a --filter "name=$NAME" -q) >> /dev/null 2>&1 || true && \
    docker run --rm -d --name $NAME -p 8080:8080 reloni/goexample