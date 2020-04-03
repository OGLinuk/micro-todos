package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	cpb "github.com/OGLinuk/micro-todos/services/create/proto"
	rpb "github.com/OGLinuk/micro-todos/services/retrieve/proto"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := []*rpb.Task{}

	req := &rpb.RetrieveAllRequest{Token: "test"}
	if resp, err := retrClient.RetrieveAll(context.Background(), req); err == nil {
		for _, cTask := range resp.Tasks {
			t = append(t, cTask)
		}
	} else {
		http.Error(w, fmt.Sprintf("handlers.go::index::createClient.Create::ERROR: %s", err.Error()), 500)
	}

	err := tpl.ExecuteTemplate(w, "index.html", t)
	if err != nil {
		http.Error(w, fmt.Sprintf("handlers.go::index::tpl.ExecuteTemplate::ERROR: %s", err.Error()), 500)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		desc := r.FormValue("desc")
		prio, err := strconv.ParseInt(r.FormValue("priority"), 10, 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("handlers.go::create::strconv.ParseInt::ERROR: %s", err.Error()), 500)
		}

		t := &cpb.Task{
			Name:        name,
			Description: desc,
			Priority:    int32(prio),
		}

		_, err = createClient.Create(context.Background(), &cpb.CreateRequest{Task: t})
		if err != nil {
			http.Error(w, fmt.Sprintf("handlers.go::create::createClient.Create::ERROR: %s", err.Error()), 500)
		}

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

	err := tpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("handlers.go::create::tpl.ExecuteTemplate::ERROR: %s", err.Error()), 500)
	}
}
