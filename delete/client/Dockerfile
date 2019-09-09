FROM golang:1.10.0
ADD . /go/src/dclient
WORKDIR /go/src/dclient
RUN go get
RUN go build -o dclient .
EXPOSE 9704
CMD ["./dclient"]
