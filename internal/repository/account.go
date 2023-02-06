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
	Get(id string) string
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

// Get account from dynamodb
func (e *AccountRepo) Get(id string) string {
	account, err := e.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_ACCOUNTS),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(id),
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

	accountID := account.Item["account_id"].S

	return *accountID
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
