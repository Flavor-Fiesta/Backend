package store

import (

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type StoreInterfaceProducto interface {
	// Read devuelve un odontologo por su id
	BuscarProducto(id int) (domain.Producto, error)
	// Create agrega un nuevo odontologo
	CrearProducto(odonto domain.Producto) error
	// Update actualiza un odontologo
	UpdateProducto(id int, p domain.Producto) error
	// Delete elimina un odontologo
	DeleteProducto(id int) error
	//
	ExistsByID(id int) bool
	// Patch
    Patch(id int, updatedFields map[string]interface{}) error
}

type StoreInterfaceImagenes interface {
	// Create agrega un nuevo odontologo
	CrearImagen(imagen domain.Imagen) error
	// Read devuelve un odontologo por su id
	BuscarImagen(id int) (domain.Imagen, error)
	// Update actualiza un odontologo
	UpdateImagen(id int, p domain.Imagen) error
	// Delete elimina un odontologo
	DeleteImagen(id int) error
	ExisteProductoParaImagen(id int) (bool, error)
    PatchImagen(id int, updatedFields map[string]interface{}) error	

}

/*type StoreInterfaceProductoImagen interface {
	// Read devuelve un paciente por su id
	BuscarProductoImagen(id int) (domain.ProductoImagen, error)
	// Create agrega un nuevo turno
	CrearProductoImagen(turno domain.ProductoImagen) error
	// Update actualiza un paciente
	UpdateProductoImagen(id int, p domain.ProductoImagen) error
	ExistsByIDProductoImagen(id int) bool
	// Delete elimina un paciente
	DeleteProductoImagen(id int) error
}*/