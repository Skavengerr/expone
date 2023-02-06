package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Skavengerr/expone/internal/domain"
)

// Get expense history from dynamodb
func (h *Handler) GetHistory(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}
}

// Insert expense to dynamodb
func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Insert(expense)
}

// Update expense to dynamodb
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Update(expense)
}

// Delete expense to dynamodb
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var expense domain.ExpenseInput
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.services.Expense.Delete(expense)
}
