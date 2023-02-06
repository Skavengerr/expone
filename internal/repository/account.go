package repository

import (
	"strconv"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Account interface {
	Create(account domain.AccountInput) error
	UpdateBalance(account domain.UpdateAccountInput) error
	Delete(id int64) error
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

// Update account's balance to dynamodb
func (e *AccountRepo) UpdateBalance(account domain.UpdateAccountInput) error {
	_, err := e.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(account.AccountID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":balance": {
				N: aws.String(strconv.Itoa(account.Balance)),
			},
		},
		UpdateExpression: aws.String("set balance=:balance"),
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err

}

// Delete account from dynamodb
func (e *AccountRepo) Delete(accountID int64) error {
	_, err := e.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(strconv.Itoa(int(accountID))),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}
