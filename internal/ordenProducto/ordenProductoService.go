package ordenProductos

import (
	"fmt"
	"log"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	CrearOrdenProducto(p domain.OrdenProducto) (domain.OrdenProducto, error)
	BuscaOrdenProducto(id int) (domain.OrdenProducto, error)
	UpdateOrdenProducto(id int, p domain.OrdenProducto) (domain.OrdenProducto, error)
	BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error)
	BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error)
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
    ordenProductoCreado, err := s.r.CrearOrdenProducto(op)
    if err != nil {
        return domain.OrdenProducto{}, err
    }
    return ordenProductoCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE ORDENPRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
	log.Printf("Service: Buscando OrdenProducto con ID: %d", id)
	OrdenProductoCreado, err := s.r.BuscaOrdenProducto(id)
	if err != nil {
		log.Printf("Service: Error al buscar OrdenProducto: %v", err)
		return domain.OrdenProducto{}, err
	}
	return OrdenProductoCreado, nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE ORDENPRODUCTO FILTRADO POR ID DE LA ORDEN<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarOrdenesProductoPorIDOrden(idOrden int) ([]domain.OrdenProducto, error) {
    ordenesProducto, err := s.r.BuscarOrdenesProductoPorIDOrden(idOrden)
    if err != nil {
        return nil, err
    }
    return ordenesProducto, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODAS LAS ORDENES PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodasLasOrdenesProducto() ([]domain.OrdenProducto, error) {
    usuarios, err := s.r.BuscarTodasLasOrdenesProducto()
    if err != nil {
        return nil, fmt.Errorf("error buscando todas las ordenes producto: %w", err)
    }
    return usuarios, nil
}



// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA  UN  ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateOrdenProducto(id int, u domain.OrdenProducto) (domain.OrdenProducto, error) {
	// Llama directamente a la actualización en el repositorio
	return s.r.UpdateOrdenProducto(id, u)
}