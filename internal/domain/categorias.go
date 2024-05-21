package domain

type Categoria struct {
	ID        int    `json:"id"`
	Nombre int   `json:"nombre" binding:"required"`
	Descripcion  string `json:"descripcion" binding:"required"`

}