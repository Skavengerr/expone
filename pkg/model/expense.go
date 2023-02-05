package model

type Expense struct {
	Date          string `json:"date" dynamodbav:"date"`
	Amount        int    `json:"amount" dynamodbav:"amount"`
	Category      string `json:"category" dynamodbav:"category"`
	Description   string `json:"description" dynamodbav:"description"`
	PaymentMethod string `json:"payment_method" dynamodbav:"payment_method"`
	AccountID     string `json:"account_id" dynamodbav:"account_id"`
	Currency      string `json:"currency" dynamodbav:"currency"`
}

type Income = Expense

type Accounts struct {
	AccountID string `json:"account_id" dynamodbav:"account_id"`
	Balance   int    `json:"balance" dynamodbav:"balance"`
}

const (
	TABLE_EXPENSE  = "expense"
	TABLE_INCOME   = "income"
	TABLE_ACCOUNTS = "accounts"
)
