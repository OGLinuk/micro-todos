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
	os.Setenv("CSERVER_PORT", "7001") // Need to figure out how to access the services .env port value

	CSPORT, err := strconv.ParseInt(os.Getenv("CSERVER_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse server port env variable ...")
	}

	log.Printf("Connected to create-service grpc server at create-service:%d ...", CSPORT)
	conn, err := grpc.Dial(fmt.Sprintf("create-service:%d", CSPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	createClient = cpb.NewTodoCreateServiceClient(conn)

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	PORT := 9001

	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Printf("Starting server on %d ...", PORT)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), nil)
}
