package routers

import (
	controllers "online-store/httpserver/controller/auth"
	"online-store/utils"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController, authService utils.AuthHelper) *gin.RouterGroup {
	userRouter := route.Group("/users")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
	}
	return userRouter
}
