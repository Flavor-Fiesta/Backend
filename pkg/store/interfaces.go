package store

import (

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type StoreInterfaceProducto interface {
	// Read devuelve un odontologo por su id
	BuscarProducto(id int) (domain.Producto, error)
	BuscarTodosLosProductos() ([]domain.Producto, error)
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
    PatchImagen(id int, updatedFields map[string]interface{}) error	
	ExistsByIDImagen(id int) bool

}

type StoreInterfaceUsuarios interface {

	CrearUsuario(usuario domain.Usuarios) error
	BuscarUsuario(id int) (domain.Usuarios, error)
	BuscarTodosLosUsuarios() ([]domain.Usuarios, error)
	UpdateUsuario(id int, p domain.Usuarios) error
	DeleteUsuario(id int) error
	ExistsByIDUsuario(id int) (bool, error)
    PatchUsuario(id int, updatedFields map[string]interface{}) error
}
type StoreInterfaceCarritos interface {

	CrearCarrito(carrito domain.Carrito) error
	UpdateCarrito(id int, p domain.Carrito) error
	DeleteCarrito(id int) error
	ExistsByIDCarrito(id int) bool
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


