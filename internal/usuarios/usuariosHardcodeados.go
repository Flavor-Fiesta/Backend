package usuarios

import (
	"golang.org/x/crypto/bcrypt"
)


type HardcodedUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RolesID  int    `json:"roles_id"`
}

func GetHardcodedUser() HardcodedUser {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	return HardcodedUser{
		ID:       1,
		Name:     "Javier",
		LastName: "Checa",
		Email:    "javito@gmail.com",
		Password: string(hashedPassword),
		RolesID:  1,
	}
}