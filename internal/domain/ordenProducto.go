package domain

// OrdenProducto representa la relación entre una orden y un producto en la base de datos
type OrdenProducto struct {
	ID         int     `json:"id"`
	ID_Orden    int     `json:"id_orden"`
	ID_Producto int     `json:"id_producto"`
	Producto Producto   `json:"producto"`
	Cantidad   int     `json:"cantidad"`
	Total      float64 `json:"total"`
}	 