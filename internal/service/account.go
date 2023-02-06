package service

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/repository"
)

type Account interface {
	Create(account domain.AccountInput) error
	UpdateBalance(accountID string, amount int64, transactionType domain.TransactionType) error
	Get(id string) string
	Delete(id string) error
}

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Create(account domain.AccountInput) error {
	return s.repo.Create(domain.AccountInput{
		AccountID: account.AccountID,
		Balance:   account.Balance,
	})
}

func (s *AccountService) UpdateBalance(accountID string, amount int64, transactionType domain.TransactionType) error {
	return s.repo.UpdateBalance(accountID, amount, transactionType)
}

func (s *AccountService) Get(id string) string {
	return s.repo.Get(id)
}

func (s *AccountService) Delete(accountID string) error {
	return s.repo.Delete(accountID)
}
