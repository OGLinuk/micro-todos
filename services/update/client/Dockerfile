FROM golang:1.10.0
ADD . /go/src/uclient
WORKDIR /go/src/uclient
RUN go get
RUN go build -o uclient .
EXPOSE 9703
CMD ["./uclient"]
