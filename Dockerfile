FROM golang

ADD . /go/src/github.com/demas/observer

RUN go get github.com/demas/observer

RUN go install github.com/demas/observer

ENTRYPOINT ["/go/bin/observer"]

EXPOSE 4000
