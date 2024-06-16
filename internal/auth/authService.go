package auth

import (
	"errors"
	"sync"

	"github.com/jfcheca/FlavorFiesta/internal/domain"
)

// Service define los métodos que debe implementar el servicio de autenticación
type Service interface {
	Login(email, password string) (domain.Usuarios, error)
	Authenticate(credentials Credentials) (string, error)
	ForgotPassword(email string) (string, error)
	ValidateToken(token string) (string, error)
}

type service struct {
	repo   Repository
	tokens map[string]string
	mu     sync.Mutex
}

// NewService crea un nuevo servicio de autenticación
func NewService(repo Repository) Service {
	return &service{
		repo:   repo,
		tokens: make(map[string]string),
	}
}

func (s *service) Login(email, password string) (domain.Usuarios, error) {
	return s.repo.Login(email, password)
}

func (s *service) Authenticate(credentials Credentials) (string, error) {
	return s.repo.Authenticate(credentials)
}

func (s *service) ForgotPassword(email string) (string, error) {
	token, err := s.repo.ForgotPassword(email)
	if err != nil {
		return "", err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token] = email

	return token, nil
}

func (s *service) ValidateToken(token string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	email, exists := s.tokens[token]
	if !exists {
		return "", errors.New("invalid token")
	}

	return email, nil
}
