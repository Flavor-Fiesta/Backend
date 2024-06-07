package estados

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
	CrearEstados(estado domain.Estado) (domain.Estado, error)
	BuscarTodosLosEstados() ([]domain.Estado, error)
	BuscarEstado(id int) (domain.Estado, error)
}

type repository struct {
	storage store.StoreInterfaceEstados
}

// NewRepository crea un nuevo repositorio para estados
func NewRepository(storage store.StoreInterfaceEstados) Repository {
	return &repository{storage}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR ESTADO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) CrearEstados(estado domain.Estado) (domain.Estado, error) {
	err := r.storage.CrearEstados(estado)
	if err != nil {
		log.Printf("Error al crear el estado %v: %v\n", estado, err)
		return domain.Estado{}, fmt.Errorf("error creando estado: %w", err)
	}
	return estado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS ESTADOS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarTodosLosEstados() ([]domain.Estado, error) {
	estados, err := r.storage.BuscarTodosLosEstados()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los estados: %w", err)
	}
	return estados, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR ESTADOS >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (r *repository) BuscarEstado(id int) (domain.Estado, error) {
	estado, err := r.storage.BuscarEstado(id)
	if err != nil {
		return domain.Estado{}, errors.New("producto not found")
	}
	return estado, nil
}
