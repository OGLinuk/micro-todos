FROM golang:1.10.0
ADD . /go/src/server
WORKDIR /go/src/server
RUN go get
RUN go build -o main .
EXPOSE 9001
CMD ["./main"]
