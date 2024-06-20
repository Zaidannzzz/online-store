package product

import (
	"net/http"
	"strconv"

	"online-store/httpserver/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("category_id"))

	products, err := h.productService.GetProductsByCategory(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
