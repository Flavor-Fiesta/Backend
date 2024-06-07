package categorias

import (
	"fmt"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type Service interface {

	CrearCategoria(p domain.Categoria) (domain.Categoria, error)
	BuscarCategoria(id int) (domain.Categoria, error)
	BuscarTodosLasCategorias() ([]domain.Categoria, error)

	DeleteCategoria(id int) error	
//	ExistsByIDCategoria(id int) (bool, error)

//UpdateCategoria(id int, p domain.Categoria) (domain.Categoria, error)
    Update(id int, p domain.Categoria) (domain.Categoria, error)

}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// CrearProducto crea un nueva categoria utilizando el repositorio y devuelve la categoria creado
func (s *service) CrearCategoria(p domain.Categoria) (domain.Categoria, error) {
    // Crear el producto utilizando el repositorio
    productoCreado, err := s.r.CrearCategoria(p)
    if err != nil {
        return domain.Categoria{}, err
    }
    return productoCreado, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE CATEGORIA POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// BuscarCategoria busca una categoria por su ID y devuelve también los datos de la categoria asociada
func (s *service) BuscarCategoria(id int) (domain.Categoria, error) {
    categoria, err := s.r.BuscarCategoria(id)
    if err != nil {
        return domain.Categoria{}, err
    }
    return categoria, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> OBTIENE TODAS LAS CATEGORIAS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) BuscarTodosLasCategorias() ([]domain.Categoria, error) {
    categoria, err := s.r.BuscarTodosLasCategorias()
    if err != nil {
        return nil, fmt.Errorf("error buscando todas las categorias: %w", err)
    }
    return categoria, nil
}


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UNA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *service) Update(id int, u domain.Categoria) (domain.Categoria, error) {
	p, err := s.r.BuscarCategoria(id)
	if err != nil {
		return domain.Categoria{}, err
	}
// Actualizar los campos de la categoría existente con los valores de la nueva categoría `u`
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Descripcion != "" {
		p.Descripcion = u.Descripcion
	}

	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Categoria{}, err
	}
	
return p, nil

}


//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) Patch(id int, updatedFields map[string]interface{}) (domain.Categoria, error) {
 // Obtener el odontólogo por su ID
 usuario, err := s.r.BuscarCategoria(id)
 if err != nil {
     return domain.Categoria{}, err
 }

 // Actualizar los campos proporcionados en updatedFields
 for field, value := range updatedFields {
     switch field {
     case "Nombre":
         if nombre, ok := value.(string); ok {
             usuario.Nombre = nombre
         }
     case "Descripcion":
         if descripcion, ok := value.(string); ok {
             usuario.Descripcion = descripcion
         }
     // Puedes añadir más campos aquí según sea necesario
     default:
         return domain.Categoria{}, fmt.Errorf("campo desconocido: %s", field)
     }
 }

 // Actualizar el odontólogo en el repositorio
 updatedUsuario, err := s.r.Update(id, usuario)
 if err != nil {
     return domain.Categoria{}, err
 }

 return updatedUsuario, nil
}
    

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ELIMINAR CATEGORIA >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *service) DeleteCategoria(id int) error {
    err := s.r.DeleteCategoria(id)
    if err != nil {
        return err
    }
    return nil
}