package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/ordenes"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type ordenHandler struct {
	s ordenes.Service
}

func NewOrdenHandler(s ordenes.Service) *ordenHandler {
	return &ordenHandler{
		s: s,
	}
}

func (h *ordenHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ord domain.Orden
		err := c.ShouldBindJSON(&ord)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: "+err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear la orden utilizando el servicio
		createdOrd, err := h.s.CrearOrden(ord)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create order"))
			return
		}
		// Devolver la orden creada con su ID asignado a la base de datos
		c.JSON(200, createdOrd)
	}
}

func (h *ordenHandler) GetOrdenByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		ord, err := h.s.BuscarOrden(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se encuentra la orden"))
			return
		}
		web.Success(c, 200, ord)
	}
}

func (h *ordenHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var ord domain.Orden
		err = c.ShouldBindJSON(&ord)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedOrd, err := h.s.UpdateOrden(id, ord)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo la orden actualizada
		c.JSON(200, updatedOrd)
	}
}

func (h *ordenHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "123456" {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			err = h.s.DeleteOrden(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			c.JSON(200, gin.H{"message": "La orden se eliminó correctamente"})
		} else {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}