package categorias

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {

	CrearCategoria(p domain.Categoria) (domain.Categoria, error)
	BuscarCategoria(id int) (domain.Categoria, error)
	BuscarTodosLasCategorias() ([]domain.Categoria, error)
	DeleteCategoria(id int) error
//	ExistsByIDCategoria(id int) (bool, error)


    Update(id int, p domain.Categoria) (domain.Categoria, error)
}

type repository struct {
	storage store.StoreInterfaceCategorias
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceCategorias) Repository {
    return &repository{storage: storage}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearCategoria(p domain.Categoria) (domain.Categoria, error) {
    // Crear el producto en el almacenamiento
    err := r.storage.CrearCategoria(p)
    if err != nil {
        // Agregar registro de error detallado
        log.Printf("Error al crear el usuario %v: %v\n", p, err)
        return domain.Categoria{}, fmt.Errorf("error creando usuario: %w", err)
    }
    return p, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarCategoria(id int) (domain.Categoria, error) {
	producto, err := r.storage.BuscarCategoria(id)
	if err != nil {
		return domain.Categoria{}, errors.New("categoria not found")
	}
	return producto, nil

}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODAS LAS CATEGORIAS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (r *repository) BuscarTodosLasCategorias() ([]domain.Categoria, error) {
	categorias, err := r.storage.BuscarTodosLasCategorias()
	if err != nil {
		return nil, fmt.Errorf("error buscando todas las categorias: %w", err)
	}
	return categorias, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) Update(id int, p domain.Categoria) (domain.Categoria, error) {

	err := r.storage.Update(p)
	if err != nil {
		return domain.Categoria{}, errors.New("error updating product")
	}
	return p, nil
}





//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) Patch(id int, updatedFields map[string]interface{}) (domain.Categoria, error) {
    // Obtener el odontólogo por su ID
    categoria, err := r.BuscarCategoria(id)
    if err != nil {
        return domain.Categoria{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "Nombre":
            if nombre, ok := value.(string); ok {
                categoria.Nombre = nombre
            }
        case "Descripcion":
            if descripcion, ok := value.(string); ok {
                categoria.Descripcion = descripcion
            }

        // Puedes añadir más campos aquí según sea necesario
        default:
            return domain.Categoria{}, errors.New("campo desconocido: " + field)
        }
    }

    // Actualizar el odontólogo en el almacenamiento
    updatedUsuario, err := r.Update(id, categoria)
    if err != nil {
        return domain.Categoria{}, err
    }

    return updatedUsuario, nil
}


//ELIMINAR

// DeleteCategoria elimina una categoria del repositorio
func (r *repository) DeleteCategoria(id int) error {
    err := r.storage.DeleteCategoria(id)
    if err != nil {
        return err
    }
    return nil
}