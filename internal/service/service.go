package service

import (
	"github.com/Skavengerr/expone/internal/repository"
)

type Services struct {
	Account     Account
	Transaction Transaction
}

func NewServices(repo *repository.Repositories) *Services {
	transactionService := NewTransactionService(repo.Transaction)
	accountService := NewAccountService(repo.Account)

	return &Services{
		Account:     accountService,
		Transaction: transactionService,
	}
}
