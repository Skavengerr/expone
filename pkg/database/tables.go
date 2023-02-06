package database

import (
	"fmt"

	"github.com/Skavengerr/expone/internal/repository"
	"github.com/Skavengerr/expone/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitTables(dynamo *dynamodb.DynamoDB) {
	CreateTableAccounts(dynamo)
	CreateTableTransactions(dynamo)
}

func CreateTableAccounts(dynamo *dynamodb.DynamoDB) {
	tableName := aws.String(repository.TABLE_ACCOUNTS)
	if isTableAlreadyExists(dynamo, tableName) {
		return
	}

	result, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("account_id"), AttributeType: aws.String("S")},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("account_id"), KeyType: aws.String("HASH")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: tableName,
	})

	fmt.Println("Created the table", result)

	if err != nil {
		util.HandleAWSError(err)
		return
	}
}

func CreateTableTransactions(dynamo *dynamodb.DynamoDB) {
	tableName := aws.String(repository.TABLE_TRANSACTIONS)
	if isTableAlreadyExists(dynamo, tableName) {
		return
	}

	result, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("account_id"), AttributeType: aws.String("S")},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("account_id"), KeyType: aws.String("HASH")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: tableName,
	})

	fmt.Println("Created the table", result)

	if err != nil {
		util.HandleAWSError(err)
		return
	}
}

// add func to check if tables already exist
func isTableAlreadyExists(dynamo *dynamodb.DynamoDB, tableName *string) bool {

	desc, err := dynamo.DescribeTable(&dynamodb.DescribeTableInput{TableName: tableName})
	if err == nil && *desc.Table.TableStatus == "ACTIVE" {
		return true
	}

	return false
}
