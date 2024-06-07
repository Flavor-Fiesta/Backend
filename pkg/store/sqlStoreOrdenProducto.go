package store

import (
	"database/sql"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type sqlStoreOrdenProductos struct {
	db *sql.DB
}

func NewSqlStoreOrdenProducto(db *sql.DB) StoreInterfaceOrdenProducto {
	return &sqlStoreOrdenProductos{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ORDENPRODUCTO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreOrdenProductos) CrearOrdenProducto(op domain.OrdenProducto) error {
	query := "INSERT INTO OrdenProducto (id_orden, id_producto, cantidad, total) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(op.ID_Orden, op.ID_Producto, op.Cantidad, op.Total)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

// BuscarOrdenProductoPorID busca una relación entre orden y producto por su ID
func (s *sqlStoreOrdenProductos) BuscaOrdenProducto(id int) (domain.OrdenProducto, error) {
	var op domain.OrdenProducto
	query := "SELECT id, id_orden, id_producto, cantidad, total FROM OrdenProducto WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&op.ID, &op.ID_Orden, &op.ID_Producto, &op.Cantidad, &op.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.OrdenProducto{}, fmt.Errorf("order product not found")
		}
		return domain.OrdenProducto{}, fmt.Errorf("error executing query: %w", err)
	}

	return op, nil
}

func (s *sqlStoreOrdenProductos) UpdateOrdenProducto(id int, op domain.OrdenProducto) error {
	query := "UPDATE OrdenProducto SET id_orden=?, id_producto=?, cantidad=?, total=? WHERE id=?;"


	// Ejecutar la consulta SQL
	result, err := s.db.Exec(query, op.ID_Orden, op.ID_Producto, op.Cantidad, op.Total, id)
	if err != nil {
		return err 
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("OrdenProducto con ID %d no encontrado", id)
	}
	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> VERIFICA SI EXISTE ORDENPRODUCTO CON ESE ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreOrdenProductos) ExistsByID(id int) bool {
	// Preparar la consulta SQL para verificar si un odontólogo con el ID dado existe
	query := "SELECT COUNT(*) FROM OrdenProducto WHERE id = ?"
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
