package category

import (
	"net/http"
	"strconv"

	"online-store/httpserver/dto"
	"online-store/httpserver/services"
	"online-store/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategories(ctx *gin.Context)
	GetCategoryByID(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
	authService     utils.AuthHelper
}

func NewCategoryController(categoryService services.CategoryService, authService utils.AuthHelper) CategoryController {
	return &categoryController{categoryService, authService}
}

func (c *categoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	category, err := c.categoryService.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var dto dto.CategoryDTO
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	category, err := c.categoryService.CreateCategory(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, category)
}
