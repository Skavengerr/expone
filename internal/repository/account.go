package repository

import (
	"fmt"
	"strconv"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Account interface {
	Create(account domain.AccountInput) error
	Get(id string) string
	UpdateBalance(accountID string, amount int64, transactionType domain.TransactionType) error
	Delete(accountID string) error
}

type AccountRepo struct {
	db *dynamodb.DynamoDB
}

func NewAccountRepo(db *dynamodb.DynamoDB) *AccountRepo {
	return &AccountRepo{db}
}

// Insert account to dynamodb
func (e *AccountRepo) Create(account domain.AccountInput) error {
	_, err := e.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Item: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(account.AccountID),
			},
			"balance": {
				N: aws.String(strconv.Itoa(account.Balance)),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

// Get account from dynamodb
func (e *AccountRepo) Get(accountID string) string {
	account, err := e.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(accountID),
			},
		},
	})
	if err != nil {
		util.HandleAWSError(err)
		return ""
	}

	// check if Item is nil
	if account.Item == nil {
		return ""
	}

	id := account.Item["account_id"].S

	return *id
}

// Update account's balance to dynamodb
func (e *AccountRepo) UpdateBalance(accountID string, amount int64, transactionType domain.TransactionType) error {
	_, err := e.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(accountID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":balance": {
				N: aws.String(strconv.FormatInt(amount, 10)),
			},
		},
		UpdateExpression: aws.String("SET balance = if_not_exists(balance, :balance) - :balance"),
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

// Delete account from dynamodb
func (e *AccountRepo) Delete(accountID string) error {
	_, err := e.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(accountID),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

func getOperationDependsOnTransactionType(transactionType domain.TransactionType) string {
	if transactionType == domain.TransactionTypeExpense {
		return "-"
	}
	fmt.Println("transactionType: ", transactionType)

	return "+"
}
