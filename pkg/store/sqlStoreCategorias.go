package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"log"
)

type sqlStoreCategorias struct {
	db *sql.DB
}

func NewSqlStoreCategorias(db *sql.DB) StoreInterfaceCategorias {
	return &sqlStoreCategorias{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UNA NUEVA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) CrearCategoria(categoria domain.Categoria) error {
	query := "INSERT INTO categorias (nombre, descripcion) VALUES (?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(categoria.Nombre, categoria.Descripcion)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}

	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR CATEGORIA POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) BuscarCategoria(id int) (domain.Categoria, error) {
	var categoria domain.Categoria
	query := "SELECT id, nombre, descripcion FROM categorias WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&categoria.ID, &categoria.Nombre, &categoria.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Categoria{}, errors.New("categoria not found")
		}
		return domain.Categoria{}, err
	}

	// Obtener los productos de la categoría
	productosQuery := "SELECT id, id_categoria, nombre, descripcion, precio, stock, ranking FROM productos WHERE id_categoria = ?"
	rows, err := s.db.Query(productosQuery, categoria.ID)
	if err != nil {
		return domain.Categoria{}, err
	}
	defer rows.Close()

	var productos []domain.Producto
	for rows.Next() {
		var producto domain.Producto
		if err := rows.Scan(&producto.ID, &producto.Id_categoria, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.Ranking); err != nil {
			return domain.Categoria{}, err
		}

		// Obtener las imágenes del producto
		imagenesQuery := "SELECT id, id_producto, titulo, url FROM imagenes WHERE id_producto = ?"
		imgRows, err := s.db.Query(imagenesQuery, producto.ID)
		if err != nil {
			return domain.Categoria{}, err
		}
		defer imgRows.Close()

		for imgRows.Next() {
			var imagen domain.Imagen
			if err := imgRows.Scan(&imagen.ID, &imagen.Id_producto, &imagen.Titulo, &imagen.Url); err != nil {
				return domain.Categoria{}, err
			}
			producto.Imagenes = append(producto.Imagenes, imagen)
		}

		productos = append(productos, producto)
	}

	categoria.Productos = productos
	return categoria, nil
}
//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODAS LAS CATEGORIAS  >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *sqlStoreCategorias) BuscarTodosLasCategorias() ([]domain.Categoria, error) {
	var categorias []domain.Categoria
	query := "SELECT id, nombre, descripcion FROM categorias"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying categorias: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var categoria domain.Categoria
		if err := rows.Scan(&categoria.ID, &categoria.Nombre, &categoria.Descripcion); err != nil {
			return nil, fmt.Errorf("error scanning categoria: %w", err)
		}

		// Obtener los productos de la categoría
		productosQuery := "SELECT id, id_categoria, nombre, descripcion, precio, stock, ranking FROM productos WHERE id_categoria = ?"
		productRows, err := s.db.Query(productosQuery, categoria.ID)
		if err != nil {
			return nil, fmt.Errorf("error querying productos: %w", err)
		}
		defer productRows.Close()

		var productos []domain.Producto
		for productRows.Next() {
			var producto domain.Producto
			if err := productRows.Scan(&producto.ID, &producto.Id_categoria, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.Ranking); err != nil {
				return nil, fmt.Errorf("error scanning producto: %w", err)
			}

			// Obtener las imágenes del producto
			imagenesQuery := "SELECT id, id_producto, titulo, url FROM imagenes WHERE id_producto = ?"
			imgRows, err := s.db.Query(imagenesQuery, producto.ID)
			if err != nil {
				return nil, fmt.Errorf("error querying imagenes: %w", err)
			}
			defer imgRows.Close()

			for imgRows.Next() {
				var imagen domain.Imagen
				if err := imgRows.Scan(&imagen.ID, &imagen.Id_producto, &imagen.Titulo, &imagen.Url); err != nil {
					return nil, fmt.Errorf("error scanning imagen: %w", err)
				}
				producto.Imagenes = append(producto.Imagenes, imagen)
			}

			productos = append(productos, producto)
		}

		categoria.Productos = productos
		categorias = append(categorias, categoria)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return categorias, nil
}



/*func (s *sqlStoreUsuarios) BuscarProductoPorID(id int) (domain.Producto, error) {
    // Preparar la consulta SQL para buscar un producto por su ID
    query := "SELECT id, nombre, codigo, categoria, fecha_alta, fecha_vencimiento FROM productos WHERE id = ?"
    
    // Ejecutar la consulta SQL y obtener el resultado
    var producto domain.Producto
    err := s.db.QueryRow(query, id).Scan(&producto.ID, &producto.Nombre, &producto.Codigo, &producto.Categoria, &producto.FechaDeAlta, &producto.FechaDeVencimiento)
    if err != nil {
        // Manejar el error, por ejemplo, devolver un error específico si no se encuentra el producto
        if err == sql.ErrNoRows {
            return domain.Producto{}, fmt.Errorf("producto con ID %d no encontrado", id)
        }
        return domain.Producto{}, err
    }

    // Devolver el producto encontrado
    return producto, nil
}
*/
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UNA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) Update(p domain.Categoria) error {
	// Preparar la consulta SQL para actualizar la imagen
	query := "UPDATE categorias SET nombre = ?, descripcion = ? WHERE id = ?;"

	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query,p.Nombre, p.Descripcion, p.ID)
	if err != nil {
		return err // Devolver el error si ocurre alguno al ejecutar la consulta
	}
	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Si no se actualizó ningún registro, significa que la imagen con el ID dado no existe
	if rowsAffected == 0 {
		return fmt.Errorf("Categoria con ID %d no encontrada", p.ID)
	}
	// Si todo fue exitoso, retornar nil
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) PatchCategoria(id int, updatedFields map[string]interface{}) error {
    // Comprobar si se proporcionan campos para actualizar
    if len(updatedFields) == 0 {
        return errors.New("no fields provided for patching")
    }

    // Construir la consulta SQL para actualizar los campos
    query := "UPDATE categorias SET"
    values := make([]interface{}, 0)
    index := 0
    for field, value := range updatedFields {
        query += fmt.Sprintf(" %s = ?", field)
        values = append(values, value)
        index++
        if index < len(updatedFields) {
            query += ","
        }
    }
    query += " WHERE id = ?"
    values = append(values, id)

    // Preparar y ejecutar la consulta SQL
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ELIMINAR UNA CATEGORIA <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) DeleteCategoria(id int) error {
	query := "DELETE FROM categorias WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> VERIFICA SI EXISTE CATEGORIA CON ESE ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreCategorias) ExistsByIDCategoria(id int) (bool, error) {
    // Consulta SQL para buscar una categoría por su ID
    query := "SELECT id FROM categorias WHERE id = ?"

    // Ejecutar la consulta SQL y escanear el resultado en una variable id
    var count int
    err := s.db.QueryRow(query, id).Scan(&count)
    if err != nil {
        // Si se produce un error, verificamos si se trata de un error de "ninguna fila encontrada"
        if err == sql.ErrNoRows {
            // No se encontró ninguna fila, por lo que la categoría no existe
            return false, nil
        }
        // Otro tipo de error, devolver el error
        return false, err
    }

    // Si se encontró una categoría con el ID dado, devolver true
    return true, nil
}