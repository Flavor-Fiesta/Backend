package store

import (
    "database/sql"
    "fmt"
    "github.com/jfcheca/FlavorFiesta/internal/domain"
)

type sqlStoreFavoritos struct {
    db *sql.DB
}

func NewSqlStoreFavoritos(db *sql.DB) StoreInterfaceFavoritos {
    return &sqlStoreFavoritos{
        db: db,
    }
}

func (s *sqlStoreFavoritos) AgregarFavorito(favorito domain.Favoritos) error {
    query := "INSERT INTO flavorfiesta.favoritos (id_usuario, id_producto) VALUES (?, ?);"
    stmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error preparing query: %w", err)
    }
    defer stmt.Close()

    res, err := stmt.Exec(favorito.Id_usuario, favorito.Id_producto)
    if err != nil {
        return fmt.Errorf("error executing query: %w", err)
    }

    _, err = res.RowsAffected()
    if err != nil {
        return fmt.Errorf("error fetching rows affected: %w", err)
    }

    return nil
}