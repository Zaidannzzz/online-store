package controllers

import (
	"fmt"
	"net/http"

	"online-store/httpserver/dto"
	"online-store/httpserver/models"
	"online-store/httpserver/services"
	"online-store/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	authService utils.AuthHelper
}

func NewUserController(
	userService services.UserService,
	authService utils.AuthHelper,
) *userController {
	return &userController{userService, authService}
}

func (c *userController) Register(ctx *gin.Context) {
	var dto dto.RegisterDto

	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(http.StatusBadRequest, err.Error())
		return
	}

	_, err = c.userService.Register(&dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &dto)
	fmt.Println(http.StatusCreated, &dto)
}

func (c *userController) Login(ctx *gin.Context) {
	var dto dto.LoginDto
	err := ctx.BindJSON(&dto)
	if err != nil {
		fmt.Println(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.userService.Login(&dto)

	if err != nil {
		fmt.Println(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		fmt.Println(http.StatusUnauthorized, err.Error())
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	accessToken, refreshToken, err := c.authService.GenerateToken(user)

	if err != nil {
		fmt.Println(http.StatusInternalServerError, err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(http.StatusOK, models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	ctx.JSON(http.StatusOK, models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
