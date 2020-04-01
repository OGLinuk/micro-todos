package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	dpb "github.com/OGLinuk/micro-todos/services/delete/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct{}
type TodoTask struct {
	ID          primitive.ObjectID `bson:_id,omitempty`
	Name        string             `bson:name`
	Description string             `bson:description`
	Priority    int32              `bson:priority`
}

var (
	db       *mongo.Client
	tododb   *mongo.Collection
	mongoCtx context.Context
)

func main() {
	log.Println("Starting delete-service gRPC server ...")
	SHOST := os.Getenv("DSERVER_HOST")
	SPORT, err := strconv.ParseInt(os.Getenv("DSERVER_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse server port env variable ...")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", SHOST, SPORT))
	if err != nil {
		log.Fatalf("Port %d already in use ...", SPORT)
	}

	srvr := grpc.NewServer()
	dpb.RegisterTodoDeleteServiceServer(srvr, &server{})

	log.Println("Initializing MongoDB connection ...")
	mongoCtx = context.Background()
	DB_URI := os.Getenv("DB_URI")
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Fatalf("MongoDB connection err: %v", err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	} else {
		log.Println("Successfully connected to MongoDB ...")
	}

	tododb = db.Database("microdb").Collection("todo")

	go func() {
		if err := srvr.Serve(listener); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	log.Printf("Server running on %s:%d ...", SHOST, SPORT)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Shutting server down ...")
	srvr.Stop()
	listener.Close()

	log.Println("Closing MongoDB connection ...")
	db.Disconnect(mongoCtx)

	log.Println("Shutdown complete ...")
}
