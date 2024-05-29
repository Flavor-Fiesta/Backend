package domain

type Carrito struct {
    CarritoID  int     `json:"carrito_id" db:"carrito_id"`
    UsuarioID  *int    `json:"usuario_id" db:"usuario_id"`
    ProductoID *int    `json:"producto_id" db:"producto_id"`
    Total      float64 `json:"total" db:"total"`
}