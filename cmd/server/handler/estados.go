package handler

import (
	"errors"
	"fmt"

	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/estados"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type estadosHandler struct {
	s estados.Service
}

// NewEstadoHandler crea un nuevo controlador de estados
func NewEstadoHandler(s estados.Service) *estadosHandler {
	return &estadosHandler{
		s: s,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR NUEVO ESTADO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *estadosHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var estado domain.Estado
		err := c.ShouldBindJSON(&estado)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: "+err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el estado utilizando el servicio
		createdEstado, err := h.s.CrearEstado(estado)
		if err != nil {
			web.Failure(c, 500, errors.New("fallo la creacion del estado, revise que los datos ingresados sean correctos"))
			return
		}
		// Devolver el estado creado con su ID asignado por la base de datos
		c.JSON(200, createdEstado)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE ESTADO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (h *estadosHandler) BuscarEstado() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		producto, err := h.s.BuscarEstado(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se encuentra"))
			return
		}
		web.Success(c, 200, producto)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTENER TODOS LOS ESTADOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *estadosHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		estados, err := h.s.BuscarTodosLosEstados()
		if err != nil {
			web.Failure(c, 500, fmt.Errorf("error buscando todos los estados: %w", err))
			return
		}
		web.Success(c, 200, estados)
	}
}
