package store

import (
	"database/sql"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"log"
)

type sqlStoreCarritos struct {
	db *sql.DB
}


func NewSqlStoreCarrito(db *sql.DB) StoreInterfaceCarritos {
	return &sqlStoreCarritos{
		db: db,
	}
}

func (s *sqlStoreCarritos) CrearCarrito(carrito domain.Carrito) error {
	query := "INSERT INTO carrito (usuario_id, producto_id, Total) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(carrito.Id_usuario, carrito.Id_producto, carrito.Total)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}

	return nil
}

func (s *sqlStoreCarritos) GetCarritoByID(id int) (domain.Carrito, error) {
    var carrito domain.Carrito
    query := "SELECT id, usuario_id, producto_id, total FROM carrito WHERE id = ?;"
    err := s.db.QueryRow(query, id).Scan(&carrito.ID, &carrito.Id_usuario, &carrito.Id_producto, &carrito.Total)
    if err != nil {
        return domain.Carrito{}, err
    }
    return carrito, nil
}


// UpdateCarrito updates an existing carrito
func (s *sqlStoreCarritos) UpdateCarrito(id int, carrito domain.Carrito) error {
	query := "UPDATE carrito SET usuario_id = ?, producto_id = ?, Total = ? WHERE id = ?;"

	result, err := s.db.Exec(query, carrito.Id_usuario, carrito.Id_producto, carrito.Total, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Carrito with ID %d not found", id)
	}

	return nil
}

// DeleteCarrito deletes a carrito by ID
func (s *sqlStoreCarritos) DeleteCarrito(id int) error {
	query := "DELETE FROM carrito WHERE id = ?;"
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

func (s *sqlStoreCarritos) ExistsByIDCarrito(id int) bool {
	query := "SELECT COUNT(*) FROM carrito WHERE id = ?"
	var count int
	err := s.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
