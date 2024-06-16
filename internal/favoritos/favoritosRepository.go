package favoritos

import (
    "fmt"
    "log"

    "github.com/jfcheca/FlavorFiesta/internal/domain"
    "github.com/jfcheca/FlavorFiesta/pkg/store"
)

type Repository interface {
    
    AgregarFavorito(favorito domain.Favoritos) (domain.Favoritos, error)
}

type repository struct {
    storage store.StoreInterfaceFavoritos
}

func NewRepository(storage store.StoreInterfaceFavoritos) Repository {
    return &repository{storage}
}



func (r *repository) AgregarFavorito(p domain.Favoritos) (domain.Favoritos, error) {
    err := r.storage.AgregarFavorito(p)
    if err != nil {
        log.Printf("Error al agregar el favorito %v: %v\n", p, err)
        return domain.Favoritos{}, fmt.Errorf("error creando favorito: %w", err)
    }
    return p, nil
}