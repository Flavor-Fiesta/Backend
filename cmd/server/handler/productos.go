package handler

import (
	"errors"
	"fmt"
	"net/http"

	//	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/internal/productos"
	"github.com/jfcheca/FlavorFiesta/pkg/web"
	//	"strings"
)

type productoHandler struct {
	s productos.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s productos.Service) *productoHandler {
	return &productoHandler{
		s: s,
	}
}

var listaProductos []domain.Producto
var ultimoProductoID int = 1

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREA UNA NUEVA BEBIDA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *productoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var producto domain.Producto
		producto.ID = ultimoProductoID
		ultimoProductoID++
		err := c.ShouldBindJSON(&producto)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
			fmt.Println("Error al hacer bind del JSON:", err)
			return
		}

		// Crear el producto utilizando el servicio
		createdProducto, err := h.s.CrearProducto(producto)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create producto"))
			return
		}
		// Devolver el producto creado con su ID asignado a la base de datos
		c.JSON(200, createdProducto)

	}

}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE PRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *productoHandler) BuscarProducto() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		producto, err := h.s.BuscarProducto(id)
		if err != nil {
			fmt.Print("aca si")
			web.Failure(c, 404, errors.New("No se encuentra"))
			fmt.Print("aca no")
			return
		}
		web.Success(c, 200, producto)
	}
}



//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *productoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var producto domain.Producto
		err = c.ShouldBindJSON(&producto)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		updatedProducto, err := h.s.UpdateProducto(id, producto)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// Devolver solo el producto actualizado
		c.JSON(200, updatedProducto) // Asegúrate de que updatedProducto tenga el ID correcto
	}
}

/// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH ACTUALIZA UN PRODUCTO O ALGUNO DE SUS CAMPOS <<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *productoHandler) Patch() gin.HandlerFunc {

	type Request struct {

	Nombre             string     `json:"nombre"`
    Codigo             string     `json:"codigo"`
    Categoria          string     `json:"categoria"`
    FechaDeAlta        string     `json:"fechaDeVencimiento"`
    FechaDeVencimiento string     `json:"fechaDeVencimiento"`

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

        // Verificar si el producto existe antes de actualizar
        _, err = h.s.BuscarProducto(id)
        if err != nil {
            web.Failure(c, http.StatusNotFound, errors.New("odontologo not found"))
            return
        }

        /// Crear una estructura de actualización solo con los campos proporcionados
		update := domain.Producto{}
		
		if r.Nombre != "" {
			update.Nombre = r.Nombre
		}
		if r.Codigo != "" {
			update.Codigo = r.Codigo
		}
		if r.Categoria != "" {
			update.Categoria = r.Categoria
		}
		if r.FechaDeAlta != "" {
			update.FechaDeAlta = r.FechaDeAlta
		}
		if r.FechaDeVencimiento != "" {
			update.FechaDeVencimiento = r.FechaDeVencimiento
		}

        // Actualizar el producto
        p, err := h.s.UpdateProducto(id, update)
        if err != nil {
            web.Failure(c, http.StatusConflict, err)
            return
        }

        web.Success(c, http.StatusOK, p)
    }
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR UN PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *productoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "123456" {
			// Permitir la eliminación del producto con el token correcto
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Failure(c, 400, errors.New("invalid id"))
				return
			}
			err = h.s.DeleteProducto(id)
			if err != nil {
				web.Failure(c, 404, err)
				return
			}
			// Se elimina el producto correctamente, enviar mensaje de éxito
			c.JSON(200, gin.H{"message": "El producto se elimino correctamente"})
		} else {
			// Token no válido
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
	}
}
