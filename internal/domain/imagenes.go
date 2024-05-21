package domain

type Imagen struct {
	ID        int    `json:"id"`
	ProductoID int   `json:"producto_id"`
	Titulo  string `json:"titulo" binding:"required"`
	Url 	  string `json:"url" binding:"required"`
}
