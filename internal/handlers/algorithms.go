package handlers

import (
	"net/http"

	"bank/internal/algorithms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HeapSort(c *gin.Context) {
	var nums []int

	if err := c.ShouldBindJSON(&nums); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.service.SortNumbers(nums)

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Search(c *gin.Context) {
	text := c.Query("text")
	pattern := c.Query("pattern")

	result := h.service.Search(text, pattern)

	c.JSON(http.StatusOK, result)
}
func (h *Handler) Kruskal(c *gin.Context) {
	var req struct {
		Vertices int               `json:"vertices"`
		Edges    []algorithms.Edge `json:"edges"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.service.GetMST(req.Vertices, req.Edges)

	c.JSON(http.StatusOK, result)
}
