package users

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"backend-PAAPO/internal/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db Repository
}

func NewService(r Repository) *Service {
	return &Service{db: r}
}

func (s *Service) Login(user dto.LoginDTO) (*User, error) {
	fmt.Println("Login: ", user)
	if user.Email == "" || user.Password == "" {
		fmt.Println("Deu erro aqui! 1")
		return nil, errors.New("email e senha são obrigatórios")
	}

	// Regex de email
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	fmt.Println("Chegou aqui!")
	if !re.MatchString(user.Email) {
		fmt.Println("Deu erro aqui! 2")
		return nil, errors.New("email inválido")
	}

	// Busca o usuário
	u, err := s.db.GetByEmail(user.Email)
	fmt.Println("chegou aqui! 2")
	if err != nil {
		fmt.Println("Deu erro aqui! 3")
		return nil, errors.New("usuário não encontrado")
	}
	if u == nil {
		fmt.Println("Deu erro aqui! 4")
		return nil, errors.New("usuário não encontrado")
	}

	// Verifica a senha usando CheckPasswordHash
	fmt.Println(u.Password)
	fmt.Println(user.Password)

	match, err := ComparePasswordAndHash(user.Password, u.Password)
	if err != nil || !match {
		fmt.Println("Deu erro aqui! 5")
		return nil, errors.New("senha incorreta")
	}

	fmt.Println("chegou aqui! 5")
	return u, nil
}

// Um método alternativo que não usa bcrypt.CompareHashAndPassword
func ComparePasswordAndHash(password, hash string) (bool, error) {
	fmt.Println("Comparando senha e hash")
	fmt.Print("Senha: ", password)
	fmt.Print("Hash: ", hash)
	if len(hash) < 60 {
		return false, errors.New("hash inválido")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *Service) CreateSession(userID uuid.UUID) (string, error) {
	sessionID := uuid.New().String()
	return sessionID, nil
}

func (s *Service) Register(user dto.RegisterDTO) (*User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email e senha são obrigatórios")
	}

	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user.Email) {
		return nil, errors.New("email inválido")
	}

	// Verifica se já existe usuário com o email
	existing, err := s.db.GetByEmail(user.Email)
	if err == nil && existing != nil {
		return nil, errors.New("usuário já registrado")
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("erro ao gerar hash da senha")
	}

	// Criação do usuário
	newUser := &User{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      string(hashedPassword),
		IsActive:      true,
		IsConfirmed:   false,
		LastUsedEmail: user.Email,
	}

	// Salvar no repositório
	newUser, err = s.db.Create(newUser)
	if err != nil {
		return nil, errors.New("erro ao salvar usuário")
	}

	return newUser, nil
}
