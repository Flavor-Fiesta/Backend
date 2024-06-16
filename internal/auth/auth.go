package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jfcheca/FlavorFiesta/internal/domain"
	"github.com/jfcheca/FlavorFiesta/pkg/store"
)

// Repository define la interfaz que debe implementar el repositorio de autenticación
type Repository interface {
	Authenticate(credentials Credentials) (string, error)
	Login(email, password string) (domain.Usuarios, error)
	ForgotPassword(email string) (string, error)
}

type repository struct {
	storage store.StoreInterfaceUsuarios
}

// NewRepository crea un nuevo repositorio de autenticación
func NewRepository(storage store.StoreInterfaceUsuarios) Repository {
	return &repository{storage}
}

func (r *repository) Login(email, password string) (domain.Usuarios, error) {
	return r.storage.BuscarUsuarioPorEmailYPassword2(email, password)
}

func (r *repository) Authenticate(credentials Credentials) (string, error) {
	user, err := r.storage.BuscarUsuarioPorEmailYPassword2(credentials.Email, credentials.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *repository) ForgotPassword(email string) (string, error) {
	user, err := r.storage.ExisteEmail2(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

var jwtKey = []byte("your-secret-key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
