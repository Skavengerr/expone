package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Skavengerr/expone/internal/repository"
	"github.com/Skavengerr/expone/internal/transport"
	"github.com/Skavengerr/expone/pkg/database"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	currentDate = time.Now().Local().Format("2006-01-02")
)

func main() {
	dynamo := database.InitDb()

	//mockAddExpenses(dynamo)
	startServer(dynamo)

}

// Initialize http server
func startServer(dynamo *dynamodb.DynamoDB) {

	// init handler
	repo := repository.NewExpenses(dynamo)
	handler := transport.NewHandler(repo)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
