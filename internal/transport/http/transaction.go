package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/gorilla/mux"
)

// init transaction router
func initTransactionRouter(transactionRouter *mux.Router, h *Handler) {
	transactionRouter.HandleFunc("/", h.transactionGetHistory).Methods(http.MethodGet)
	transactionRouter.HandleFunc("/add", h.transactionOperation).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/{id:[0-9]+}", h.transactionDelete).Methods(http.MethodDelete)
}

// Get transaction history from dynamodb
func (h *Handler) transactionGetHistory(w http.ResponseWriter, r *http.Request) {
	var transaction domain.TransactionInput
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		fmt.Errorf("Error decoding transaction: %v", err)
		return
	}
}

// Insert transaction to dynamodb
func (h *Handler) transactionOperation(w http.ResponseWriter, r *http.Request) {
	var transaction domain.TransactionInput
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Println("transactionOperation() error:", err)
		sendErrorResponse(w, http.StatusBadRequest, "Error decoding transaction")
		return
	}

	if transaction.AccountID == "" || transaction.Amount == 0 || transaction.TransactionType == "" {
		log.Println("transactionOperation() error: required fields are empty")
		sendErrorResponse(w, http.StatusBadRequest, "Error: required fields are empty")
		return
	}

	accountID := h.services.Account.Get(transaction.AccountID)
	if accountID == "" {
		log.Println("transactionOperation() error:", err)
		sendErrorResponse(w, http.StatusBadRequest, "Error: accountID does not exist")
		return
	}

	h.services.Transaction.Operation(transaction)
}

// Delete transaction to dynamodb
func (h *Handler) transactionDelete(w http.ResponseWriter, r *http.Request) {
	var transaction domain.TransactionInput
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		fmt.Errorf("Error decoding transaction: %v", err)
		return
	}

	h.services.Transaction.Delete(transaction)
}
