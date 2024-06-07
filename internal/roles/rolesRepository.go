package roles

import (
//	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {


    CrearRol(p domain.Rol) (domain.Rol, error)
	BuscarTodosLosRoles() ([]domain.Rol, error)
	CambiarRol(usuarioID int, nuevoRol string) error
//	BuscarImagen(id int) (domain.Imagen, error)
//	UpdateImagen(id int, p domain.Imagen) (domain.Imagen, error)
//	DeleteImagen(id int) error
}

type repository struct {
	storage store.StoreInterfaceRoles
	
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceRoles) Repository {
	return &repository{storage}
}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR ROL >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearRol(p domain.Rol) (domain.Rol, error) {
	err := r.storage.CrearRol(p)
	if err != nil {
		log.Printf("Error al crear el producto %v: %v\n", p, err)
		return domain.Rol{}, fmt.Errorf("error creando producto: %w", err)
	}
	return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS ROLES >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (r *repository) BuscarTodosLosRoles() ([]domain.Rol, error) {
	roles, err := r.storage.BuscarTodosLosRoles()
	if err != nil {
		return nil, fmt.Errorf("error buscando todas las categorias: %w", err)
	}
	return roles, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CAMBIAR DE ROL >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

/*func (r *repository) CambiarRol(usuarioID int, nuevoRol string) error {
    return r.storage.CambiarRol(usuarioID, nuevoRol)
}
*/

func (r *repository) CambiarRol(usuarioID int, nuevoRol string) error {
    err := r.storage.CambiarRol(usuarioID, nuevoRol)
    if err != nil {
        return fmt.Errorf("error cambiando rol: %w", err)
    }
    return nil
}

