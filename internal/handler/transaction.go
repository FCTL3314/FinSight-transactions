package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo repository.TransactionRepo
}

func NewHandler(repo repository.TransactionRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/transactions", h.createTransaction)
	r.Get("/transactions", h.listTransactions)
	return r
}

func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	var input models.CreateTransaction
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(r.Context(), &input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

func (h *Handler) listTransactions(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}
