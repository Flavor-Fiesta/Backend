package domain

type Carrito struct {
    ID  int     `json:"id"`
    Id_usuario  *int    `json:"usuario_id" db:"usuario_id"`
    Id_producto *int    `json:"producto_id" db:"producto_id"`
    Total      float64 `json:"total" db:"total"`
}