package database

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamo *dynamodb.DynamoDB

func InitDb() *dynamodb.DynamoDB {
	dynamo = connectDynamoDB()
	return dynamo
}

// connect dynamodb and return dynamodb client
func connectDynamoDB() *dynamodb.DynamoDB {
	os.Setenv("AWS_PROFILE", "your_profile_name")

	db := dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", "expone"),
	})))

	return db
}
