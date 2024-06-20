package routers

import (
	"online-store/httpserver/controller/order"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderHandler *order.OrderHandler) {
	orderRoutes := router.Group("/order")
	{
		orderRoutes.POST("/checkout", orderHandler.Checkout)
	}
}
