FROM reloni/goexample-base as builder

ADD ./src ./src/app
RUN go install -i $GOPATH/src/app

FROM alpine:3.10
ENV CGO_ENABLED=0 \
    GOPATH=/go \
    PATH=go/bin:$PATH
WORKDIR $GOPATH

RUN apk --update --no-cache add curl
COPY --from=builder $GOPATH/bin/app $GOPATH/bin/

EXPOSE 8080
USER nobody

CMD ["./bin/app"]
HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost:8080 || exit 1