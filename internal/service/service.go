package service

import (
	"github.com/Skavengerr/expone/internal/repository"
)

type Services struct {
	Account Account
	Expense Expense
}

func NewServices(repo *repository.Repositories) *Services {
	expenseService := NewExpenseService(repo.Expense)
	accountService := NewAccountService(repo.Account)

	return &Services{
		Account: accountService,
		Expense: expenseService,
	}
}
