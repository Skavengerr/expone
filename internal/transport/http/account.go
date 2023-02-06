package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Skavengerr/expone/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// init account router
func initAccountRouter(accountRouter *mux.Router, h *Handler) {
	accountRouter.HandleFunc("/create", h.accountCreate).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{id:[0-9]+}", h.accountUpdateBalance).Methods(http.MethodPatch)
	accountRouter.HandleFunc("/{id:[0-9]+}", h.accountDelete).Methods(http.MethodDelete)
}

func (h *Handler) accountCreate(w http.ResponseWriter, r *http.Request) {
	var account domain.AccountInput
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Error: while decoding account")
		return
	}

	if account.Name == "" {
		log.Println("createAccount() error: name is empty")
		sendErrorResponse(w, http.StatusBadRequest, "Error: name is empty")

		return
	}

	account.Balance = 0
	account.AccountID = uuid.New().String()

	h.services.Account.Insert(account)
}

func (h *Handler) accountUpdateBalance(w http.ResponseWriter, r *http.Request) {
	var account domain.UpdateAccountInput

	if account.AccountID == "" || account.Balance == 0 {
		log.Println("updateAccountBalance() error: accountID or balance is empty")
		sendErrorResponse(w, http.StatusBadRequest, "Error: accountID or balance is empty")

		return
	}

	h.services.Account.Update(account)
}

func (h *Handler) accountDelete(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteBook() error:", err)
		sendErrorResponse(w, http.StatusBadRequest, "Error: while getting id from request")

		return
	}

	if id == 0 {
		log.Println("deleteBook() error: id is empty")
		sendErrorResponse(w, http.StatusBadRequest, "Error: Id should not be empty")

		return
	}

	err = h.services.Account.Delete(id)
	if err != nil {
		log.Println("deleteBook() error:", err)

		sendErrorResponse(w, http.StatusInternalServerError, "Error: Error while deleting book")
		return
	}

	w.WriteHeader(http.StatusOK)
}