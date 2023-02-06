package service

import (
	"github.com/Skavengerr/expone/internal/repository"
)

type Services struct {
	Expense Expense
}

func NewServices(repo *repository.Repositories) *Services {
	expenseService := NewExpenseService(repo.Expense)

	return &Services{
		Expense: expenseService,
	}
}
