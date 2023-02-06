package http

import (
	"github.com/Skavengerr/expone/internal/domain"
	"github.com/Skavengerr/expone/internal/service"
	"github.com/gorilla/mux"
)

type ExpenseHandler interface {
	GetHistory(expense domain.ExpenseInput) error
	Insert(expense domain.ExpenseInput) error
	Update(expense domain.ExpenseInput) error
	Delete(expense domain.ExpenseInput) error
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

	expenseRouter := router.PathPrefix("/expense").Subrouter()
	accountRouter := router.PathPrefix("/account").Subrouter()

	initExpenseRouter(expenseRouter, h)
	initAccountRouter(accountRouter, h)

	return router
}
