package store

import (

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type StoreInterfaceProducto interface {
	BuscarProducto(id int) (domain.Producto, error)
	BuscarTodosLosProductos() ([]domain.Producto, error)
	CrearProducto(producto domain.Producto) error
	UpdateProducto(id int, p domain.Producto) error
	DeleteProducto(id int) error
	ExistsByID(id int) bool
    Patch(id int, updatedFields map[string]interface{}) (domain.Producto, error)
	ObtenerNombreCategoria(id int) (string, error) // Añadir este método
}

type StoreInterfaceImagenes interface {
//	CrearImagen(imagen domain.Imagen) error
	CrearImagenes(imagenes []domain.Imagen) error
	BuscarImagen(id int) (domain.Imagen, error)
	UpdateImagen(id int, p domain.Imagen) error
	DeleteImagen(id int) error
	PatchImagen(id int, updatedFields map[string]interface{}) error	
	ExistsByIDImagen(id int) bool
}

type StoreInterfaceUsuarios interface {
	ExisteEmail(email string) (bool, error) 
	ExisteEmail2(email string) (domain.Usuarios, error)        
    ExisteCelular(celular string) (bool, error)
	CrearUsuario(usuario domain.Usuarios) error
	BuscarUsuario(id int) (domain.Usuarios, error)
	//BuscarUsuarioPorEmailYPassword(email, password string) (domain.Usuarios, error)
	BuscarUsuarioPorEmailYPassword(email, password string) (bool, error)
	BuscarUsuarioPorEmailYPassword2(email, password string) (domain.Usuarios, error)
	BuscarUsuarioPorEmailYPassword3(email, password string) (bool, error, domain.Usuarios)
	BuscarTodosLosUsuarios() ([]domain.Usuarios, error)

	DeleteUsuario(id int) error
	ExistsByIDUsuario(id int) (bool, error)

	Update(usuario domain.Usuarios) error
	UpdatePassword(id int, newPassword string) (domain.Usuarios, error)
	PatchUsuario(id int, updatedFields map[string]interface{}) error
}

type StoreInterfaceCategorias interface {
	CrearCategoria(categoria domain.Categoria) error
	BuscarCategoria(id int) (domain.Categoria, error)
	BuscarTodosLasCategorias() ([]domain.Categoria, error)
	DeleteCategoria(id int) error
	ExistsByIDCategoria(id int) (bool, error)
	Update(categoria domain.Categoria) error
	PatchCategoria(id int, updatedFields map[string]interface{}) error
}

type StoreInterfaceRoles interface{
	CrearRol (rol domain.Rol) error
	BuscarTodosLosRoles() ([]domain.Rol, error)
	CambiarRol(usuarioID int, nuevoRol string) error
//	BuscarRol(id int) (domain.Rol, error)
}

type StoreInterfaceEstados interface{
    CrearEstados (estado domain.Estado) error
    BuscarTodosLosEstados() ([]domain.Estado, error)
    BuscarEstado(id int) (domain.Estado, error)
}

type StoreInterfaceOrdenes interface {
	CrearOrden(orden domain.Orden) error
	BuscarOrden(id int) (domain.Orden, error)
	UpdateOrden(id int, orden domain.Orden) error
	DeleteOrden(id int) error
	ExistsByIDOrden(id int) bool
	BuscarOrdenPorUsuarioYEstado(UserID, Estado string) (bool, error)
	BuscarOrdenPorUsuarioYEstado2(UserID, Estado string) (bool, error, domain.Orden)
}

type StoreInterfaceOrdenProducto interface{
    CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error)
    BuscaOrdenProducto(id int) (domain.OrdenProducto, error)
    UpdateOrdenProducto(id int, p domain.OrdenProducto) error
    ExistsByID(id int) bool
	BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error)
	BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error)

}

type StoreInterfaceFavoritos interface{
	AgregarFavorito(favorito domain.Favoritos) error

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