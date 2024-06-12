package auth

import (
	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

// Service define los métodos que debe implementar el servicio de autenticación
type Service interface {
	Login(email, password string) (domain.Usuarios, error)
	Authenticate(credentials Credentials) (string, error)
}

type service struct {
	repo Repository
}

// NewService crea un nuevo servicio de autenticación
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Login(email, password string) (domain.Usuarios, error) {
	return s.repo.Login(email, password)
}

func (s *service) Authenticate(credentials Credentials) (string, error) {
	return s.repo.Authenticate(credentials)
}