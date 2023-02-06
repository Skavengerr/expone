package domain

type AccountInput struct {
	AccountID string `json:"id" dynamodbav:"id"`
	Name      string `json:"name" dynamodbav:"name"`
	Balance   int    `json:"balance" dynamodbav:"balance"`
}

type UpdateAccountInput struct {
	AccountID string `json:"id" dynamodbav:"id"`
	Balance   int    `json:"balance" dynamodbav:"balance"`
}
