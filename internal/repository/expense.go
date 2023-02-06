package repository

import (
	"strconv"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ExpensesRepo struct {
	db *dynamodb.DynamoDB
}

func NewExpensesRepo(db *dynamodb.DynamoDB) *ExpensesRepo {
	return &ExpensesRepo{db}
}

// Insert expense to dynamodb
func (e *ExpensesRepo) Insert(expense domain.ExpenseInput) error {
	_, err := e.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_EXPENSE),
		Item: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(expense.AccountID),
			},
			"date": {
				S: aws.String(expense.Date),
			},
			"amount": {
				N: aws.String(strconv.Itoa(expense.Amount)),
			},
			"category": {
				S: aws.String(expense.Category),
			},
			"description": {
				S: aws.String(expense.Description),
			},
			"payment_method": {
				S: aws.String(expense.PaymentMethod),
			},
			"currency": {
				S: aws.String(expense.Currency),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}

// Update expense to dynamodb
func (e *ExpensesRepo) Update(expense domain.ExpenseInput) error {
	_, err := e.db.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(TABLE_EXPENSE),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(expense.AccountID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":amount": {
				N: aws.String(strconv.Itoa(expense.Amount)),
			},
			":category": {
				S: aws.String(expense.Category),
			},
			":description": {
				S: aws.String(expense.Description),
			},
			":payment_method": {
				S: aws.String(expense.PaymentMethod),
			},
			":currency": {
				S: aws.String(expense.Currency),
			},
		},
		UpdateExpression: aws.String("set amount=:amount, category=:category, description=:description, payment_method=:payment_method, currency=:currency"),
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err

}

// Delete expense from dynamodb
func (e *ExpensesRepo) Delete(expense domain.ExpenseInput) error {
	_, err := e.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(TABLE_EXPENSE),
		Key: map[string]*dynamodb.AttributeValue{
			"account_id": {
				S: aws.String(expense.AccountID),
			},
		},
	})

	if err != nil {
		util.HandleAWSError(err)
	}

	return err
}
