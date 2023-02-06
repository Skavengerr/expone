package http

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/service"
	"github.com/gorilla/mux"
)

type TransactionHandler interface {
	GetHistory(transaction domain.TransactionInput) error
	Insert(transaction domain.TransactionInput) error
	Update(transaction domain.TransactionInput) error
	Delete(transaction domain.TransactionInput) error
}

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	transactionRouter := router.PathPrefix("/transaction").Subrouter()
	accountRouter := router.PathPrefix("/account").Subrouter()

	initTransactionRouter(transactionRouter, h)
	initAccountRouter(accountRouter, h)

	return router
}
