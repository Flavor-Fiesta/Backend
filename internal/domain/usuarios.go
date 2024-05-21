package domain

type Usuarios struct {
	ID        int    `json:"id"`
	Nombre int   `json:"nombre" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Telefono 	  string `json:"telefono" binding:"required"`
}
