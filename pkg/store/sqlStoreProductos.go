package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"log"
)

type sqlStoreProductos struct {
	db *sql.DB
}

func NewSqlStoreProductos(db *sql.DB) StoreInterfaceProducto {
	return &sqlStoreProductos{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) CrearProducto(producto domain.Producto) error {
	query := "INSERT INTO productos (id_categoria, nombre, descripcion, precio, stock, ranking) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(producto.Id_categoria, producto.Nombre, producto.Descripcion, producto.Precio, producto.Stock, producto.Ranking)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

func (s *sqlStoreProductos) ObtenerNombreCategoria(idCategoria int) (string, error) {
	var nombreCategoria string
	query := "SELECT nombre FROM categorias WHERE id = ?"
	err := s.db.QueryRow(query, idCategoria).Scan(&nombreCategoria)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("categoria not found")
		}
		return "", fmt.Errorf("error fetching category name: %w", err)
	}
	return nombreCategoria, nil
}
/*func (s *sqlStoreProductos) CrearProducto(producto domain.Producto) error {
    query := "INSERT INTO productos (id_categoria, nombre, descripcion, precio, stock, ranking) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(producto.Id_categoria, producto.Nombre, producto.Descripcion, producto.Precio, producto.Stock, producto.Ranking)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}

	return nil
}
*/
/*    stmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error al preparar la consulta SQL: %w", err)
    }
    defer stmt.Close()

    result, err := stmt.Exec(producto.Id_categoria, producto.Nombre, producto.Descripcion, producto.Precio, producto.Stock, producto.Ranking)
    if err != nil {
        return fmt.Errorf("error al ejecutar la consulta SQL para insertar producto: %w", err)
    }
    
    // Obtener el número de filas afectadas
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al obtener el número de filas afectadas: %w", err)
    }
    if rowsAffected != 1 {
        return fmt.Errorf("se esperaba que se afectara una fila, pero se afectaron %d filas", rowsAffected)
    }

    // Obtener el ID del producto insertado
    productoID, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("error al obtener el ID del producto insertado: %w", err)
    }

    // Asignar el ID del producto
    producto.ID = int(productoID)

    // Consultar el nombre de la categoría
    var categoriaNombre string
    err = s.db.QueryRow("SELECT nombre FROM categorias WHERE id = ?", producto.Id_categoria).Scan(&categoriaNombre)
    if err != nil {
        return fmt.Errorf("error al obtener el nombre de la categoría: %w", err)
    }

    return nil
}*/

    //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR PRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) BuscarProducto(id int) (domain.Producto, error) {
    var producto domain.Producto

    // Consulta principal del producto
    query := `SELECT id, id_categoria, nombre, descripcion, precio, stock, ranking
              FROM productos WHERE id = ?`

    err := s.db.QueryRow(query, id).Scan(
        &producto.ID,
        &producto.Id_categoria,
        &producto.Nombre,
        &producto.Descripcion,
        &producto.Precio,
        &producto.Stock,
        &producto.Ranking,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            return domain.Producto{}, errors.New("producto not found")
        }
        return domain.Producto{}, err
    }

    // Consulta de la categoría del producto
    if producto.Id_categoria != 0 {
        categoriaQuery := `SELECT nombre FROM categorias WHERE id = ?`
        var categoria string
        err = s.db.QueryRow(categoriaQuery, producto.Id_categoria).Scan(&categoria)
        if err != nil {
            return domain.Producto{}, err
        }
        producto.Categoria = categoria // Asignar la categoría al producto
    }

    // Inicializar la lista de imágenes
    producto.Imagenes = []domain.Imagen{}

    // Consulta de las imágenes del producto
    imagenesQuery := `SELECT id, id_producto, titulo, url FROM imagenes WHERE id_producto = ?`
    rows, err := s.db.Query(imagenesQuery, producto.ID)
    if err != nil {
        return domain.Producto{}, err
    }
    defer rows.Close()

    for rows.Next() {
        var imagen domain.Imagen
        if err := rows.Scan(&imagen.ID, &imagen.Id_producto, &imagen.Titulo, &imagen.Url); err != nil {
            return domain.Producto{}, err
        }
        producto.Imagenes = append(producto.Imagenes, imagen)
    }

    return producto, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR TODOS LOS PRODUCTOS<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (s *sqlStoreProductos) BuscarTodosLosProductos() ([]domain.Producto, error) {
    var productos []domain.Producto
    query := `
        SELECT 
            p.id, p.nombre, p.descripcion, c.nombre as categoria, p.precio, p.stock, p.ranking 
        FROM 
            productos p
        LEFT JOIN 
            categorias c ON p.id_categoria = c.id
    `

    rows, err := s.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying productos: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var producto domain.Producto
        var categoriaNombre string // Variable adicional para la categoría

        if err := rows.Scan(&producto.ID, &producto.Nombre, &producto.Descripcion, &categoriaNombre, &producto.Precio, &producto.Stock, &producto.Ranking); err != nil {
            return nil, fmt.Errorf("error scanning producto: %w", err)
        }

        producto.Categoria = categoriaNombre // Asignar la categoría al producto

        // Obtener las imágenes del producto
        imagenesQuery := "SELECT id, id_producto, titulo, url FROM imagenes WHERE id_producto = ?"
        imagenesRows, err := s.db.Query(imagenesQuery, producto.ID)
        if err != nil {
            return nil, err
        }
        defer imagenesRows.Close()

        for imagenesRows.Next() {
            var imagen domain.Imagen
            if err := imagenesRows.Scan(&imagen.ID, &imagen.Id_producto, &imagen.Titulo, &imagen.Url); err != nil {
                return nil, err
            }
            producto.Imagenes = append(producto.Imagenes, imagen)
        }

        productos = append(productos, producto)
    }

    if err := rows.Err(); err != nil { // Asegúrate de que esta condición esté evaluando un booleano
        return nil, fmt.Errorf("error iterating rows: %w", err)
    }

    return productos, nil
}


// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UN PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) UpdateProducto(id int, p domain.Producto) error {
	// Preparar la consulta SQL para actualizar el producto
	query := "UPDATE productos SET nombre = ?, descripcion = ?, categoria = ?, fecha_alta = ?, fecha_vencimiento = ? WHERE id = ?;"

	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query, p.Nombre, p.Descripcion, p.Precio, p.Stock,p.Ranking, id)
	if err != nil {
		return err // Devolver el error si ocurre alguno al ejecutar la consulta
	}
	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Si no se actualizó ningún registro, significa que el odontólogo con el ID dado no existe
	if rowsAffected == 0 {
		return fmt.Errorf("Producto con ID %d no encontrado", id)
	}
	// Si todo fue exitoso, retornar nil
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (s *sqlStoreProductos) Patch(id int, updatedFields map[string]interface{}) (domain.Producto, error) {
    // Comprobar si se proporcionan campos para actualizar
    if len(updatedFields) == 0 {
        return domain.Producto{}, errors.New("no fields provided for patching")
    }

    // Construir la consulta SQL para actualizar los campos
    query := "UPDATE productos SET"
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
        return domain.Producto{}, err
    }
    defer stmt.Close()

    _, err = stmt.Exec(values...)
    if err != nil {
        return domain.Producto{}, err
    }

    // Si la actualización fue exitosa, puedes devolver el producto actualizado
    // Sin embargo, en este caso, como la actualización podría haber afectado a múltiples registros,
    // No tienes una representación directa del producto actualizado.
    // Puedes optar por devolver un producto vacío o uno con el ID correspondiente.
    // Aquí devolveré un producto vacío.
    return domain.Producto{}, nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ELIMINAR UN PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) DeleteProducto(id int) error {
	query := "DELETE FROM productos WHERE id = ?;"
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

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> VERIFICA SI EXISTE PRODUCTO CON ESE ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) ExistsByID(id int) bool {
	// Preparar la consulta SQL para verificar si un odontólogo con el ID dado existe
	query := "SELECT COUNT(*) FROM productos WHERE id = ?"
	// Ejecutar la consulta SQL y obtener el número de filas devueltas
	var count int
	err := s.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		// Manejar el error, por ejemplo, loguearlo o devolver false si ocurre un error
		return false
	}
	// Si el número de filas devueltas es mayor que cero, significa que el odontólogo con el ID dado existe
	return count > 0
}
