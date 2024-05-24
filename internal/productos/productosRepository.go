package productos

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {

	BuscarProducto(id int) (domain.Producto, error)
    BuscarTodosLosProductos() ([]domain.Producto, error)
	CrearProducto(p domain.Producto) (domain.Producto, error)
	UpdateProducto(id int, p domain.Producto) (domain.Producto, error)
	DeleteProducto(id int) error
}

type repository struct {
	storage store.StoreInterfaceProducto
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceProducto) Repository {
    return &repository{storage: storage}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearProducto(p domain.Producto) (domain.Producto, error) {
    // Crear el producto en el almacenamiento
    err := r.storage.CrearProducto(p)
    if err != nil {
        // Agregar registro de error detallado
        log.Printf("Error al crear el producto %v: %v\n", p, err)
        return domain.Producto{}, fmt.Errorf("error creando producto: %w", err)
    }
    return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarProducto(id int) (domain.Producto, error) {
	producto, err := r.storage.BuscarProducto(id)
	if err != nil {
		return domain.Producto{}, errors.New("product not found")
	}
	return producto, nil

}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS PRODUCTOS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (r *repository) BuscarTodosLosProductos() ([]domain.Producto, error) {
	productos, err := r.storage.BuscarTodosLosProductos()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los productos: %w", err)
	}
	return productos, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) UpdateProducto(id int, p domain.Producto) (domain.Producto, error) {
	// Verificar si el producto existe por su ID
	if !r.storage.ExistsByID(id) {
		return domain.Producto{}, fmt.Errorf("Producto con ID %d no encontrado", id)
	}
	// Actualizar el producto en el almacenamiento
	err := r.storage.UpdateProducto(id, p)
	if err != nil {
		return domain.Producto{}, err
	}

	return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) Patch(id int, updatedFields map[string]interface{}) (domain.Producto, error) {
    // Obtener el producto por su ID
    producto, err := r.BuscarProducto(id)
    if err != nil {
        return domain.Producto{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "Nombre":
            if nombre, ok := value.(string); ok {
                producto.Nombre = nombre
            }
        case "Codigo":
            if codigo, ok := value.(string); ok {
                producto.Codigo = codigo
            }
        case "Categoria":
            if categoria, ok := value.(string); ok {
                producto.Categoria = categoria
            }
        case "Fecha_De_Alta":
            if fecha_de_alta, ok := value.(string); ok {
                producto.FechaDeAlta = fecha_de_alta
            }
        case "Fecha_De_Vencimiento":
            if fecha_de_vencimiento, ok := value.(string); ok {
                producto.FechaDeVencimiento = fecha_de_vencimiento
            }
            
        }
    }

    // Actualizar el producto en el almacenamiento
    updatedProducto, err := r.UpdateProducto(id, producto)
    if err != nil {
        return domain.Producto{}, err
    }

    return updatedProducto, nil
}


// DeleteProducto elimina un producto del repositorio
func (r *repository) DeleteProducto(id int) error {
    err := r.storage.DeleteProducto(id)
    if err != nil {
        return err
    }
    return nil
}