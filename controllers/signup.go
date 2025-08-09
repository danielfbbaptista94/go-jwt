package controllers

import (
	loggerhandler "go-jwt/configuration/loggerHandler"
	"go-jwt/configuration/validation"
	requestdto "go-jwt/controllers/requestDTO"
	domain "go-jwt/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) Signup(c *gin.Context) {
	loggerhandler.Info("Init SignUp controller")
	var signup requestdto.SignupDTO

	if err := c.ShouldBindJSON(&signup); err != nil {
		loggerhandler.Error("Error trying to marshal object, error=%s\n", err)
		errorHandler := validation.ValidateUserError(err)
		c.JSON(errorHandler.Code, errorHandler)
		return
	}

	userDomain := domain.NewUserDomain(signup.Email, signup.Password)
	if err := uc.service.CreateUser(userDomain); err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusCreated, gin.H{})
}
