package main

import (
	"context"
	"fmt"
	"log"

	todopb "github.com/OGLinuk/micro-todos/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Update(ctx context.Context, req *todopb.UpdateRequest) (*todopb.UpdateResponse, error) {
	log.Println("Updating ...")

	t := req.GetTask()

	updateFilter := bson.D{{"_id", t.Id}}

	data := bson.D{{
		"$set", bson.D{
			{"name", t.GetName()},
			{"description", t.GetDescription()},
			{"priority", t.GetPriority()},
		},
	}}

	_, err := tododb.UpdateOne(mongoCtx, updateFilter, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &todopb.UpdateResponse{Task: t}, nil
}
