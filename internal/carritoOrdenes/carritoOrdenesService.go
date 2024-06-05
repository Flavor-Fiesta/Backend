package carritoOrdenes

import (
	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	GetCarritoOrdenesByID(id int) (domain.CarritosOrden, error)
	CrearCarritoOrdenes(p domain.CarritosOrden) (domain.CarritosOrden, error)
	DeleteCarritoOrdenes(id int) error
	UpdateCarritoOrdenes(id int, p domain.CarritosOrden) (domain.CarritosOrden, error)
}

type service struct {
	r Repository
}

// NewService creates a new service for carritoOrdenes
func NewService(r Repository) Service {
	return &service{r}
}

// CrearCarritoOrdenes creates a new carritoOrden
func (s *service) CrearCarritoOrdenes(p domain.CarritosOrden) (domain.CarritosOrden, error) {
	p, err := s.r.CrearCarritoOrdenes(p)
	if err != nil {
		return domain.CarritosOrden{}, err
	}
	return p, nil
}

func (s *service) GetCarritoOrdenesByID(id int) (domain.CarritosOrden, error) {
	p, err := s.r.GetCarritoOrdenesByID(id)
	if err != nil {
		return domain.CarritosOrden{}, err
	}
	return p, nil
}

// UpdateCarritoOrdenes updates an existing carritoOrden
func (s *service) UpdateCarritoOrdenes(id int, u domain.CarritosOrden) (domain.CarritosOrden, error) {
	return s.r.UpdateCarritoOrdenes(id, u)
}

// DeleteCarritoOrdenes deletes a carritoOrden by ID
func (s *service) DeleteCarritoOrdenes(id int) error {
	err := s.r.DeleteCarritoOrdenes(id)
	if err != nil {
		return err
	}
	return nil
}
