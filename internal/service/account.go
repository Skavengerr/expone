package service

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/repository"
)

type Account interface {
	Insert(account domain.AccountInput) error
	Get(id string) string
	Delete(id int64) error
}

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) Insert(account domain.AccountInput) error {
	return s.repo.Create(domain.AccountInput{
		AccountID: account.AccountID,
		Balance:   account.Balance,
	})
}

func (s *AccountService) Get(id string) string {
	return s.repo.Get(id)
}

func (s *AccountService) Delete(id int64) error {
	return s.repo.Delete(id)
}
