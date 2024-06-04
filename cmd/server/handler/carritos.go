package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/carritos"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type carritoHandler struct {
	s carritos.Service
}

func NewCarritoHandler(s carritos.Service) *carritoHandler {
	return &carritoHandler{
		s: s,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA UN NUEVO CARRITO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var carrito domain.Carrito
		err := c.ShouldBindJSON(&carrito)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el carrito utilizando el servicio
		createdCarrito, err := h.s.CrearCarrito(carrito)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create carrito"))
			return
		}
		// Devolver el carrito creado con su ID asignado a la base de datos
		c.JSON(200, createdCarrito)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE CARRITO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoHandler) GetCarritoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		carrito, err := h.s.GetCarritoByID(id)
		if err != nil {
			fmt.Print("aca si")
			web.Failure(c, 404, errors.New("No se encuentra"))
			fmt.Print("aca no")
			return
		}
		web.Success(c, 200, carrito)
	}
}


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UN CARRITO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var carrito domain.Carrito
		err = c.ShouldBindJSON(&carrito)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedCarrito, err := h.s.UpdateCarrito(id, carrito)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo el carrito actualizado
		c.JSON(200, updatedCarrito) // Asegúrate de que updatedCarrito tenga el ID correcto
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR UN CARRITO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *carritoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "123456" {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			err = h.s.DeleteCarrito(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			c.JSON(200, gin.H{"message": "El carrito se elimino correctamente"})
		} else {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}
