package handler

import (
	"errors"
	"fmt"
	"log"

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
var ordenProd domain.OrdenProducto
var ultimaOrdenProdID int = 1
// Post crea una nueva relación orden-producto
func (h *ordenProductoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
        var orden domain.Orden
        orden.ID = ultimaOrdenProdID
        ultimaOrdenProdID++
        err := c.ShouldBindJSON(&ordenProd)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
            fmt.Println("Error al hacer bind del JSON:", err)
            return
        }

        ordenProd, err = h.s.CrearOrdenProducto(ordenProd)
        if err != nil {
            web.Failure(c, 500, errors.New("failed to create order product: " + err.Error()))
            return
        }

        // Devolver la relación creada con los detalles del producto
        web.Success(c, 200, ordenProd)
    }
}
/*func (h *ordenProductoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
        var ordenProd domain.OrdenProducto
        ordenProd.ID = ultimaOrdenProdID
        ultimaOrdenProdID++
        err := c.ShouldBindJSON(&ordenProd)
        if err != nil {
            web.Failure(c, 400, errors.New("invalid json: " + err.Error()))
            fmt.Println("Error al hacer bind del JSON:", err)
            return
        }

		ordenProd, err = h.s.CrearOrdenProducto(ordenProd)
		if err != nil {
			web.Failure(c, 500, errors.New("failed to create order product: " + err.Error()))
			return
		}

		// Devolver la relación creada
		web.Success(c, 200, ordenProd)
	}
}*/

// BuscarOrdenProducto busca una relación orden-producto por su ID
func (h *ordenProductoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		log.Printf("Buscando OrdenProducto con ID: %d", id)
		op, err := h.s.BuscaOrdenProducto(id)
		if err != nil {
			log.Printf("Error al buscar OrdenProducto: %v", err)
			web.Failure(c, 404, errors.New("No se encuentra"))
			return
		}
		web.Success(c, 200, op)
	}
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> FILTRA LA ORDEN PRODUCTO POR ID DE ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (h *ordenProductoHandler) BuscarPorIDOrden() gin.HandlerFunc {
    return func(c *gin.Context) {
        idOrdenParam := c.Param("idOrden")
        idOrden, err := strconv.Atoi(idOrdenParam)
        if err != nil {
            web.Failure(c, 400, errors.New("Invalid order ID"))
            return
        }
        ordenesProducto, err := h.s.BuscarOrdenesProductoPorIDOrden(idOrden)
        if err != nil {
            web.Failure(c, 500, fmt.Errorf("failed to fetch order products: %w", err))
            return
        }
        web.Success(c, 200, ordenesProducto)
    }
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODAS LAS ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (h *ordenProductoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios, err := h.s.BuscarTodasLasOrdenesProducto()
		if err != nil {
			web.Failure(c, 500, fmt.Errorf("error buscando todos los usuarios: %w", err))
			return
		}
		web.Success(c, 200, usuarios)
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
