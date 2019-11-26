#!/bin/bash
set -e

NAME=gotest

docker container rm -f $(docker container ps --filter "name=$NAME" -q) || true && \
    docker run --rm -d --name $NAME -p 8080:8080 reloni/goexample