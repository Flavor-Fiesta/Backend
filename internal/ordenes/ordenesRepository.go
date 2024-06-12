package ordenes

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
	BuscarOrden(id int) (domain.Orden, error)
	CrearOrden(o domain.Orden) (domain.Orden, error)
	UpdateOrden(id int, o domain.Orden) (domain.Orden, error)
	DeleteOrden(id int) error
	BuscarOrdenPorUsuarioYEstado(UserID, Estado string) (bool, error)
	BuscarOrdenPorUsuarioYEstado2(UserID, Estado string) (bool, error, domain.Orden)
}

type repository struct {
	storage store.StoreInterfaceOrdenes
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterfaceOrdenes) Repository {
	return &repository{storage}
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearOrden(o domain.Orden) (domain.Orden, error) {
	// Crear la orden en el almacenamiento
	err := r.storage.CrearOrden(o)
	if err != nil {
		return domain.Orden{}, errors.New("error creando orden")
	}
	return o, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR ORDEN POR ID >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarOrden(id int) (domain.Orden, error) {
	orden, err := r.storage.BuscarOrden(id)
	if err != nil {
		return domain.Orden{}, errors.New("orden not found")
	}
	return orden, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR ORDEN POR ID USUARIO Y ESTADO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarOrdenPorUsuarioYEstado(UserID, Estado string) (bool, error) {
    exists, err := r.storage.BuscarOrdenPorUsuarioYEstado(UserID, Estado)
    if err != nil {
        return false, errors.New("order not found")
    }
    return exists, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR ORDEN POR ID USUARIO Y ESTADO CON TODOS LOS DATOS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarOrdenPorUsuarioYEstado2(UserID, Estado string) (bool, error, domain.Orden) {
    exists, err, orden := r.storage.BuscarOrdenPorUsuarioYEstado2(UserID, Estado)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, errors.New("orden not found"), domain.Orden{}
        }
        return false, err, domain.Orden{}
    }
    return exists, nil, orden
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZAR ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) UpdateOrden(id int, o domain.Orden) (domain.Orden, error) {
	// Verificar si la orden existe por su ID
	if !r.storage.ExistsByIDOrden(id) {
		return domain.Orden{}, fmt.Errorf("Orden con ID %d no encontrada", id)
	}

	// Actualizar la orden en el almacenamiento
	err := r.storage.UpdateOrden(id, o)
	if err != nil {
		return domain.Orden{}, err
	}

	return o, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) Patch(id int, updatedFields map[string]interface{}) (domain.Orden, error) {
    // Obtener la orden por su ID
    orden, err := r.BuscarOrden(id)
    if err != nil {
        return domain.Orden{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "FechaOrden":
            if fechaOrden, ok := value.(string); ok {
                orden.FechaOrden = fechaOrden
            }
        case "Total":
            if total, ok := value.(float64); ok {
                orden.Total = total
            }
        // Puedes añadir más campos aquí según sea necesario
        default:
            return domain.Orden{}, errors.New("campo desconocido: " + field)
        }
    }

    // Actualizar la orden en el almacenamiento
    updatedOrden, err := r.UpdateOrden(id, orden)
    if err != nil {
        return domain.Orden{}, err
    }

    return updatedOrden, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) DeleteOrden(id int) error {
	err := r.storage.DeleteOrden(id)
	if err != nil {
		return err
	}
	return nil
}