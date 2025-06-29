package product

import (
	"strconv"

	productDomain "go-clean/internal/domain/product"
	"go-clean/internal/usecase/product"
	"go-clean/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	productUseCase product.UseCase
}

func NewHandler(productUseCase product.UseCase) *Handler {
	return &Handler{
		productUseCase: productUseCase,
	}
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var req product.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	p, err := h.productUseCase.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case product.ErrProductExists:
			response.Conflict(c, "Product already exists")
		case product.ErrInvalidInput:
			response.BadRequest(c, "Invalid input data", err.Error())
		default:
			response.InternalServerError(c, "Failed to create product")
		}
		return
	}

	response.Created(c, "Product created successfully", p)
}

func (h *Handler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid product ID", "Product ID must be a valid UUID")
		return
	}

	p, err := h.productUseCase.GetProductByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			response.NotFound(c, "Product not found")
		default:
			response.InternalServerError(c, "Failed to retrieve product")
		}
		return
	}

	response.OK(c, "Product retrieved successfully", p)
}

func (h *Handler) GetProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	search := c.Query("search")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.BadRequest(c, "Invalid limit parameter", "Limit must be a valid integer")
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		response.BadRequest(c, "Invalid offset parameter", "Offset must be a valid integer")
		return
	}

	var products []*productDomain.Product
	if search != "" {
		products, err = h.productUseCase.SearchProducts(c.Request.Context(), search, limit, offset)
	} else {
		products, err = h.productUseCase.GetProducts(c.Request.Context(), limit, offset)
	}

	if err != nil {
		response.InternalServerError(c, "Failed to retrieve products")
		return
	}

	response.OK(c, "Products retrieved successfully", products)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid product ID", "Product ID must be a valid UUID")
		return
	}

	var req product.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	p, err := h.productUseCase.UpdateProduct(c.Request.Context(), id, &req)
	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			response.NotFound(c, "Product not found")
		case product.ErrProductExists:
			response.Conflict(c, "Product with this name already exists")
		case product.ErrInvalidInput:
			response.BadRequest(c, "Invalid input data", err.Error())
		default:
			response.InternalServerError(c, "Failed to update product")
		}
		return
	}

	response.OK(c, "Product updated successfully", p)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid product ID", "Product ID must be a valid UUID")
		return
	}

	err = h.productUseCase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			response.NotFound(c, "Product not found")
		default:
			response.InternalServerError(c, "Failed to delete product")
		}
		return
	}

	response.OK(c, "Product deleted successfully", nil)
}

func (h *Handler) UpdateStock(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid product ID", "Product ID must be a valid UUID")
		return
	}

	var req struct {
		Stock int `json:"stock" validate:"gte=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	err = h.productUseCase.UpdateStock(c.Request.Context(), id, req.Stock)
	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			response.NotFound(c, "Product not found")
		default:
			response.InternalServerError(c, "Failed to update stock")
		}
		return
	}

	response.OK(c, "Stock updated successfully", nil)
}

func (h *Handler) GetLowStockProducts(c *gin.Context) {
	thresholdStr := c.DefaultQuery("threshold", "10")
	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		response.BadRequest(c, "Invalid threshold parameter", "Threshold must be a valid integer")
		return
	}

	products, err := h.productUseCase.GetLowStockProducts(c.Request.Context(), threshold)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve low stock products")
		return
	}

	response.OK(c, "Low stock products retrieved successfully", products)
}