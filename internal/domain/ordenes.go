package domain

type Orden struct {
    ID    int       `json:"id"`
    ID_Usuario    int     `json:"id_usuario"`
    ID_Estado int `json:"id_estado"`
    FechaOrden string `json:"fechaOrden" db:"fechaOrden"`
    Total      float64   `json:"total" db:"total"`
}