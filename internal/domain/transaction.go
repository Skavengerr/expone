package domain

type TransactionType string

const (
	TransactionTypeIncome  TransactionType = "income"
	TransactionTypeExpense TransactionType = "expense"
)

type TransactionInput struct {
	AccountID       string          `json:"account_id" dynamodbav:"account_id" validate:"required"`
	Amount          int64           `json:"amount" dynamodbav:"amount" validate:"required"`
	TransactionType TransactionType `json:"transaction_type" dynamodbav:"transaction_type" validate:"required,oneof=income expense"`
	// TODO: add current date as default
	Date          string `json:"date" dynamodbav:"date" default:"2020-01-01"`
	Currency      string `json:"currency" dynamodbav:"currency" default:"USD"`
	Description   string `json:"description" dynamodbav:"description"`
	PaymentMethod string `json:"payment_method" dynamodbav:"payment_method"`
	Category      string `json:"category" dynamodbav:"category"`
}
