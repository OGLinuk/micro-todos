package main

import (
	"context"
	"fmt"
	"log"

	dpb "github.com/OGLinuk/micro-todos/services/delete/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Delete(ctx context.Context, req *dpb.DeleteRequest) (*dpb.DeleteResponse, error) {
	tID := req.GetId()

	log.Printf("Deleting %s ...", tID)
	id, err := primitive.ObjectIDFromHex(tID)
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

	return &dpb.DeleteResponse{Success: result.DeletedCount}, nil
}
