package carritos

import (
	//"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {
	CrearCarrito(p domain.Carrito) (domain.Carrito, error)
	DeleteCarrito(id int) error
	UpdateCarrito(id int, p domain.Carrito) (domain.Carrito, error)
}

type service struct {
	r Repository
}

// NewService creates a new service for carritos
func NewService(r Repository) Service {
	return &service{r}
}

// CrearCarrito creates a new carrito
func (s *service) CrearCarrito(p domain.Carrito) (domain.Carrito, error) {
	p, err := s.r.CrearCarrito(p)
	if err != nil {
		return domain.Carrito{}, err
	}
	return p, nil
}



// UpdateCarrito updates an existing carrito
func (s *service) UpdateCarrito(id int, u domain.Carrito) (domain.Carrito, error) {
	return s.r.UpdateCarrito(id, u)
}

// Patch updates specific fields of a carrito

/*
func (s *service) Patch(id int, updatedFields map[string]interface{}) (domain.Carrito, error) {
	carrito, err := s.r.BuscarCarrito(id)
	if err != nil {
		return domain.Carrito{}, err
	}

	for field, value := range updatedFields {
		switch field {
		case "UsuarioID":
			if usuarioID, ok := value.(int); ok {
				carrito.UsuarioID = usuarioID
			}
		case "ProductoID":
			if productoID, ok := value.(int); ok {
				carrito.ProductoID = productoID
			}
		case "Total":
			if total, ok := value.(float64); ok {
				carrito.Total = total
			}
		default:
			return domain.Carrito{}, fmt.Errorf("campo desconocido: %s", field)
		}
	}

	updatedCarrito, err := s.r.UpdateCarrito(id, carrito)
	if err != nil {
		return domain.Carrito{}, err
	}

	return updatedCarrito, nil
}
*/


// DeleteCarrito deletes a carrito by ID
func (s *service) DeleteCarrito(id int) error {
	err := s.r.DeleteCarrito(id)
	if err != nil {
		return err
	}
	return nil
}
