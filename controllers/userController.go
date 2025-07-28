package controllers

import (
	"go-jwt/domain/user/services"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service services.UserDomainService
}

func NewUserControllerInterface(
	seviceInterface services.UserDomainService,
) UserControllerInterface {
	return &userController{service: seviceInterface}
}

type UserControllerInterface interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
