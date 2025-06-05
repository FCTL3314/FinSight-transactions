package handler

import (
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repository.TransactionRepo
}

func NewHandler(repo repository.TransactionRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/transactions", h.createTransaction)
	r.GET("/transactions", h.listTransactions)
}

func (h *Handler) createTransaction(c *gin.Context) {
	var input models.CreateTransaction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := h.repo.Create(c.Request.Context(), input.ToTransaction()); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (h *Handler) listTransactions(c *gin.Context) {
	list, err := h.repo.FindAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}
