package main

import (
	"context"
	"fmt"
	"log"

	todopb "github.com/OGLinuk/micro-todos/delete/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Delete(ctx context.Context, req *todopb.DeleteRequest) (*todopb.DeleteResponse, error) {
	log.Println("Deleting ...")

	tId := req.GetId()
	id, err := primitive.ObjectIDFromHex(tId)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	delFilter := bson.D{{"_id", id}}

	result, err := tododb.DeleteOne(mongoCtx, delFilter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &todopb.DeleteResponse{Success: result.DeletedCount}, nil
}
