package repository

import (
	"strconv"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Transaction interface {
	Operation(transaction domain.TransactionInput) error
	History(accountID string) ([]domain.TransactionInput, error)
	Delete(transaction domain.TransactionInput) error
}

type TransactionsRepo struct {
	db *dynamodb.DynamoDB
}

func NewTransactionsRepo(db *dynamodb.DynamoDB) *TransactionsRepo {
	return &TransactionsRepo{db}
}

// Insert transaction to dynamodb
func (e *TransactionsRepo) Operation(transaction domain.TransactionInput) error {
	_, err := e.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_TRANSACTIONS),
		Item: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(transaction.AccountID),
			},
			"date": {
				S: aws.String(transaction.Date),
			},
			"amount": {
				N: getAmountDependsOnTransactionType(transaction.Amount, transaction.TransactionType),
			},
			"transaction_type": {
				S: aws.String(string(transaction.TransactionType)),
			},
			"category": {
				S: aws.String(transaction.Category),
			},
			"description": {
				S: aws.String(transaction.Description),
			},
			"payment_method": {
				S: aws.String(transaction.PaymentMethod),
			},
			"currency": {
				S: aws.String(transaction.Currency),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

// Get transactions history for account from dynamodb
func (e *TransactionsRepo) History(accountID string) ([]domain.TransactionInput, error) {
	transactions, err := e.db.Query(&dynamodb.QueryInput{
		TableName: aws.String(TABLE_TRANSACTIONS),
		KeyConditions: map[string]*dynamodb.Condition{
			"account_id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(accountID),
					},
				},
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	var history []domain.TransactionInput

	for _, transaction := range transactions.Items {
		amount, _ := strconv.ParseInt(*transaction["amount"].N, 10, 64)

		history = append(history, domain.TransactionInput{
			AccountID:       *transaction["account_id"].S,
			Date:            *transaction["date"].S,
			Amount:          amount,
			Category:        *transaction["category"].S,
			Description:     *transaction["description"].S,
			PaymentMethod:   *transaction["payment_method"].S,
			Currency:        *transaction["currency"].S,
			TransactionType: domain.TransactionType(*transaction["transaction_type"].S),
		})
	}

	return history, err
}

// Delete transaction from dynamodb
func (e *TransactionsRepo) Delete(transaction domain.TransactionInput) error {
	_, err := e.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_TRANSACTIONS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(transaction.AccountID),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

func getAmountDependsOnTransactionType(amount int64, transactionType domain.TransactionType) *string {
	if transactionType == domain.TransactionTypeExpense {
		amount = -amount
	}

	return aws.String(strconv.FormatInt(amount, 10))
}
