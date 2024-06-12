package ordenProductos

import (
	"errors"
	"fmt"
	"log"

	//	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
    CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error)
    BuscaOrdenProducto(id int) (domain.OrdenProducto, error)
	UpdateOrdenProducto(id int, p domain.OrdenProducto) (domain.OrdenProducto, error)
	BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error)
	BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error)
}

type repository struct {
    storage store.StoreInterfaceOrdenProducto
}

func NewRepository(storage store.StoreInterfaceOrdenProducto) Repository {
    return &repository{storage: storage}
}


func (r *repository) CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error) {
    ordenProductoCreado, err := r.storage.CrearOrdenProducto(op)
    if err != nil {
        return domain.OrdenProducto{}, err
    }
    return ordenProductoCreado, nil
}
/*func (r *repository) CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error)  {
    err := r.storage.CrearOrdenProducto(op)
	if err != nil {
		log.Printf("Error al crear el producto %v: %v\n", op, err)
		return domain.OrdenProducto{}, fmt.Errorf("error creando producto: %w", err)
	}
	return op, nil
}*/

func (r *repository) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
    log.Printf("Repository: Buscando OrdenProducto con ID: %d", id)
    op, err := r.storage.BuscaOrdenProducto(id)
    if err != nil {
        log.Printf("Repository: Error al buscar OrdenProducto: %v", err)
        return domain.OrdenProducto{}, errors.New("order product not found")
    }
    return op, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR ORDEN PRODUCTO POR ID DE ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error) {
    log.Printf("Repository: Buscando OrdenesProducto con ID de orden: %d", idOrden)
    ordenesProducto, err := r.storage.BuscarOrdenesProductoPorIDOrden(idOrden)
    if err != nil {
        log.Printf("Repository: Error al buscar OrdenesProducto por ID de orden: %v", err)
        return nil, fmt.Errorf("error searching order products by order ID: %w", err)
    }
    return ordenesProducto, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODAS LAS ORDEN PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error) {
	usuarios, err := r.storage.BuscarTodasLasOrdenesProducto()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los usuarios: %w", err)
	}
	return usuarios, nil
}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) UpdateOrdenProducto(id int, p domain.OrdenProducto) (domain.OrdenProducto, error) {
	// Verificar si el producto existe por su ID
	if !r.storage.ExistsByID(id) {
		return domain.OrdenProducto{}, fmt.Errorf("Producto con ID %d no encontrado", id)
	}
	// Actualizar el producto en el almacenamiento
	err := r.storage.UpdateOrdenProducto(id, p)
	if err != nil {
		return domain.OrdenProducto{}, err
	}

	return p, nil
}
