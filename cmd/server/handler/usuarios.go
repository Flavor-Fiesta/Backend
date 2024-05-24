package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/usuarios"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type usuariosHandler struct {
	s usuarios.Service
}

// NewImagenHandler crea un nuevo controller de imagenes
func NewUsuarioHandler(s usuarios.Service) *usuariosHandler {
	return &usuariosHandler{
		s: s,
	}
}

var listaUsuarios []domain.Usuarios
var ultimoUsuarioID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA NUEVA USUARIO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *usuariosHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario domain.Usuarios
		usuario.ID = ultimoUsuarioID
		ultimoUsuarioID++
		err := c.ShouldBindJSON(&usuario)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el producto utilizando el servicio
		createdUsuario, err := h.s.CrearUsuario(usuario)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create producto"))
			return
		}
		// Devolver el producto creado con su ID asignado a la base de datos
		c.JSON(200, createdUsuario)

	}

}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE USUARIO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *usuariosHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		usuario, err := h.s.BuscarUsuario(id)
		if err != nil {
			fmt.Print("aca si")
			web.Failure(c, 404, errors.New("No se encuentra"))
			fmt.Print("aca no")
			return
		}
		web.Success(c, 200, usuario)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODOS LOS URUARIOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (h *usuariosHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios, err := h.s.BuscarTodosLosUsuarios()
		if err != nil {
			web.Failure(c, 500, fmt.Errorf("error buscando todos los usuarios: %w", err))
			return
		}
		web.Success(c, 200, usuarios)
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *usuariosHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var usuario domain.Usuarios
		err = c.ShouldBindJSON(&usuario)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedImagen, err := h.s.UpdateUsuario(id, usuario)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo la imagen actualizado
		c.JSON(200, updatedImagen) // Asegúrate de que updatedImagen tenga el ID correcto
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH ACTUALIZA UNA IMAGEN O ALGUNO DE SUS CAMPOS <<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *usuariosHandler) Patch() gin.HandlerFunc {

    type Request struct {
        Nombre  string `json:"nombre"`
        Email    string `json:"email"`
        Telefono string `json:"telefono"`
		Password string `json:"password"`
    }

    return func(c *gin.Context) {
       
		/*token := c.GetHeader("TOKEN")
        if token == "" || token != os.Getenv("TOKEN") {
            web.Failure(c, http.StatusUnauthorized, errors.New("invalid token"))
            return
        }*/

        var r Request
        if err := c.ShouldBindJSON(&r); err != nil {
            web.Failure(c, http.StatusBadRequest, errors.New("invalid JSON"))
            return
        }
        idParam := c.Param("id")
        id, err := strconv.Atoi(idParam)
        if err != nil {
            web.Failure(c, http.StatusBadRequest, errors.New("invalid ID"))
            return
        }

        // Verificar si la imagen existe antes de actualizar
        _, err = h.s.BuscarUsuario(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, errors.New("odontologo not found"))
            return
        }

        // Crear una estructura de actualización solo con los campos proporcionados
        update := domain.Usuarios{}
        if r.Nombre != "" {
            update.Nombre = r.Nombre
        }
        if r.Email != "" {
            update.Email = r.Email
        }
		if r.Telefono != "" {
            update.Telefono = r.Telefono
        }
		if r.Password != "" {
            update.Password = r.Password
        }

        // Actualizar la imagen
        p, err := h.s.UpdateUsuario(id, update)
        if err != nil {
            web.Failure(c, http.StatusConflict, err)
            return
        }

        web.Success(c, http.StatusOK, p)
    }
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR UNA IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *usuariosHandler) DeleteUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "123456" {
			// Permitir la eliminación de la imagen con el token correcto
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			err = h.s.DeleteUsuario(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			// Se elimina el usuario correctamente, enviar mensaje de éxito
			c.JSON(200, gin.H{"message": "El usuario se elimino correctamente"})
		} else {
			// Token no válido
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}
