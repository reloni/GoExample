#!/bin/bash
set -e

NAME=gotest

docker container rm -f $(docker container ps -a --filter "name=$NAME" -q) >> /dev/null 2>&1 || true && \
    docker run -d --name $NAME -p 8080:8080 -p 40000:40000 --security-opt="seccomp:unconfined" reloni/goexampledebug

#docker run -d --name $NAME --rm -p 8080:8080 -p 40000:40000 --security-opt="apparmor=unconfined" --cap-add=SYS_PTRACE reloni/goexampledebug