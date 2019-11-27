FROM golang:1.13.4-alpine3.10 as builder

RUN apk add --no-cache git

ENV GOPATH /go 
WORKDIR $GOPATH

ADD ./src ./src/app
RUN go get "github.com/golang/glog"
RUN go get "github.com/gorilla/mux"
RUN go install -i $GOPATH/src/app

FROM alpine:3.10
ENV GOPATH /go 
WORKDIR $GOPATH

COPY --from=builder $GOPATH/bin $GOPATH/bin

EXPOSE 8080
USER nobody

CMD ["./bin/app"]