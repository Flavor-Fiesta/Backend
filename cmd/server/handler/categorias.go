package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/categorias"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
)

type categoriasHandler struct {
	s categorias.Service
}

// NewCategoriaHandler crea un nuevo controller de categorias
func NewCategoriaHandler(s categorias.Service) *categoriasHandler {
	return &categoriasHandler{
		s: s,
	}
}

var listaCategorias []domain.Usuarios
var ultimaCategoriaID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA NUEVA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *categoriasHandler) Post() gin.HandlerFunc {
    return func(c *gin.Context) {
        var categoria domain.Categoria
        categoria.ID = ultimaCategoriaID
        ultimaCategoriaID++
        err := c.ShouldBindJSON(&categoria)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
            fmt.Println("Error al hacer bind del JSON:", err)
            return
        }

        // Inicializar el campo Productos como un array vacío para devolver la consulta en postman con array vacio y no NULL
        categoria.Productos = []domain.Producto{}

        // Crear la categoría utilizando el servicio
        createdCategoria, err := h.s.CrearCategoria(categoria)
        if err != nil {
            web.Failure(c, 500, errors.New("failed to create categoria"))
            return
        }

        // Devolver la categoría creada con su ID asignado a la base de datos
        c.JSON(200, createdCategoria)
    }
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE CATEGORIA POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *categoriasHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		categoria, err := h.s.BuscarCategoria(id)
		if err != nil {
			web.Failure(c, 404, errors.New("No se encuentra"))
			return
		}
		web.Success(c, 200, categoria)
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODAS LAS CATEGORIAS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *categoriasHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios, err := h.s.BuscarTodosLasCategorias()
		if err != nil {
			web.Failure(c, 500, fmt.Errorf("error buscando todos los usuarios: %w", err))
			return
		}
		web.Success(c, 200, usuarios)
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *categoriasHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var categoria domain.Categoria
		err = c.ShouldBindJSON(&categoria)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedCategoria, err := h.s.Update(id, categoria)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo la categoria actualizado
		c.JSON(200, updatedCategoria) 
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH ACTUALIZA UNA CATEGORIA O ALGUNOS DE SUS CAMPOS <<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *categoriasHandler) Patch() gin.HandlerFunc {

    type Request struct {
        Nombre  	   string `json:"nombre"`
        Descripcion    string `json:"descripcion"`
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

        // Verificar si la categoria existe antes de actualizar
        _, err = h.s.BuscarCategoria(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, errors.New("categoria not found"))
            return
        }

        // Crear una estructura de actualización solo con los campos proporcionados
        update := domain.Categoria{}
        if r.Nombre != "" {
            update.Nombre = r.Nombre
        }
        if r.Descripcion != "" {
            update.Descripcion = r.Descripcion
        }
        // Actualizar la categoria
        p, err := h.s.Update(id, update)
        if err != nil {
            web.Failure(c, http.StatusConflict, err)
            return
        }

        web.Success(c, http.StatusOK, p)
    }
}

/*func (h *categoriasHandler) Patch() gin.HandlerFunc {

    type Request struct {
        Nombre  	   string `json:"nombre"`
        Descripcion    string `json:"descripcion"`
    }

    return func(c *gin.Context) {
       
		/*token := c.GetHeader("TOKEN")
        if token == "" || token != os.Getenv("TOKEN") {
            web.Failure(c, http.StatusUnauthorized, errors.New("invalid token"))
            return
        }*/

   /*     var r Request
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

        // Verificar si la categoria existe antes de actualizar
        _, err = h.s.BuscarCategoria(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, errors.New("categoria not found"))
            return
        }

        // Crear una estructura de actualización solo con los campos proporcionados
        update := domain.Categoria{}
        if r.Nombre != "" {
            update.Nombre = r.Nombre
        }
        if r.Descripcion != "" {
            update.Descripcion = r.Descripcion
        }

        // Actualizar la categoria
        p, err := h.s.UpdateCategoria(id, update)
        if err != nil {
            web.Failure(c, http.StatusConflict, err)
            return
        }
        web.Success(c, http.StatusOK, p)
    }
}
*/
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR UNA CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *categoriasHandler) DeleteCategoria() gin.HandlerFunc {
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
			err = h.s.DeleteCategoria(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			// Se elimina la categoria correctamente, enviar mensaje de éxito
			c.JSON(200, gin.H{"message": "La categoria se elimino correctamente"})
		} else {
			// Token no válido
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}
