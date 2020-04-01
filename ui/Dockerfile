FROM golang:1.14.0
ADD . /go/src/ui
WORKDIR /go/src/ui
RUN go get
RUN go build -o ui .
EXPOSE 9001
CMD ["./ui"]
