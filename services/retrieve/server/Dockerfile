FROM golang:1.10.0
ADD . /go/src/rserver
WORKDIR /go/src/rserver
RUN go get
RUN go build -o rserver .
EXPOSE 9002
CMD ["./rserver"]
