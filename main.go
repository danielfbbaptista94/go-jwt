package main

import (
	"go-jwt/configuration/initializers"
	loggerhandler "go-jwt/configuration/loggerHandler"
	"go-jwt/controllers"
	"go-jwt/controllers/routes"
	"go-jwt/domain/user/services"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDatabase()
}

func main() {
	loggerhandler.Info("About to start application")

	service := services.NewUserDomainService()
	userController := controllers.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
