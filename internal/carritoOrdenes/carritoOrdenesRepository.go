package carritoOrdenes

import (
	"errors"
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
	GetCarritoOrdenesByID(id int) (domain.CarritosOrden, error)
	CrearCarritoOrdenes(carrito domain.CarritosOrden) (domain.CarritosOrden, error)
	UpdateCarritoOrdenes(id int, p domain.CarritosOrden) (domain.CarritosOrden, error)
	DeleteCarritoOrdenes(id int) error
}

type repository struct {
	storage store.StoreInterfaceCarritoOrdenes
}

// NewRepository creates a new repository
func NewRepository(storage store.StoreInterfaceCarritoOrdenes) Repository {
	return &repository{storage}
}

func (r *repository) GetCarritoOrdenesByID(id int) (domain.CarritosOrden, error) {
	carritoOrden, err := r.storage.GetCarritoOrdenesByID(id)
	if err != nil {
		return domain.CarritosOrden{}, err
	}
	return carritoOrden, nil
}

func (r *repository) CrearCarritoOrdenes(carrito domain.CarritosOrden) (domain.CarritosOrden, error) {
	err := r.storage.CrearCarritoOrdenes(carrito)
	if err != nil {
		return domain.CarritosOrden{}, errors.New("error creating carritoOrden")
	}
	return carrito, nil
}

func (r *repository) UpdateCarritoOrdenes(id int, p domain.CarritosOrden) (domain.CarritosOrden, error) {
	if !r.storage.ExistsByIDCarritoOrdenes(id) {
		return domain.CarritosOrden{}, fmt.Errorf("CarritoOrden with ID %d not found", id)
	}

	err := r.storage.UpdateCarritoOrdenes(id, p)
	if err != nil {
		return domain.CarritosOrden{}, err
	}

	return p, nil
}

func (r *repository) DeleteCarritoOrdenes(id int) error {
	err := r.storage.DeleteCarritoOrdenes(id)
	if err != nil {
		return err
	}
	return nil
}
