package ordenes

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	BuscarOrden(id int) (domain.Orden, error)
	CrearOrden(p domain.Orden) (domain.Orden, error)
	DeleteOrden(id int) error
	UpdateOrden(id int, p domain.Orden) (domain.Orden, error)
	Patch(id int, updatedFields map[string]interface{}) (domain.Orden, error)
}

type service struct {
	r Repository
}

// NewService creates a new service for ordenes
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UNA NUEVA ORDEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) CrearOrden(p domain.Orden) (domain.Orden, error) {
	p, err := s.r.CrearOrden(p)
	if err != nil {
		return domain.Orden{}, err
	}
	return p, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE ORDEN POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarOrden(id int) (domain.Orden, error) {
	p, err := s.r.BuscarOrden(id)
	if err != nil {
		return domain.Orden{}, err
	}
	return p, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UNA ORDEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateOrden(id int, u domain.Orden) (domain.Orden, error) {
	return s.r.UpdateOrden(id, u)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) Patch(id int, updatedFields map[string]interface{}) (domain.Orden, error) {
	orden, err := s.r.BuscarOrden(id)
	if err != nil {
		return domain.Orden{}, err
	}

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

		default:
			return domain.Orden{}, fmt.Errorf("campo desconocido: %s", field)
		}
	}

	updatedOrden, err := s.r.UpdateOrden(id, orden)
	if err != nil {
		return domain.Orden{}, err
	}

	return updatedOrden, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR ORDEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) DeleteOrden(id int) error {
	err := s.r.DeleteOrden(id)
	if err != nil {
		return err
	}
	return nil
}