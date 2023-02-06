package database

import (
	"os"

	"github.com/Skavengerr/expone/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamo *dynamodb.DynamoDB

func InitDb(cfg config.Config) *dynamodb.DynamoDB {
	dynamo = connectDynamoDB(cfg)
	InitTables(dynamo)

	return dynamo
}

// connect dynamodb and return dynamodb client
func connectDynamoDB(cfg config.Config) *dynamodb.DynamoDB {
	os.Setenv("AWS_PROFILE", cfg.AWSProfile)

	db := dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: credentials.NewSharedCredentials("", cfg.DbName),
	})))

	return db
}
