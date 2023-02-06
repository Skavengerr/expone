package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Skavengerr/expone/config"
	"github.com/Skavengerr/expone/internal/repository"
	"github.com/Skavengerr/expone/internal/service"
	delivery "github.com/Skavengerr/expone/internal/transport/http"
	"github.com/Skavengerr/expone/pkg/database"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	cfg, err := config.InitViper(".")
	if err != nil {
		log.Fatal(err)
	}
	dynamo := database.InitDb(cfg)

	startServer(dynamo)
}

func startServer(dynamo *dynamodb.DynamoDB) {
	repos := repository.NewRepositories(dynamo)
	services := service.NewServices(repos)
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
