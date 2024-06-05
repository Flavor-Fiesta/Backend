package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/carritoOrdenes"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type carritoOrdenesHandler struct {
	s carritoOrdenes.Service
}

func NewCarritoOrdenesHandler(s carritoOrdenes.Service) *carritoOrdenesHandler {
	return &carritoOrdenesHandler{
		s: s,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA UN NUEVO CARRITO ORDEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoOrdenesHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var carritoOrden domain.CarritosOrden
		err := c.ShouldBindJSON(&carritoOrden)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el carritoOrden utilizando el servicio
		createdCarritoOrden, err := h.s.CrearCarritoOrdenes(carritoOrden)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create carritoOrden"))
			return
		}
		// Devolver el carritoOrden creado con su ID asignado a la base de datos
		c.JSON(200, createdCarritoOrden)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE CARRITO ORDEN POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoOrdenesHandler) GetCarritoOrdenesByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		carritoOrden, err := h.s.GetCarritoOrdenesByID(id)
		if err != nil {
			fmt.Print("aca si")
			web.Failure(c, 404, errors.New("No se encuentra"))
			fmt.Print("aca no")
			return
		}
		web.Success(c, 200, carritoOrden)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UN CARRITO ORDEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoOrdenesHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var carritoOrden domain.CarritosOrden
		err = c.ShouldBindJSON(&carritoOrden)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedCarritoOrden, err := h.s.UpdateCarritoOrdenes(id, carritoOrden)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo el carritoOrden actualizado
		c.JSON(200, updatedCarritoOrden) // Asegúrate de que updatedCarritoOrden tenga el ID correcto
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR UN CARRITO ORDEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoOrdenesHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "123456" {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			err = h.s.DeleteCarritoOrdenes(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			c.JSON(200, gin.H{"message": "El carritoOrden se elimino correctamente"})
		} else {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}
