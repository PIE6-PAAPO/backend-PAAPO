package users

import (
	"backend-PAAPO/internal/models"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"time"

	"backend-PAAPO/internal/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
	"golang.org/x/crypto/bcrypt"
)

// Service é a interface que define os métodos do serviço de usuários
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

// getJWTSecret loads the JWT secret at runtime
func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	return []byte(secret)
}

func GenerateJWT(userID uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
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

	// Use atomic group assignment to prevent race conditions
	newUser, err := s.createUserWithAtomicGroupAssignment(user, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// createUserWithAtomicGroupAssignment creates a user with atomic group assignment
// to prevent race conditions during concurrent registrations
func (s *Service) createUserWithAtomicGroupAssignment(user dto.RegisterDTO, hashedPassword string) (*User, error) {
	// Start a database transaction
	tx := s.db.GetDB().Begin()
	if tx.Error != nil {
		return nil, errors.New("erro ao iniciar transação")
	}

	// Defer rollback in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Count users in each group within the transaction
	var testCount, controlCount int64

	if err := tx.Model(&User{}).Where("is_test_group = ?", true).Count(&testCount).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao contar grupo de teste")
	}

	if err := tx.Model(&User{}).Where("is_test_group = ?", false).Count(&controlCount).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao contar grupo de controle")
	}

	// Determine which group to assign (smaller group gets the new user)
	isTestGroup := controlCount >= testCount

	// Create the user within the transaction
	newUser := &User{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      hashedPassword,
		Role:          string(models.Patient),
		IsActive:      true,
		IsConfirmed:   false,
		LastUsedEmail: user.Email,
		IsTestGroup:   isTestGroup,
	}

	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("erro ao salvar usuário")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("erro ao confirmar transação")
	}

	return newUser, nil
}

// gera um código de 6 dígitos para recuperação de senha
func gerarCodigoRecuperacao() (string, error) {
	numeros := "0123456789"
	codigo := make([]byte, 6)

	for i := range codigo {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(numeros))))
		if err != nil {
			return "", err
		}
		codigo[i] = numeros[num.Int64()]
	}
	return string(codigo), nil
}

// SolicitarRecuperacaoSenha inicia o processo de recuperação de senha
func (s *Service) SolicitarRecuperacaoSenha(email string) error {
	// Verifica se o email é válido
	if email == "" {
		return nil
	}

	// Busca o usuário pelo email
	usuario, err := s.db.GetByEmail(email)
	if err != nil {
		// Se não achar o usuário, não retorna erro (por segurança)
		return nil
	}

	// Gera um código de 6 dígitos
	codigo, err := gerarCodigoRecuperacao()
	if err != nil {
		return fmt.Errorf("erro ao gerar código: %v", err)
	}

	// Define que o código expira em 1 hora
	expiraEm := time.Now().Add(1 * time.Hour)
	usuario.RecoveryCode = &codigo
	usuario.RecoveryCodeExpiresAt = &expiraEm

	// Salva o código no banco
	if _, err := s.db.Update(usuario); err != nil {
		return fmt.Errorf("erro ao salvar código: %v", err)
	}

	// Envia o email com o código
	if err := enviarEmailRecuperacao(email, codigo); err != nil {
		return fmt.Errorf("erro ao enviar email: %v", err)
	}

	return nil
}

// enviarEmailRecuperacao envia um email com o código de recuperação
func enviarEmailRecuperacao(email, codigo string) error {
	// Cria o cliente do Mailjet
	mailjetClient := mailjet.NewMailjetClient(
		os.Getenv("MJ_APIKEY_PUBLIC"),
		os.Getenv("MJ_APIKEY_PRIVATE"),
	)

	// Monta a mensagem
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "naoresponda@paapo.com.br",
				Name:  "PAAPO",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
				},
			},
			Subject:  "Redefinição de Senha - PAAPO",
			TextPart: fmt.Sprintf("Seu código de recuperação é: %s", codigo),
			HTMLPart: fmt.Sprintf(`
				<h2>Redefinição de Senha</h2>
				<p>Olá,</p>
				<p>Recebemos uma solicitação para redefinir sua senha. Use o código abaixo para continuar:</p>
				<div style="font-size: 24px; font-weight: bold; letter-spacing: 5px; margin: 20px 0; padding: 15px; background: #f5f5f5; display: inline-block; border-radius: 5px;">%s</div>
				<p>Este código é válido por 1 hora. Se você não solicitou esta alteração, por favor, ignore este email.</p>
				<p>Atenciosamente,<br>Equipe PAAPO</p>
			`, codigo),
		},
	}

	// Envia o email
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	return err
}

// ConfirmarTrocaSenha confirma a troca de senha com o código de recuperação
func (s *Service) ConfirmarTrocaSenha(email, codigo, novaSenha string) error {
	// Valida os parâmetros
	if email == "" || codigo == "" || novaSenha == "" {
		return errors.New("dados inválidos")
	}

	// Busca o usuário pelo email
	usuario, err := s.db.GetByEmail(email)
	if err != nil {
		// Se não achar, não fala que não existe (por segurança)
		return errors.New("código inválido ou expirado")
	}

	// Verifica se o código está correto
	if usuario.RecoveryCode == nil || *usuario.RecoveryCode != codigo {
		return errors.New("código inválido ou expirado")
	}

	// Verifica se o código não expirou
	if usuario.RecoveryCodeExpiresAt == nil || usuario.RecoveryCodeExpiresAt.Before(time.Now()) {
		return errors.New("código expirado")
	}

	// Gera o hash da nova senha
	hash, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha")
	}

	// Atualiza a senha e limpa o código de recuperação
	usuario.Password = string(hash)
	usuario.RecoveryCode = nil
	usuario.RecoveryCodeExpiresAt = nil

	// Salva as alterações
	if _, err := s.db.Update(usuario); err != nil {
		return errors.New("erro ao atualizar senha")
	}

	return nil
}

func CreateUserWithGroupAssignment(repo Repository, user *User) (*User, error) {
	testCount, _ := repo.CountByTestGroup(true)
	controlCount, _ := repo.CountByTestGroup(false)

	// Alternate automatically
	user.IsTestGroup = controlCount > testCount

	return repo.Create(user)
}
