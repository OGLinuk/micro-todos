FROM golang:1.10.0
ADD . /go/src/rclient
WORKDIR /go/src/rclient
RUN go get
RUN go build -o rclient .
EXPOSE 9702
CMD ["./rclient"]
