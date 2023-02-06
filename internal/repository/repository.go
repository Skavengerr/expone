package repository

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_ACCOUNTS = "accounts"
	TABLE_EXPENSE  = "expense"
	TABLE_INCOME   = "income"
)

type Expense interface {
	Insert(expense domain.ExpenseInput) error
	Update(expense domain.ExpenseInput) error
	Delete(expense domain.ExpenseInput) error
}

type Repositories struct {
	Expense Expense
}

func NewRepositories(dynamo *dynamodb.DynamoDB) *Repositories {
	return &Repositories{
		Expense: NewExpensesRepo(dynamo),
	}
}
