package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	cpb "github.com/OGLinuk/micro-todos/create/proto"
	"google.golang.org/grpc"
)

var (
	tpl          *template.Template
	createClient cpb.TodoCreateServiceClient
)

func init() {
	os.Setenv("SERVER_HOST", "0.0.0.0")
	os.Setenv("SERVER_PORT", "7001")

	SHOST := os.Getenv("SERVER_HOST")
	SPORT, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse server port env variable ...")
	}

	log.Printf("Client connected to grpc server on %s:%d ...", SHOST, SPORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", SHOST, SPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	createClient = cpb.NewTodoCreateServiceClient(conn)
}

func main() {
	PORT := 9001

	http.HandleFunc("/", index)
	log.Printf("Starting server on %d ...", PORT)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), nil)
}
