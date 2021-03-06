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

func (s *server) Retrieve(ctx context.Context, req *rpb.RetrieveRequest) (*rpb.RetrieveResponse, error) {
	log.Println("Retrieving ...")

	tId := req.GetId()
	id, err := primitive.ObjectIDFromHex(tId)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	retrFilter := bson.D{{"_id", id}}

	doc := bson.D{}
	err = tododb.FindOne(mongoCtx, retrFilter).Decode(&doc)
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

	return &rpb.RetrieveResponse{Task: tdTask}, nil
}
