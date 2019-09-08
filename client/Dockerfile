FROM golang:1.10.0
ADD . /go/src/client
WORKDIR /go/src/client
RUN go get
RUN go build -o main .
EXPOSE 9002
CMD ["./main"]
