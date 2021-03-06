package main

import (
	"context"
	"fmt"
	"log"

	rpb "github.com/OGLinuk/micro-todos/services/retrieve/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) RetrieveAll(ctx context.Context, req *rpb.RetrieveAllRequest) (*rpb.RetrieveAllResponse, error) {
	log.Println("Retrieving all ...")

	cur, err := tododb.Find(mongoCtx, bson.D{})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	defer cur.Close(mongoCtx)

	var tasks []*rpb.Task

	for cur.Next(mongoCtx) {
		doc := bson.D{}

		err := cur.Decode(&doc)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		mapped := doc.Map()
		tdTask := &rpb.Task{
			Id:          mapped["_id"].(primitive.ObjectID).Hex(),
			Name:        mapped["name"].(string),
			Description: mapped["description"].(string),
			Priority:    mapped["priority"].(int32),
		}

		tasks = append(tasks, tdTask)
	}

	return &rpb.RetrieveAllResponse{Tasks: tasks}, nil
}
