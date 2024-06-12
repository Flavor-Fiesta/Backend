package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

var jwtKey = []byte("11111") // Define tu clave secreta

// Claims es una estructura que contiene la información de los claims del JWT
type Claims struct {
    UserID int    `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

// Middleware de autenticación
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Verificar si la ruta es la de creación de roles
        if c.FullPath() == "/roles/crear" && c.Request.Method == "POST" {
            // Si es la ruta de creación de roles, permitir el acceso sin token
            c.Next()
            return
        }
		 // Verificar si la ruta es la de obtener todos los roles
		 if c.FullPath() == "/roles/" && c.Request.Method == "GET" {
            // Si es la ruta de obtener todos los roles, permitir el acceso sin token
            c.Next()
            return
        }

        // En otras rutas, verificar el token normalmente
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            web.Failure(c, http.StatusUnauthorized, errors.New("no authorization header provided"))
            c.Abort()
            return
        }

        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            web.Failure(c, http.StatusUnauthorized, errors.New("invalid token"))
            c.Abort()
            return
        }

        // Almacenar el rol y el userID en el contexto
        c.Set("user_id", claims.UserID)
        c.Set("rol", claims.Role)
        c.Next()
    }
}