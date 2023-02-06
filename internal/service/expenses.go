package service

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/repository"
)

type Expense interface {
	Insert(expense domain.ExpenseInput) error
	Update(expense domain.ExpenseInput) error
	Delete(expense domain.ExpenseInput) error
}

type ExpenseService struct {
	repo repository.Expense
}

func NewExpenseService(repo repository.Expense) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) Insert(expense domain.ExpenseInput) error {
	return s.repo.Insert(domain.ExpenseInput{
		AccountID:     expense.AccountID,
		Date:          expense.Date,
		Amount:        expense.Amount,
		Category:      expense.Category,
		Description:   expense.Description,
		PaymentMethod: expense.PaymentMethod,
		Currency:      expense.Currency,
	})
}

func (s *ExpenseService) Update(expense domain.UpdateExpenseInput) error {
	return s.repo.Update(expense)
}

func (s *ExpenseService) Delete(expense domain.ExpenseInput) error {
	return s.repo.Delete(expense)
}
