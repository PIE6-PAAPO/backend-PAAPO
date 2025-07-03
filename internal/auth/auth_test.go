package auth

import (
	"backend-PAAPO/internal/dto"
	"backend-PAAPO/internal/users"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *users.Service {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&users.User{})
	repo := users.NewUserRepository(db)
	return users.NewService(repo)
}

func TestRegisterAndLogin(t *testing.T) {
	service := setupTestDB()
	reg := dto.RegisterDTO{
		FirstName: "Maria",
		LastName:  "Teste",
		Email:     "maria@teste.com",
		Password:  "senha12345",
	}
	user, err := service.Register(reg)
	if err != nil {
		t.Fatalf("Erro ao registrar: %v", err)
	}
	if user.Email != reg.Email {
		t.Errorf("Email não bate: %s", user.Email)
	}

	login := dto.LoginDTO{
		Email:    reg.Email,
		Password: reg.Password,
	}
	user2, err := service.Login(login)
	if err != nil {
		t.Fatalf("Erro ao fazer login: %v", err)
	}
	if user2.Email != reg.Email {
		t.Errorf("Login retornou usuário errado")
	}
}

func TestLoginComSenhaErrada(t *testing.T) {
	service := setupTestDB()
	reg := dto.RegisterDTO{
		FirstName: "João",
		LastName:  "Teste",
		Email:     "joao@teste.com",
		Password:  "senha12345",
	}
	_, err := service.Register(reg)
	if err != nil {
		t.Fatalf("Erro ao registrar: %v", err)
	}
	login := dto.LoginDTO{
		Email:    reg.Email,
		Password: "errada",
	}
	_, err = service.Login(login)
	if err == nil {
		t.Error("Login com senha errada deveria falhar")
	}
}
