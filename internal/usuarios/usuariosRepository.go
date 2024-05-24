package usuarios

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {

	BuscarUsuario(id int) (domain.Usuarios, error)
    BuscarTodosLosUsuarios() ([]domain.Usuarios, error)
	CrearUsuario(p domain.Usuarios) (domain.Usuarios, error)
	UpdateUsuario(id int, p domain.Usuarios) (domain.Usuarios, error)
	DeleteUsuario(id int) error
}

type repository struct {
	storage store.StoreInterfaceUsuarios
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceUsuarios) Repository {
    return &repository{storage}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearUsuario(p domain.Usuarios) (domain.Usuarios, error) {
    // Crear el producto en el almacenamiento
    err := r.storage.CrearUsuario(p)
    if err != nil {
        // Agregar registro de error detallado
        log.Printf("Error al crear el usuario %v: %v\n", p, err)
        return domain.Usuarios{}, fmt.Errorf("error creando usuario: %w", err)
    }
    return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarUsuario(id int) (domain.Usuarios, error) {
	usuario, err := r.storage.BuscarUsuario(id)
	if err != nil {
		return domain.Usuarios{}, errors.New("usuario not found")
	}
	return usuario, nil

}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (r *repository) BuscarTodosLosUsuarios() ([]domain.Usuarios, error) {
	usuarios, err := r.storage.BuscarTodosLosUsuarios()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los usuarios: %w", err)
	}
	return usuarios, nil
}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) UpdateUsuario(id int, p domain.Usuarios) (domain.Usuarios, error) {
    // Verificar si el usuario existe por su ID
    exists, err := r.storage.ExistsByIDUsuario(id)
    if err != nil {
        return domain.Usuarios{}, fmt.Errorf("error al verificar si el usuario existe: %v", err)
    }
    if !exists {
        return domain.Usuarios{}, fmt.Errorf("usuario con ID %d no encontrado", id)
    }

    // Actualizar el usuario en el almacenamiento
    err = r.storage.UpdateUsuario(id, p)
    if err != nil {
        return domain.Usuarios{}, fmt.Errorf("error al actualizar el usuario: %v", err)
    }

    return p, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH USUARIO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) PatchUsuario(id int, updatedFields map[string]interface{}) (domain.Usuarios, error) {
    // Obtener el usuario por su ID
    usuario, err := r.BuscarUsuario(id)
    if err != nil {
        return domain.Usuarios{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "Nombre":
            if nombre, ok := value.(string); ok {
                usuario.Nombre = nombre
            }
        case "Email":
            if email, ok := value.(string); ok {
                usuario.Email = email
            }
        case "Telefono":
            if telefono, ok := value.(string); ok {
                usuario.Telefono = telefono
            }
        case "Password":
            if password, ok := value.(string); ok {
                usuario.Password = password
            }
            
        }
    }

// Actualizar la imagen en el almacenamiento
updatedImagen, err := r.UpdateUsuario(id, usuario)
if err != nil {
    return domain.Usuarios{}, err
}

return updatedImagen, nil
}


// DeleteProducto elimina un producto del repositorio
func (r *repository) DeleteUsuario(id int) error {
    err := r.storage.DeleteUsuario(id)
    if err != nil {
        return err
    }
    return nil
}