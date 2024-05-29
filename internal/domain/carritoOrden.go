package domain

type CarritosOrden struct {
    OrdenID       int     `json:"orden_id" db:"orden_id"`
    CarritoID     int     `json:"carrito_id" db:"carrito_id"`
    Cantidad      int     `json:"cantidad" db:"cantidad"`
    PrecioUnitario float64 `json:"precio_unitario" db:"precio_unitario"`
}