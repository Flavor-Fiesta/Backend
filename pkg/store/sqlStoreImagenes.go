package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"log"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStoreImagen(db *sql.DB) StoreInterfaceImagenes {
	return &sqlStore{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UNA NUEVA IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStore) CrearImagen(imagen domain.Imagen) error {
	query := "INSERT INTO imagenes (id_producto, titulo, url) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(imagen.Id_producto, imagen.Titulo, imagen.Url)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}

	return nil
}
func (s *sqlStore) ExisteProductoParaImagen(id int) (bool, error) {
    // Consulta SQL para buscar un producto por su ID
    query := "SELECT id FROM productos WHERE id = ?"

    // Ejecutar la consulta SQL y escanear el resultado en una variable id
    var count int
    err := s.db.QueryRow(query, id).Scan(&count)
    if err != nil {
        // Si se produce un error, verificamos si se trata de un error de "ninguna fila encontrada"
        if err == sql.ErrNoRows {
            // No se encontró ninguna fila, por lo que el producto no existe
            return false, nil
			fmt.Print(" REVISANDO EN SQL STORE:  ", query, count )
        }
        // Otro tipo de error, devolver el error
        return false, err
    }

    // Si se encontró un producto con el ID dado, devolver true
    return true, nil
}



// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR IMAGEN POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStore) BuscarImagen(id int) (domain.Imagen, error) {
	var imagen domain.Imagen
	query := "SELECT id, titulo, url FROM imagenes WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&imagen.ID, &imagen.Titulo, &imagen.Url)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Imagen{}, errors.New("imagen not found")
		}
		return domain.Imagen{}, err
	}

	return imagen, nil
}

func (s *sqlStore) BuscarProductoPorID(id int) (domain.Producto, error) {
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

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> ACTUALIZA UNA IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStore) UpdateImagen(id int, p domain.Imagen) error {
	// Preparar la consulta SQL para actualizar la imagen
	query := "UPDATE imagenes SET producto_id = ?, titulo = ?, url = ? WHERE id = ?;"

	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query,p.Id_producto, p.Titulo, p.Url,id)
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
		return fmt.Errorf("Imagen con ID %d no encontrado", id)
	}
	// Si todo fue exitoso, retornar nil
	return nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> PATCH IMAGEN >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (s *sqlStore) PatchImagen(id int, updatedFields map[string]interface{}) error {
    // Comprobar si se proporcionan campos para actualizar
    if len(updatedFields) == 0 {
        return errors.New("no fields provided for patching")
    }

    // Construir la consulta SQL para actualizar los campos
    query := "UPDATE imagenes SET"
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

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ELIMINAR UNA IMAGEN <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStore) DeleteImagen(id int) error {
	query := "DELETE FROM imagenes WHERE id = ?;"
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
func (s *sqlStore) ExistsByIDImagen(id int) bool {
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
//************************************************************************************************************************************//



