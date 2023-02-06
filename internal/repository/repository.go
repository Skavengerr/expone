package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_ACCOUNTS     = "accounts"
	TABLE_TRANSACTIONS = "transaction"
)

type Repositories struct {
	Account     Account
	Transaction Transaction
}

func NewRepositories(dynamo *dynamodb.DynamoDB) *Repositories {
	return &Repositories{
		Account:     NewAccountRepo(dynamo),
		Transaction: NewTransactionsRepo(dynamo),
	}
}
