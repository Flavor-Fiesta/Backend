package productos

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {

	BuscarProducto(id int) (domain.Producto, error)
    BuscarTodosLosProductos() ([]domain.Producto, error)
	CrearProducto(p domain.Producto) (domain.Producto, error)
	UpdateProducto(id int, p domain.Producto) (domain.Producto, error)
	DeleteProducto(id int) error
    ObtenerNombreCategoria(id int) (string, error) // Nuevo método agregado
    

}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// CrearProducto crea un nuevo producto utilizando el repositorio y devuelve el producto creado
func (s *service) CrearProducto(p domain.Producto) (domain.Producto, error) {
    // Crear el producto utilizando el repositorio
    productoCreado, err := s.r.CrearProducto(p)
    if err != nil {
        return domain.Producto{}, err
    }
    return productoCreado, nil
}

func (s *service) ObtenerNombreCategoria(idCategoria int) (string, error) {
    return s.r.ObtenerNombreCategoria(idCategoria)
}


/*func (s *service) CrearProducto(p domain.Producto) (domain.Producto, error) {
    // Crear el producto utilizando el repositorio
    productoCreado, err := s.r.CrearProducto(p)
    if err != nil {
        return domain.Producto{}, err
    }
    return productoCreado, nil
}
*/

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE PRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// BuscarProducto busca un producto por su ID y devuelve también los datos de la imagen asociada
func (s *service) BuscarProducto(id int) (domain.Producto, error) {
	producto, err := s.r.BuscarProducto(id)
	if err != nil {
		return domain.Producto{}, err
	}
	return producto, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODOS LOS PRODUCTOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodosLosProductos() ([]domain.Producto, error) {
    productos, err := s.r.BuscarTodosLosProductos()
    if err != nil {
        return nil, fmt.Errorf("error buscando todos los productos: %w", err)
    }
    return productos, nil
}


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA  UN  PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) UpdateProducto(id int, u domain.Producto) (domain.Producto, error) {
	// Llama directamente a la actualización en el repositorio
	return s.r.UpdateProducto(id, u)
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) Patch(id int, updatedFields map[string]interface{}) (domain.Producto, error) {
    // Obtener el paciente por su ID
    producto, err := s.r.BuscarProducto(id)
    if err != nil {
        return domain.Producto{}, err
    }

    // Actualizar los campos proporcionados en updatedFields
    for field, value := range updatedFields {
        switch field {
        case "Nombre":
            if nombre, ok := value.(string); ok {
                producto.Nombre = nombre
            }
        case "Descripcion":
            if descripcion, ok := value.(string); ok {
                producto.Descripcion = descripcion
            }
 /*       case "Categoria":
            if categoria, ok := value.(string); ok {
                producto.Categoria = categoria
            }*/
        case "Precio":
            if precio, ok := value.(float64); ok {
                producto.Precio = precio
            }
        case "Stock":
            if stock, ok := value.(int); ok {
                producto.Stock = stock
            }
        case "Ranking":
            if ranking, ok := value.(float64); ok {
                producto.Ranking = ranking
            }
        // Puedes añadir más campos aquí según sea necesario
        default:
            return domain.Producto{}, fmt.Errorf("campo desconocido: %s", field)
        }
    }

    // Actualizar el producto en el repositorio
    updatedProducto, err := s.r.UpdateProducto(id, producto)
    if err != nil {
        return domain.Producto{}, err
    }

    return updatedProducto, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) DeleteProducto(id int) error {
    err := s.r.DeleteProducto(id)
    if err != nil {
        return err
    }
    return nil
}