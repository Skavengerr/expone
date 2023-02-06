package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Skavengerr/expone/internal/repository"
	"github.com/Skavengerr/expone/internal/service"
	delivery "github.com/Skavengerr/expone/internal/transport/http"
	"github.com/Skavengerr/expone/pkg/database"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	dynamo := database.InitDb()

	startServer(dynamo)
}

// Initialize http server
func startServer(dynamo *dynamodb.DynamoDB) {
	// init deps
	repos := repository.NewRepositories(dynamo)
	services := service.NewServices(repos)
	fmt.Println("services", services)
	handlers := delivery.NewHandler(services)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
