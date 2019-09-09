FROM golang:1.10.0
ADD . /go/src/cclient
WORKDIR /go/src/cclient
RUN go get
RUN go build -o cclient .
EXPOSE 9701
CMD ["./cclient"]
