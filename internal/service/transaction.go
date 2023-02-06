package service

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/repository"
)

type Transaction interface {
	Operation(transaction domain.TransactionInput) error
	Delete(transaction domain.TransactionInput) error
}

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Operation(transaction domain.TransactionInput) error {
	return s.repo.Operation(domain.TransactionInput{
		AccountID:       transaction.AccountID,
		Date:            transaction.Date,
		Amount:          transaction.Amount,
		Category:        transaction.Category,
		Description:     transaction.Description,
		PaymentMethod:   transaction.PaymentMethod,
		Currency:        transaction.Currency,
		TransactionType: transaction.TransactionType,
	})
}

func (s *TransactionService) Delete(transaction domain.TransactionInput) error {
	return s.repo.Delete(transaction)
}
