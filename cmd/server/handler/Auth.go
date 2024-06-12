package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/auth"
)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Authenticate the user and get the token
		token, err := h.service.Authenticate(auth.Credentials{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Get user details for the response
		user, err := h.service.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve user details"})
			return
		}

		// Return the user details and token in the response
		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
	}
}