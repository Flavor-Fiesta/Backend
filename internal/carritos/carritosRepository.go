package carritos

import (
	"errors"
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
	CrearCarrito(p domain.Carrito) (domain.Carrito, error)
	UpdateCarrito(id int, p domain.Carrito) (domain.Carrito, error)
	DeleteCarrito(id int) error
}

type repository struct {
	storage store.StoreInterfaceCarritos
}

// NewRepository creates a new repository
func NewRepository(storage store.StoreInterfaceCarritos) Repository {
	return &repository{storage}
}

func (r *repository) CrearCarrito(p domain.Carrito) (domain.Carrito, error) {
	err := r.storage.CrearCarrito(p)
	if err != nil {
		return domain.Carrito{}, errors.New("error creating carrito")
	}
	return p, nil
}


// UpdateCarrito updates an existing carrito
func (r *repository) UpdateCarrito(id int, p domain.Carrito) (domain.Carrito, error) {
	if !r.storage.ExistsByIDCarrito(id) {
		return domain.Carrito{}, fmt.Errorf("Carrito with ID %d not found", id)
	}

	err := r.storage.UpdateCarrito(id, p)
	if err != nil {
		return domain.Carrito{}, err
	}

	return p, nil
}

// DeleteCarrito deletes a carrito by ID
func (r *repository) DeleteCarrito(id int) error {
	err := r.storage.DeleteCarrito(id)
	if err != nil {
		return err
	}
	return nil
}
