package usuarios

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {

	BuscarUsuario(id int) (domain.Usuarios, error)
    BuscarTodosLosUsuarios() ([]domain.Usuarios, error)
	CrearUsuario(p domain.Usuarios) (domain.Usuarios, error)
	UpdateUsuario(id int, p domain.Usuarios) (domain.Usuarios, error)
	DeleteUsuario(id int) error

}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO USUARIO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// CrearProducto crea un nuevo usuario utilizando el repositorio y devuelve el usuario creado
func (s *service) CrearUsuario(p domain.Usuarios) (domain.Usuarios, error) {
    // Crear el producto utilizando el repositorio
    usuarioCreado, err := s.r.CrearUsuario(p)
    if err != nil {
        return domain.Usuarios{}, err
    }
    return usuarioCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE USUARIO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// BuscarProducto busca un usuario por su ID y devuelve también los datos de la imagen asociada
func (s *service) BuscarUsuario(id int) (domain.Usuarios, error) {
	p, err := s.r.BuscarUsuario(id)
	if err != nil {
		return domain.Usuarios{}, err
	}
	return p, nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODOS LOS USUARIOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodosLosUsuarios() ([]domain.Usuarios, error) {
	usuarios, err := s.r.BuscarTodosLosUsuarios()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los usuarios: %w", err)
	}
	return usuarios, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA  UN  PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateUsuario(id int, u domain.Usuarios) (domain.Usuarios, error) {
	// Llama directamente a la actualización en el repositorio
	return s.r.UpdateUsuario(id, u)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) PatchUsuario(id int, updatedFields map[string]interface{}) (domain.Usuarios, error) {
    // Obtener el paciente por su ID
    usuario, err := s.r.BuscarUsuario(id)
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
            if codigo, ok := value.(string); ok {
                usuario.Email = codigo
            }
        case "Telefono":
            if telefono, ok := value.(string); ok {
                usuario.Telefono = telefono
            }
        case "Password":
            if password, ok := value.(string); ok {
                usuario.Password = password
            }

        // Puedes añadir más campos aquí según sea necesario
        default:
            return domain.Usuarios{}, fmt.Errorf("campo desconocido: %s", field)
        }
    }

    // Actualizar el producto en el repositorio
    updatedUsuario, err := s.r.UpdateUsuario(id, usuario)
    if err != nil {
        return domain.Usuarios{}, err
    }

    return updatedUsuario, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) DeleteUsuario(id int) error {
    err := s.r.DeleteUsuario(id)
    if err != nil {
        return err
    }
    return nil
}