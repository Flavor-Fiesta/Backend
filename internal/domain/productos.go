package domain


type Producto struct {
	ID                 int             `json:"id"`
	Nombre             string          `json:"nombre"`
	Descripcion        string          `json:"descripcion"`
	Precio             float64         `json:"precio"`
	Stock              int             `json:"stock"`
	Ranking            float64         `json:"ranking"`
	Id_categoria	   int			   `json:"id_categoria"`
	Categoria		   string	       `json:"categoria"`
	Imagenes           []Imagen        `json:"imagenes"`
}
