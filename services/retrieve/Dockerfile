FROM golang:1.14.0
ADD . /go/src/rserver
WORKDIR /go/src/rserver
RUN go get
RUN go build -o rserver .
EXPOSE 7002
CMD ["./rserver"]
