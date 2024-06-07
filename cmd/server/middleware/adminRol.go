package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

// Middleware para verificar el rol de ADMIN
func AdminRoleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        rol, exists := c.Get("rol")
        if !exists || rol != "ADMIN" {
            web.Failure(c, http.StatusForbidden, errors.New("forbidden: you do not have permission to change roles"))
            c.Abort()
            return
        }
        c.Next()
    }
}