package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout allows a user to log out and clear the cookie Authorization
// @Summary User Logout
// @Description Allows a user to log out and clear the cookie Authorization.
// @Tags Authentication
// @Success 200 {string} "Message: Logged out successfully"
// @Router /logout [post]
func (uc *userController) Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
