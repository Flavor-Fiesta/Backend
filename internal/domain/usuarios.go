package domain

type Usuarios struct {
	ID        int    `json:"id"`
	Nombre string   `json:"nombre" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Telefono 	  string `json:"telefono" binding:"required"`
	Password	  string `json:"password" binding:"required"`
	Id_rol     	  int    `json:"id_rol" binding:"required"`
}
