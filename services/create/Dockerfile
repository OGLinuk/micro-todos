FROM golang:1.14.0
ADD . /go/src/cserver
WORKDIR /go/src/cserver
RUN go get
RUN go build -o cserver .
EXPOSE 7001
CMD ["./cserver"]
