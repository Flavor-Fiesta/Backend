package store

import (
	"database/sql"
	"fmt"
	"errors"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

type sqlStoreEstados struct {
	db *sql.DB
}

func NewSqlStoreEstados(db *sql.DB) StoreInterfaceEstados {
	return &sqlStoreEstados{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ESTADO <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreEstados) CrearEstados(estado domain.Estado) error {
	query := "INSERT INTO estados (nombre) VALUES (?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(estado.Nombre)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS ESTADOS <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreEstados) BuscarTodosLosEstados() ([]domain.Estado, error) {
	var estados []domain.Estado
	query := "SELECT id, nombre FROM estados"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying estados: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var estado domain.Estado
		if err := rows.Scan(&estado.ID, &estado.Nombre); err != nil {
			return nil, fmt.Errorf("error scanning estado: %w", err)
		}
		estados = append(estados, estado)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return estados, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  BUSCAR ESTADO POR ID <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreEstados) BuscarEstado(id int) (domain.Estado, error) {
	var estado domain.Estado
	query := "SELECT id, nombre FROM estados WHERE id = ?"

	err := s.db.QueryRow(query, id).Scan(&estado.ID, &estado.Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Estado{}, errors.New("estado not found")
		}
		return domain.Estado{}, err
	}

	return estado, nil
}