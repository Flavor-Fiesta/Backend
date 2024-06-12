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
    Login(email, password string) (domain.Usuarios, error) // Agregar este método

}


//ACA USO LA INTERFAZ DE USUARIO PARA TRAER LA FUNCION BUSCARUSUARIOPORMAILYPASWORD
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
    // Autenticar las credenciales del usuario
    user, err := r.storage.BuscarUsuarioPorEmailYPassword2(credentials.Email, credentials.Password)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    // Generar el token JWT
    token, err := GenerateToken(user.Email)
    if err != nil {
        return "", err
    }

    // Devolver el token en la respuesta
    return token, nil
}



var jwtKey = []byte("your-secret-key") // Cambia esto por tu propia clave secreta

// Credentials representa las credenciales de inicio de sesión
type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Claims representa la estructura del token JWT
type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

// GenerateToken genera un token JWT para un usuario autenticado
func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // El token expira en 24 horas

    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}