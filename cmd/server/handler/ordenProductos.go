package handler

import (
	"errors"
	"fmt"

	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/ordenProducto"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type ordenProductoHandler struct {
	s ordenProductos.Service
}

// NewOrdenProductoHandler crea un nuevo handler para OrdenProducto
func NewOrdenProductoHandler(s ordenProductos.Service) *ordenProductoHandler {
	return &ordenProductoHandler{
		s: s,
	}
}

// Post crea una nueva relación orden-producto
func (h *ordenProductoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var op domain.OrdenProducto
		err := c.ShouldBindJSON(&op)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		op, err = h.s.CrearOrdenProducto(op)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create order product: " + err.Error()))
			return
		}

		// Devolver la relación creada
		web.Success(c, 200, op)
	}
}

// BuscarOrdenProducto busca una relación orden-producto por su ID
func (h *ordenProductoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		op, err := h.s.BuscaOrdenProducto(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se encuentra"))
			return
		}
		web.Success(c, 200, op)
	}
}

// Put actualiza una relación orden-producto existente
func (h *ordenProductoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var op domain.OrdenProducto
		err = c.ShouldBindJSON(&op)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		op, err = h.s.UpdateOrdenProducto(id, op)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo la relación actualizada
		c.JSON(200, op)
	}
}
