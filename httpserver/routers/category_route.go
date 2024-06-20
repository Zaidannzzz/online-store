package routers

import (
	"online-store/httpserver/controller/category"
	"online-store/utils"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(route *gin.RouterGroup, categoryController category.CategoryController, authService utils.AuthHelper) {
	route.Use(authService.ValidateToken())
	{
		route.GET("/categories", categoryController.GetAllCategories)
		route.GET("/categories/:id", categoryController.GetCategoryByID)
		route.POST("/categories", categoryController.CreateCategory)
	}
}
