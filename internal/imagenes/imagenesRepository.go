package imagenes

import (
	"errors"
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {

	BuscarImagen(id int) (domain.Imagen, error)
	CrearImagen(p domain.Imagen) (domain.Imagen, error)
	UpdateImagen(id int, p domain.Imagen) (domain.Imagen, error)
	DeleteImagen(id int) error
}

type repository struct {
	storage store.StoreInterfaceImagenes
	
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceImagenes) Repository {
	return &repository{storage}
}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearImagen(p domain.Imagen) (domain.Imagen, error) {
	// Crear el odontólogo en el almacenamiento
	err := r.storage.CrearImagen(p)
	if err != nil {
		return domain.Imagen{}, errors.New("error creando paciente")
	}
	return p, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR IMAGEN POR ID >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarImagen(id int) (domain.Imagen, error) {
	product, err := r.storage.BuscarImagen(id)
	if err != nil {
		return domain.Imagen{}, errors.New("imagen not found")
	}
	return product, nil

}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) UpdateImagen(id int, p domain.Imagen) (domain.Imagen, error) {
	// Verificar si la imagen existe por su ID
	if !r.storage.ExistsByIDImagen(id) {
		return domain.Imagen{}, fmt.Errorf("Imagen con ID %d no encontrada", id)
	}

	// Actualizar la imagen en el almacenamiento
	err := r.storage.UpdateImagen(id, p)
	if err != nil {
		return domain.Imagen{}, err
	}

	return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) Patch(id int, updatedFields map[string]interface{}) (domain.Imagen, error) {
    // Obtener la imagen por su ID
    imagen, err := r.BuscarImagen(id)
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
            return domain.Imagen{}, errors.New("campo desconocido: " + field)
        }
    }

    // Actualizar la imagen en el almacenamiento
    updatedImagen, err := r.UpdateImagen(id, imagen)
    if err != nil {
        return domain.Imagen{}, err
    }

    return updatedImagen, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) DeleteImagen(id int) error {
	err := r.storage.DeleteImagen(id)
	if err != nil {
		return err
	}
	return nil
}
