package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_ACCOUNTS = "accounts"
	TABLE_EXPENSE  = "expense"
	TABLE_INCOME   = "income"
)

type Repositories struct {
	Account Account
	Expense Expense
}

func NewRepositories(dynamo *dynamodb.DynamoDB) *Repositories {
	return &Repositories{
		Account: NewAccountRepo(dynamo),
		Expense: NewExpensesRepo(dynamo),
	}
}
