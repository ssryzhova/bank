package handlers

import (
	"bank/internal/algorithms"
	"bank/internal/customerror"
	"bank/internal/models"

	"github.com/gin-gonic/gin"
)

type IService interface {
	CreateAccount(req models.CreateAccountRequest) customerror.Error
	CloseAccount(email string) customerror.Error
	GetAccount(email string) (models.BankAccount, customerror.Error)

	AmountOperation(operation string, amount float64, account models.BankAccount) customerror.Error
	Transfer(req models.TransferRequest) customerror.Error
	SortNumbers(nums []int) []int
	Search(text, pattern string) []int
	GetMST(vertices int, edges []algorithms.Edge) []algorithms.Edge
}

type Handler struct {
	service IService
}

func New(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func Init(r *gin.Engine, h *Handler) {
	r.POST("/account/create", h.CreateAccount)
	r.POST("/account/close/:email", h.CloseAccount)

	r.GET("/balance/:email", h.GetBalance)
	r.POST("/amount/:email", h.AmountOperation)

	r.POST("/transfer", h.Transfer)

	r.POST("/algorithms/heapsort", h.HeapSort)
	r.GET("/algorithms/search", h.Search)
	r.POST("/algorithms/kruskal", h.Kruskal)
}
