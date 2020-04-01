FROM golang:1.14.0
ADD . /go/src/userver
WORKDIR /go/src/userver
RUN go get
RUN go build -o userver .
EXPOSE 7003
CMD ["./userver"]
