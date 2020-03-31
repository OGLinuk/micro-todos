FROM golang:1.10.0
ADD . /go/src/cserver
WORKDIR /go/src/cserver
RUN go get
RUN go build -o cserver .
EXPOSE 9001
CMD ["./cserver"]
