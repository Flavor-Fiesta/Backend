package domain

type Categoria struct {
	ID           int    `json:"id"`
	Nombre 	     string   `json:"nombre" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`
	Productos   []Producto `json:"productos"`

}