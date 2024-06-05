package store

import (
	"database/sql"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type sqlStoreCarritoOrdenes struct {
	db *sql.DB
}

func NewSqlStoreCarritoOrdenes(db *sql.DB) StoreInterfaceCarritoOrdenes {
	return &sqlStoreCarritoOrdenes{
		db: db,
	}
}

func (s *sqlStoreCarritoOrdenes) CrearCarritoOrdenes(carrito domain.CarritosOrden) error {
	query := "INSERT INTO carritos_ordenes (id_carrito, id_orden, cantidad, precioUnitario) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(carrito.Id_carrito, carrito.Id_orden, carrito.Cantidad, carrito.PrecioUnitario)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}

	return nil
}

func (s *sqlStoreCarritoOrdenes) GetCarritoOrdenesByID(id int) (domain.CarritosOrden, error) {
	var carrito domain.CarritosOrden
	query := "SELECT id, id_carrito, id_orden, cantidad, precioUnitario FROM carritos_ordenes WHERE id = ?;"
	err := s.db.QueryRow(query, id).Scan(&carrito.ID, &carrito.Id_carrito, &carrito.Id_orden, &carrito.Cantidad, &carrito.PrecioUnitario)
	if err != nil {
		return domain.CarritosOrden{}, err
	}
	return carrito, nil
}

func (s *sqlStoreCarritoOrdenes) UpdateCarritoOrdenes(id int, carrito domain.CarritosOrden) error {
	query := "UPDATE carritos_ordenes SET id_carrito = ?, id_orden = ?, cantidad = ?, precioUnitario = ? WHERE id = ?;"
	result, err := s.db.Exec(query, carrito.Id_carrito, carrito.Id_orden, carrito.Cantidad, carrito.PrecioUnitario, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("CarritoOrden with ID %d not found", id)
	}

	return nil
}

func (s *sqlStoreCarritoOrdenes) DeleteCarritoOrdenes(id int) error {
	query := "DELETE FROM carritos_ordenes WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStoreCarritoOrdenes) ExistsByIDCarritoOrdenes(id int) bool {
	query := "SELECT COUNT(*) FROM carritos_ordenes WHERE id = ?"
	var count int
	err := s.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
