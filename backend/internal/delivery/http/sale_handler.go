package http

import (
	"go-clean/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	saleUsecase usecase.SaleUsecase
}

func NewSaleHandler(saleUsecase usecase.SaleUsecase) *SaleHandler {
	return &SaleHandler{saleUsecase: saleUsecase}
}

func (h *SaleHandler) CreateSale(c *gin.Context) {
	var req usecase.CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, err := h.saleUsecase.CreateSale(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sale)
}

func (h *SaleHandler) GetSale(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
		return
	}

	sale, err := h.saleUsecase.GetSale(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}

	c.JSON(http.StatusOK, sale)
}

func (h *SaleHandler) GetAllSales(c *gin.Context) {
	sales, err := h.saleUsecase.GetAllSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}

func (h *SaleHandler) GetSalesByCustomer(c *gin.Context) {
	customerID, err := strconv.ParseUint(c.Param("customer_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	sales, err := h.saleUsecase.GetSalesByCustomer(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}