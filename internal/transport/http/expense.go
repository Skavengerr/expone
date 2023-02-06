package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/gorilla/mux"
)

// init expense router
func initExpenseRouter(expenseRouter *mux.Router, h *Handler) {
	expenseRouter.HandleFunc("/", h.expenseGetHistory).Methods(http.MethodGet)
	expenseRouter.HandleFunc("/add", h.expenseInsert).Methods(http.MethodPost)
	expenseRouter.HandleFunc("/{id:[0-9]+}", h.expenseUpdate).Methods(http.MethodPatch)
	expenseRouter.HandleFunc("/{id:[0-9]+}", h.expenseDelete).Methods(http.MethodDelete)
}

// Get expense history from dynamodb
func (h *Handler) expenseGetHistory(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}
}

// Insert expense to dynamodb
func (h *Handler) expenseInsert(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Insert(expense)
}

// Update expense to dynamodb
func (h *Handler) expenseUpdate(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Update(expense)
}

// Delete expense to dynamodb
func (h *Handler) expenseDelete(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Delete(expense)
}
