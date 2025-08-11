package main

import (
	"go-jwt/configuration/initializers"
	loggerhandler "go-jwt/configuration/loggerHandler"
	"go-jwt/controllers"
	"go-jwt/controllers/routes"
	"go-jwt/domain/user/services"
	"go-jwt/models/user/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Swagger API
// @version 1.0
// @description A API for User Authentication
// @termsOfService http://swagger.io/terms/
func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDatabase()
}

func main() {
	loggerhandler.Info("About to start application")

	repository := repositories.NewUserRepository()
	service := services.NewUserDomainService(repository)
	userController := controllers.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
