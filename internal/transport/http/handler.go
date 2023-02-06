package http

import (
	"fmt"
	"net/http"

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
	fmt.Println("NewHandler: services", services)
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Expenses router
	expenseRouter := r.PathPrefix("/expense").Subrouter()
	{
		expenseRouter.HandleFunc("/", h.GetHistory).Methods(http.MethodGet)
		expenseRouter.HandleFunc("/add", h.Insert).Methods(http.MethodPost)
		expenseRouter.HandleFunc("/{id:[0-9]+}", h.Update).Methods(http.MethodPatch)
		expenseRouter.HandleFunc("/{id:[0-9]+}", h.Delete).Methods(http.MethodDelete)
	}

	return r
}
