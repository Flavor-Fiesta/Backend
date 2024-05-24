package domain


type Producto struct {
	ID                 int             `json:"id"`
	Nombre             string          `json:"nombre"`
	Codigo             string          `json:"codigo"`
	Categoria          string          `json:"categoria"`
	FechaDeAlta        string          `json:"fecha_de_alta"`
	FechaDeVencimiento string          `json:"fecha_de_vencimiento"`
	Imagenes           []Imagen `json:"imagenes"`
	//	Imagenes []Imagen `json:"imagenes"`
	//	Imagen []Imagen `json:"imagen"`
	//	ImagenUrl string `json:"imagen_url"`
}
