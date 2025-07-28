package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
