package estados

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	CrearEstado(estado domain.Estado) (domain.Estado, error)
	BuscarTodosLosEstados() ([]domain.Estado, error)
	BuscarEstado(id int) (domain.Estado, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio para estados
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ESTADO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) CrearEstado(estado domain.Estado) (domain.Estado, error) {
	estadoCreado, err := s.r.CrearEstados(estado)
	if err != nil {
		return domain.Estado{}, err
	}
	return estadoCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE PRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// BuscarProducto busca un producto por su ID y devuelve también los datos de la imagen asociada
func (s *service) BuscarEstado(id int) (domain.Estado, error) {
	estado, err := s.r.BuscarEstado(id)
	if err != nil {
		return domain.Estado{}, err
	}
	return estado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTENER TODOS LOS ESTADOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodosLosEstados() ([]domain.Estado, error) {
	estados, err := s.r.BuscarTodosLosEstados()
	if err != nil {
		return nil, fmt.Errorf("error buscando todos los estados: %w", err)
	}
	return estados, nil
}
