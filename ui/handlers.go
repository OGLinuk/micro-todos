package main

import (
	"context"
	"log"
	"net/http"

	cpb "github.com/OGLinuk/micro-todos/create/proto"
)

func index(w http.ResponseWriter, r *http.Request) {
	resp, err := createClient.Create(context.Background(), &cpb.CreateRequest{
		Token: "test",
		Task: &cpb.Task{
			Id:          "id00001",
			Name:        "test",
			Description: "this is a desc",
			Priority:    1,
		},
	})

	if err != nil {
		log.Printf("index::createClient.Create::ERROR: %s", err.Error())
	}

	log.Println(resp.Task)
}
