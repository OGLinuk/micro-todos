FROM golang:1.10.0
ADD . /go/src/dserver
WORKDIR /go/src/dserver
RUN go get
RUN go build -o dserver .
EXPOSE 9004
CMD ["./dserver"]
