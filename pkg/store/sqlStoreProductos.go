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
    query := "INSERT INTO productos (nombre, codigo, categoria, fecha_alta, fecha_vencimiento) VALUES (?, ?, ?, ?, ?);"
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error al preparar la consulta SQL: %w", err)
    }
    defer stmt.Close()

    result, err := stmt.Exec(producto.Nombre, producto.Codigo, producto.Categoria, producto.FechaDeAlta, producto.FechaDeVencimiento)
    if err != nil {
        return fmt.Errorf("error al ejecutar la consulta SQL para insertar producto: %w", err)
    }

    productoID, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("error al obtener el ID del producto insertado: %w", err)
    }

    // Insertar imágenes asociadas al producto
    for _, imagen := range producto.Imagenes {
        query := "INSERT INTO imagenes (producto_id, titulo, url) VALUES (?, ?, ?);"
        stmt, err := s.db.Prepare(query)
        if err != nil {
            return fmt.Errorf("error al preparar la consulta SQL para insertar imagen: %w", err)
        }
        defer stmt.Close()

        _, err = stmt.Exec(productoID, imagen.Titulo, imagen.Url)
        if err != nil {
            return fmt.Errorf("error al ejecutar la consulta SQL para insertar imagen: %w", err)
        }
    }

    return nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR PRODUCTO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) BuscarProducto(id int) (domain.Producto, error) {
	var producto domain.Producto
	query := "SELECT id, nombre, codigo, categoria, fecha_alta, fecha_vencimiento FROM productos WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&producto.ID, &producto.Nombre, &producto.Codigo, &producto.Categoria, &producto.FechaDeAlta, &producto.FechaDeVencimiento)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Producto{}, errors.New("producto not found")
		}
		return domain.Producto{}, err
	}
    return producto, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UN PRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreProductos) UpdateProducto(id int, p domain.Producto) error {
	// Preparar la consulta SQL para actualizar el producto
	query := "UPDATE productos SET nombre = ?, codigo = ?, categoria = ?, fecha_alta = ?, fecha_vencimiento = ? WHERE id = ?;"

	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query, p.Nombre, p.Codigo, p.Categoria, p.FechaDeAlta, p.FechaDeVencimiento, id)
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
		return fmt.Errorf("Odontologo con ID %d no encontrado", id)
	}
	// Si todo fue exitoso, retornar nil
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH PRODUCTO >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (s *sqlStoreProductos) Patch(id int, updatedFields map[string]interface{}) error {
    // Comprobar si se proporcionan campos para actualizar
    if len(updatedFields) == 0 {
        return errors.New("no fields provided for patching")
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
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(values...)
    if err != nil {
        return err
    }

    return nil
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
