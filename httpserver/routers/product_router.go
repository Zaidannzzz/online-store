package routers

import (
	"online-store/httpserver/controller/product"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productHandler *product.ProductHandler) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/category/:category_id", productHandler.GetProductsByCategory)
	}
}
