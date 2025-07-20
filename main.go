package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"go-jwt/repositories"
	"go-jwt/services"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	UserRepository := repositories.NewUserRepository()
	UserService := services.NewUserService(UserRepository)
	UserController := controllers.NewUserController(UserService)

	r.POST("/signup", UserController.Signup)
	r.POST("/login", UserController.Login)
	r.GET("/validate", middleware.RequireAuth, UserController.Validate)
	r.GET("/logout", middleware.RequireAuth, UserController.Logout)

	r.Run()
}
