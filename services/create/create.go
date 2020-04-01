package main

import (
	"context"
	"fmt"
	"log"

	cpb "github.com/OGLinuk/micro-todos/services/create/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Create(ctx context.Context, req *cpb.CreateRequest) (*cpb.CreateResponse, error) {
	t := req.GetTask()

	data := TodoTask{
		Name:        t.GetName(),
		Description: t.GetDescription(),
		Priority:    t.GetPriority(),
	}

	checked := bson.D{}
	checkFilter := bson.D{{"name", data.Name}}
	err := tododb.FindOne(mongoCtx, checkFilter).Decode(&checked)
	if err != nil {
		log.Println("Inserting new document ...")
		result, err := tododb.InsertOne(mongoCtx, data)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		oid := result.InsertedID.(primitive.ObjectID)
		t.Id = oid.Hex()
	} else {
		log.Println("Found existing:", checked)
		mapped := checked.Map()
		t.Id = mapped["_id"].(primitive.ObjectID).Hex()
		t.Name = mapped["name"].(string)
		t.Description = mapped["description"].(string)
		t.Priority = mapped["priority"].(int32)
	}

	return &cpb.CreateResponse{Task: t}, nil
}
