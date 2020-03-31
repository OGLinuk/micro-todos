FROM golang:1.10.0
ADD . /go/src/userver
WORKDIR /go/src/userver
RUN go get
RUN go build -o userver .
EXPOSE 9003
CMD ["./userver"]
