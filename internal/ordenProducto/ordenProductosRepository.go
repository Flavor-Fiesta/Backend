package ordenProductos

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
    CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error)
    BuscaOrdenProducto(id int) (domain.OrdenProducto, error)
	UpdateOrdenProducto(id int, p domain.OrdenProducto) (domain.OrdenProducto, error)
}

type repository struct {
    storage store.StoreInterfaceOrdenProducto
}

func NewRepository(storage store.StoreInterfaceOrdenProducto) Repository {
    return &repository{storage: storage}
}

func (r *repository) CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error)  {
    err := r.storage.CrearOrdenProducto(op)
	if err != nil {
		log.Printf("Error al crear el producto %v: %v\n", op, err)
		return domain.OrdenProducto{}, fmt.Errorf("error creando producto: %w", err)
	}
	return op, nil
}

func (r *repository) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
    op, err := r.storage.BuscaOrdenProducto(id)
    if err != nil {
        return domain.OrdenProducto{}, errors.New("order product not found")
    }
    return op, nil
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
