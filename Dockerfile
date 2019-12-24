FROM reloni/goexample-base as builder

ADD ./src/service ./src/app
ADD ./src/healthcheck ./src/healthcheck
RUN go install -i $GOPATH/src/app && \
    go install -i $GOPATH/src/healthcheck

FROM alpine:3.10
ENV CGO_ENABLED=0 \
    GOPATH=/go \
    PATH=go/bin:$PATH
WORKDIR $GOPATH

COPY --from=builder $GOPATH/bin/app $GOPATH/bin/
COPY --from=builder $GOPATH/bin/healthcheck $GOPATH/bin/

EXPOSE 8080
USER nobody

CMD ["./bin/app"]
HEALTHCHECK --interval=15s --timeout=3s \
    CMD ./bin/healthcheck || exit 1