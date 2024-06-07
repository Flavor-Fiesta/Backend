package imagenes

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {

	BuscarImagen(id int) (domain.Imagen, error)
	CrearImagenes(imagenes []domain.Imagen) error
	DeleteImagen(id int) error
	UpdateImagen(id int, p domain.Imagen) (domain.Imagen, error)

	
}



type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) CrearImagenes(imagenes []domain.Imagen) error {
    err := s.r.CrearImagenes(imagenes)
    if err != nil {
        return err
    }
    return nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE IMAGEN POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarImagen(id int) (domain.Imagen, error) {
	p, err := s.r.BuscarImagen(id)
	if err != nil {
		return domain.Imagen{}, err
	}
	return p, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA  UNA  IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateImagen(id int, u domain.Imagen) (domain.Imagen, error) {
	// Llama directamente a la actualización en el repositorio
	return s.r.UpdateImagen(id, u)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) Patch(id int, updatedFields map[string]interface{}) (domain.Imagen, error) {
    // Obtener la imagen por su ID
    imagen, err := s.r.BuscarImagen(id)
    if err != nil {
        return domain.Imagen{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "Titulo":
            if titulo, ok := value.(string); ok {
                imagen.Titulo = titulo
            }
        case "Url":
            if url, ok := value.(string); ok {
                imagen.Url = url
            }
        // Puedes añadir más campos aquí según sea necesario
        default:
            return domain.Imagen{}, fmt.Errorf("campo desconocido: %s", field)
        }
    }

    // Actualizar la imagen en el repositorio
    updatedImagen, err := s.r.UpdateImagen(id, imagen)
    if err != nil {
        return domain.Imagen{}, err
    }

    return updatedImagen, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) DeleteImagen(id int) error {
	err := s.r.DeleteImagen(id)
	if err != nil {
		return err
	}
	return nil
}

/*func (s *service) BuscarProducto(id int) (domain.Producto, bool, error) {
  // Llama al método del repositorio que busca un producto por su ID
  producto, err := s.r.BuscarProductoPorID(id)
  if err != nil {
	  // Manejar el error, por ejemplo, loguearlo o devolver un error
	  return domain.Producto{}, false, err
  }

  // Verificar si se encontró un producto con el ID proporcionado
  if producto.ID == 0 {
	  // Si no se encontró ningún producto con ese ID, devuelve false y un producto vacío
	  return domain.Producto{}, false, nil
  }

  // Si se encontró un producto con el ID proporcionado, devuelve el producto y true
  return producto, true, nil
}*/