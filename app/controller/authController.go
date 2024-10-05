package controller

import (
	"bank/app/model/dto/request"
	"bank/app/model/dto/response"
	"bank/app/service/impl"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *impl.AuthService
}

// Login handles login
func (c *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := c.AuthService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.LoginResponse{UserID: userID.ID})
}

// Logout handles logout
func (c *AuthController) Logout(ctx *gin.Context) {
	userID := ctx.Param("id") // Get user ID from path variable

	if err := c.AuthService.Logout(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return // if error occured send or return the error
	}

	ctx.JSON(http.StatusOK, response.LogoutResponse{Message: "Logout successful"})
}
