package controllers

import (
	"github.com/gin-gonic/gin"
	loggerhandler "go-jwt/configuration/loggerHandler"
	"go-jwt/configuration/validation"
	requestdto "go-jwt/controllers/requestDTO"
	"net/http"
)

// Login allows a user to log in and obtain an authentication token
// @Summary User Login
// @Description Allows a user to log in and obtain an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginDTO body requestdto.LoginDTO true "User login credentials"
// @Success 200
// @Cookies {string} Authorization "Authentication token"
// @Failure 403 {object} errorhandler.ErrorHandler "error: Invalid login credentials"
// @Router /login [post]
func (uc *userController) Login(c *gin.Context) {
	loggerhandler.Info("Init Login controller")
	var loginDTO requestdto.LoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		loggerhandler.Error("Error trying to marshal object, error=%s\n", err)
		errorHandler := validation.ValidateUserError(err)
		c.JSON(errorHandler.Code, errorHandler)
		return
	}

	_, token, err := uc.service.Login(loginDTO)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err,
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}
