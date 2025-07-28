package routes

import (
	"go-jwt/configuration/middleware"
	"go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	rg *gin.RouterGroup,
	userController controllers.UserControllerInterface,
) {

	rg.GET("/logout", middleware.RequireAuth, userController.Logout)
	rg.POST("/signup", userController.Signup)
	rg.POST("/login", userController.Login)
}
