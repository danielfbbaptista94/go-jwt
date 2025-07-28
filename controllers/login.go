package controllers

import (
	loggerhandler "go-jwt/configuration/loggerHandler"
	"go-jwt/configuration/validation"
	requestdto "go-jwt/controllers/requestDTO"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (uc *userController) Login(c *gin.Context) {
	loggerhandler.Info("Init Login controller")
	var loginDTO requestdto.LoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		loggerhandler.Error("Error trying to marshal object, error=%s\n", err)
		errorHandler := validation.ValidateUserError(err)
		c.JSON(errorHandler.Code, errorHandler)
		return
	}

	user, err := uc.service.FindUser(loginDTO.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(loginDTO.Password))
	if errPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is wrong",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.GetEmail(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
