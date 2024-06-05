package domain

type CarritosOrden struct {
	ID             int     `json:"id" `
	Id_carrito     int     `json:"id_carrito" `
	Id_orden       int     `json:"id_orden" `
	Cantidad       int     `json:"cantidad" `
	PrecioUnitario float64 `json:"precioUnitario" `
}
