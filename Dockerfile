FROM reloni/goexample-base as builder

ADD ./src ./src/app
RUN go install -i $GOPATH/src/app

FROM alpine:3.10
ENV CGO_ENABLED=0 \
    GOPATH=/go \
    PATH=go/bin:$PATH
WORKDIR $GOPATH

COPY --from=builder $GOPATH/bin $GOPATH/bin

EXPOSE 8080
USER nobody

CMD ["./bin/app"]