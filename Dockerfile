FROM golang:1.13.4-alpine3.10

RUN apk add --no-cache git

WORKDIR $GOPATH

ADD ./src ./src/app

RUN go get "github.com/golang/glog"
RUN go get "github.com/gorilla/mux"

RUN go install -i $GOPATH/src/app

RUN rm -rf $GOPATH/src

EXPOSE 8080
USER nobody

CMD ["app"]