package routers

import (
	controllers "online-store/httpserver/controller/auth"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController) *gin.RouterGroup {
	userRouter := route.Group("/users")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
	}
	return userRouter
}
