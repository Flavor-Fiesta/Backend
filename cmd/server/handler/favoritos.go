package handler

import (
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jfcheca/FlavorFiesta/internal/domain"
    "github.com/jfcheca/FlavorFiesta/internal/favoritos"
    "github.com/jfcheca/FlavorFiesta/pkg/web"
)

type favoritosHandler struct {
    s favoritos.Service
}

// NewFavoritosHandler crea un nuevo controlador de favoritos
func NewFavoritosHandler(s favoritos.Service) *favoritosHandler {
    return &favoritosHandler{
        s: s,
    }
}

var favorito domain.Favoritos
var ultimofavoritoID int = 1

// Post maneja la solicitud para agregar un nuevo favorito
func (h *favoritosHandler) Post() gin.HandlerFunc {
    return func(c *gin.Context) {
        var fav domain.Favoritos
        fav.ID = ultimofavoritoID
        ultimofavoritoID++
        err := c.ShouldBindJSON(&fav)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
            fmt.Println("Error al hacer bind del JSON:", err)
            return
        }

        // Crear el favorito utilizando el servicio
        createdFavorito, err := h.s.AgregarFavorito(fav)
        if err != nil {
            web.Failure(c, 500, errors.New("fallo la creacion del favorito, revise que los datos ingresados sean correctos"))
            fmt.Println("Error al agregar el favorito", fav, ":", err)
            return
        }

        // Devolver el favorito creado con su ID asignado
        c.JSON(200, createdFavorito)
    }
}