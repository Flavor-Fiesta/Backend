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
	ExisteProductoParaImagen(id int) (bool, error)
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
		// Crear la imagen en el almacenamiento
		err := r.storage.CrearImagen(p)
		if err != nil {
			return domain.Imagen{}, errors.New("error creando odontologo")
		}
		return p, nil
	}

	func (r *repository) ExisteProductoParaImagen(id int) (bool, error) {
		// Llamar a la función en tu SQL store para verificar la existencia de la imagen por su ID
		exists, err := r.storage.ExisteProductoParaImagen(id)
		if err != nil {
			fmt.Print("ACA SI LLEGA Y PASA")
			return false, err // Devuelve el error si ocurre alguno
		}
		fmt.Print("REVISANDO SI TRAE ALGO: ", exists)
		return exists, nil // Devuelve el resultado de la verificación
		
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
	exists, err := r.storage.ExisteProductoParaImagen(id)
	if err != nil {
		return domain.Imagen{}, err // Devuelve el error si ocurre alguno
	}
	if !exists {
		return domain.Imagen{}, fmt.Errorf("Odontologo con ID %d no encontrado", id)
	}

	// Actualizar la imagen en el almacenamiento
	err = r.storage.UpdateImagen(id, p)
	if err != nil {
		return domain.Imagen{}, err
	}
	// Retorna la imagen actualizada
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
