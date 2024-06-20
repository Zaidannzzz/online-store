package cart

import (
	"net/http"
	"strconv"

	"online-store/httpserver/dto"
	"online-store/httpserver/services"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	var CartItemDTO dto.CartItemDTO
	if err := c.ShouldBindJSON(&CartItemDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	err := h.cartService.AddToCart(userID.(uint), CartItemDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

func (h *CartHandler) ViewCart(c *gin.Context) {
	userID, _ := c.Get("userID")

	cart, err := h.cartService.GetCartByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	cartItemID, _ := strconv.Atoi(c.Param("product_id"))

	err := h.cartService.RemoveFromCart(uint(cartItemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, _ := c.Get("userID")

	err := h.cartService.ClearCart(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}
