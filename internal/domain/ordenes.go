package domain

type Orden struct {
    ID    int       `json:"id"`
    FechaOrden string `json:"fechaOrden" db:"fechaOrden"`
    Total      float64   `json:"total" db:"total"`
    Estado     string    `json:"estado" db:"estado"`
}