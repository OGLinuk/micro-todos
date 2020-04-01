package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	cpb "github.com/OGLinuk/micro-todos/services/create/proto"
	dpb "github.com/OGLinuk/micro-todos/services/delete/proto"
	rpb "github.com/OGLinuk/micro-todos/services/retrieve/proto"
	upb "github.com/OGLinuk/micro-todos/services/update/proto"

	"google.golang.org/grpc"
)

var (
	tpl          *template.Template
	createClient cpb.TodoCreateServiceClient
	retrClient   rpb.TodoRetrieveServiceClient
	updateClient upb.TodoUpdateServiceClient
	deleteClient dpb.TodoDeleteServiceClient
)

func init() {
	// Create-service connection
	conn, err := grpc.Dial(fmt.Sprintf("create-service:%d", 7001), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	createClient = cpb.NewTodoCreateServiceClient(conn)

	// Retrieve-service connection
	conn, err = grpc.Dial(fmt.Sprintf("retrieve-service:%d", 7002), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	retrClient = rpb.NewTodoRetrieveServiceClient(conn)

	// Update-service connection
	conn, err = grpc.Dial(fmt.Sprintf("update-service:%d", 7003), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	updateClient = upb.NewTodoUpdateServiceClient(conn)

	// Delete-service connection
	conn, err = grpc.Dial(fmt.Sprintf("delete-service:%d", 7004), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	deleteClient = dpb.NewTodoDeleteServiceClient(conn)

	log.Println("Successfully connected to all services ...")

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	HOST := os.Getenv("UI_HOST")
	PORT, err := strconv.ParseInt(os.Getenv("UI_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse server port env variable ...")
	}
	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	defer log.Printf("Starting server at %s:%d ...", HOST, PORT)
	http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), nil)
}
