FROM golang:1.14.0
ADD . /go/src/dserver
WORKDIR /go/src/dserver
RUN go get
RUN go build -o dserver .
EXPOSE 7004
CMD ["./dserver"]
