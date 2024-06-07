package domain

type Imagen struct {
	ID           int    `json:"id"`
	Id_producto  int   `json:"id_producto"`
	Titulo       string `json:"titulo" binding:"required"`
	Url 	     string `json:"url" binding:"required"`
}
