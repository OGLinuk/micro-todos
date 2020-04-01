package main

import (
	"context"
	"fmt"
	"net/http"

	cpb "github.com/OGLinuk/micro-todos/create/proto"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := &cpb.Task{}

	task := &cpb.Task{
		Name:        "test",
		Description: "this is a desc",
		Priority:    int32(3),
	}

	req := &cpb.CreateRequest{Token: "test", Task: task}
	if resp, err := createClient.Create(context.Background(), req); err == nil {
		t = resp.Task
	} else {
		http.Error(w, fmt.Sprintf("handlers.go::index::createClient.Create::ERROR: %s", err.Error()), 500)
	}

	err := tpl.ExecuteTemplate(w, "index.html", t)

	if err != nil {
		http.Error(w, fmt.Sprintf("handlers.go::index::tpl.ExecuteTemplate::ERROR: %s", err.Error()), 500)
	}
}
