package handler

import (
	"errors"
	"fmt"
//	"net/http"
//	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/roles"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type rolesHandler struct {
	s roles.Service
}

// NewImagenHandler crea un nuevo controller de imagenes
func NewRolHandler(s roles.Service) *rolesHandler {
	return &rolesHandler{
		s: s,
	}
}

var listaRoles []domain.Rol
var ultimoRolID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA NUEVO ROL <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *rolesHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rol domain.Rol
		rol.ID = ultimoRolID
		ultimoRolID++
		err := c.ShouldBindJSON(&rol)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el rol utilizando el servicio
		createdRol, err := h.s.CrearRol(rol)
		if err != nil {
			web.Failure(c, 500, errors.New("Fallo la creacion de usuario, revise que los datos ingresados sean hmmmm"))
			return
		}
		// Devolver el producto creado con su ID asignado a la base de datos
		c.JSON(200, createdRol)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODOS LOS ROLES <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *rolesHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := h.s.BuscarTodosLosRoles()
		if err != nil {
			web.Failure(c, 500, fmt.Errorf("error buscando todos los usuarios: %w", err))
			return
		}
		web.Success(c, 200, roles)
	}
}

//
func (h *rolesHandler) CambiarRol() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Verificar que el usuario tenga rol de ADMIN
        // Asumiendo que tienes una forma de obtener el rol del usuario actual
        usuarioRol := c.GetString("rol") // ejemplo de cómo podrías obtener el rol del usuario
        if usuarioRol != "ADMIN" {
            web.Failure(c, 403, errors.New("forbidden: you do not have permission to change roles"))
            return
        }

        var req struct {
            UsuarioID int    `json:"usuario_id"`
            NuevoRol  string `json:"nuevo_rol"`
        }
        err := c.ShouldBindJSON(&req)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
            return
        }

        // Cambiar el rol del usuario utilizando el servicio
        err = h.s.CambiarRol(req.UsuarioID, req.NuevoRol)
        if err != nil {
            web.Failure(c, 500, errors.New("failed to change user role: " + err.Error()))
            return
        }

        web.Success(c, 200, "user role changed successfully")
    }
}