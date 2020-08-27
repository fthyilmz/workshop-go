package handler

import (
	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Loging
// @Summary Login user
// @ID login
// @Accept  json
// @Produce  json
// @Tags Login
// @Param body body model.LoginCredentials true "body"
// @Success 200 {string} Token "{"token": XXX}"
// @Router /login [post]
func (h *LoginHandler) Login(c *gin.Context) {
	var credential model.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		c.JSON(200, gin.H{"error": "test"})
		return
	}
	var token string

	isUserAuthenticated := h.LoginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		token = h.JWtService.GenerateToken(credential.Username, true)
	}

	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}
