FROM golang:latest

RUN mkdir -p /go/src/callisto

ADD . /go/src/callisto

RUN apt-get update && apt-get install -y xdg-utils

WORKDIR /go/src/callisto

RUN go get
RUN go get github.com/tockins/realize


CMD realize start --run