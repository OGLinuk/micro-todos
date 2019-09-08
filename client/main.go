package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	todopb "github.com/OGLinuk/micro-todos/proto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	Client todopb.TodoServiceClient
)

func main() {
	SHOST := os.Getenv("SERVER_HOST")
	SPORT, err := strconv.ParseInt(os.Getenv("SERVER_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse server port env variable ...")
	}

	log.Printf("Client connected to grpc server on %s:%d ...", SHOST, SPORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", SHOST, SPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	Client = todopb.NewTodoServiceClient(conn)

	var g *gin.Engine
	g = gin.Default()

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Fatalf("Failed to parse debug env variable ...")
	}

	if debug == true {
		log.Println("Starting in debug mode ...")
		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		cfg.AddAllowHeaders("access-control-allow-origin")
		g.Use(cors.New(cfg))
	} else {
		log.Println("Starting client in production mode ...")
		gin.SetMode(gin.ReleaseMode)
	}

	g.GET("/create/:name/:desc/:priority", CreateHandler)
	g.GET("/retrieve/:id/", RetrieveHandler)
	g.GET("/update/:id/:name/:desc/:priority", UpdateHandler)
	g.GET("/delete/:id", DeleteHandler)
	g.GET("/retrall", RetrieveAllHandler)

	CHOST := os.Getenv("CLIENT_HOST")
	CPORT, err := strconv.ParseInt(os.Getenv("CLIENT_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("Could not parse client port env variable ...")
	}

	log.Printf("Client running on %s:%d ...", CHOST, CPORT)
	if err := g.Run(fmt.Sprintf("%s:%d", CHOST, CPORT)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}