package ordenProductos

import (

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	CrearOrdenProducto(p domain.OrdenProducto) (domain.OrdenProducto, error)
	BuscaOrdenProducto(id int) (domain.OrdenProducto, error)
	UpdateOrdenProducto(id int, p domain.OrdenProducto) (domain.OrdenProducto, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) CrearOrdenProducto(op domain.OrdenProducto) (domain.OrdenProducto, error) {
    OrdenProductoCreado, err := s.r.CrearOrdenProducto(op)
    if err != nil {
        return domain.OrdenProducto{}, err
    }
    return OrdenProductoCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE ORDENPRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
	OrdenProductoCreado, err := s.r.BuscaOrdenProducto(id)
	if err != nil {
		return domain.OrdenProducto{}, err
	}
	return OrdenProductoCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA  UN  ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateOrdenProducto(id int, u domain.OrdenProducto) (domain.OrdenProducto, error) {
	// Llama directamente a la actualización en el repositorio
	return s.r.UpdateOrdenProducto(id, u)
}