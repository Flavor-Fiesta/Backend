package domain

type Orden struct {
    OrdenID    int       `json:"orden_id" db:"orden_id"`
    FechaOrden string `json:"fecha_orden" db:"fecha_orden"`
    Total      float64   `json:"total" db:"total"`
    Estado     string    `json:"estado" db:"estado"`
}