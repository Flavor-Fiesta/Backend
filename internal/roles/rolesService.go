package roles

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {


	CrearRol(p domain.Rol) (domain.Rol, error)
	BuscarTodosLosRoles() ([]domain.Rol, error)
	CambiarRol(usuarioID int, nuevoRol string) error
//	UpdateProducto(id int, p domain.Producto) (domain.Producto, error)
//	DeleteProducto(id int) error
//  ObtenerNombreCategoria(id int) (string, error) // Nuevo método agregado
    

}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ROL <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// CrearRol crea un nuevo producto utilizando el repositorio y devuelve el producto creado
func (s *service) CrearRol(p domain.Rol) (domain.Rol, error) {
    // Crear el producto utilizando el repositorio
    rolCreado, err := s.r.CrearRol(p)
    if err != nil {
        return domain.Rol{}, err
    }
    return rolCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODOS LOS ROLES <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodosLosRoles() ([]domain.Rol, error) {
    roles, err := s.r.BuscarTodosLosRoles()
    if err != nil {
        return nil, fmt.Errorf("error buscando todas las categorias: %w", err)
    }
    return roles, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CAMBIAR DE ROL >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
/*func (s *service) CambiarRol(usuarioID int, nuevoRol string) error {
    // Llamar al repositorio para cambiar el rol del usuario
    return s.r.CambiarRol(usuarioID, nuevoRol)
}*/
func (s *service) CambiarRol(usuarioID int, nuevoRol string) error {
    err := s.r.CambiarRol(usuarioID, nuevoRol)
    if err != nil {
        return err
    }
    return nil
}