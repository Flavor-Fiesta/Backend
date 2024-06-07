package store

import (
	"database/sql"
//	"errors"
	"fmt"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
//	"log"
)

type sqlStoreRoles struct {
	db *sql.DB
}

func NewSqlStoreRoles(db *sql.DB) StoreInterfaceRoles {
	return &sqlStoreRoles{
		db: db,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CREAR UN NUEVO ROL <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
func (s *sqlStoreRoles) CrearRol(rol domain.Rol) error {
	query := "INSERT INTO roles (nombre) VALUES (?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(rol.Nombre)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	return nil
}
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BUSCAR TODOS LOS ROLES <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

func (s *sqlStoreRoles) BuscarTodosLosRoles() ([]domain.Rol, error) {
	var roles []domain.Rol
	query := "SELECT id, nombre FROM roles"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying usuarios: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rol domain.Rol
		if err := rows.Scan(&rol.ID, &rol.Nombre); err != nil {
			return nil, fmt.Errorf("error scanning usuario: %w", err)
		}
		roles = append(roles, rol)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return roles, nil
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> CAMBIAR DE ROL >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (s *sqlStoreRoles) CambiarRol(usuarioID int, nuevoRol string) error {
    query := "UPDATE usuarios SET rol = ? WHERE id = ?"
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error preparing query: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(nuevoRol, usuarioID)
    if err != nil {
        return fmt.Errorf("error executing query: %w", err)
    }

    return nil
}