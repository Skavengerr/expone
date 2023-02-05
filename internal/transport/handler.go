package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Skavengerr/expone/pkg/model"
	"github.com/gorilla/mux"
)

type Expense interface {
	Insert(expense model.Expense)
	Update(expense model.Expense)
	Delete(expense model.Expense)
}

type Handler struct {
	expenseService Expense
}

func NewHandler(expense Expense) *Handler {
	return &Handler{
		expenseService: expense,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/addExpense", h.Insert).Methods(http.MethodPost)
	router.HandleFunc("/expenses/{id}", h.Update).Methods(http.MethodPatch)
	router.HandleFunc("/expenses/{id}", h.Delete).Methods(http.MethodDelete)

	return router
}

// Insert expense to dynamodb
func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.expenseService.Insert(expense)
}

// Update expense to dynamodb
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.expenseService.Update(expense)
}

// Delete expense to dynamodb
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Errorf("Error decoding expense: %v", err)
		return
	}

	h.expenseService.Delete(expense)
}
