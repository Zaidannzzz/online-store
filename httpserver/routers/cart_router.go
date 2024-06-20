package routers

import (
	"online-store/httpserver/controller/cart"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine, cartHandler *cart.CartHandler) {
	cartRoutes := router.Group("/cart")
	{
		cartRoutes.POST("/add", cartHandler.AddToCart)
		cartRoutes.GET("/", cartHandler.ViewCart)
		cartRoutes.DELETE("/remove/:product_id", cartHandler.RemoveFromCart)
		cartRoutes.DELETE("/clear", cartHandler.ClearCart)
	}
}
