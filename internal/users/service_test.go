package users

import (
	"backend-PAAPO/internal/dto"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type testRepo struct {
	*userRepository
}

func setupTestDB() *userRepository {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return &userRepository{db: db}
}

func TestGroupAssignmentBalance(t *testing.T) {
	repo := setupTestDB()
	service := NewService(repo)

	// Registra 11 usuários
	numUsers := 11
	for i := 0; i < numUsers; i++ {
		_, err := service.Register(dto.RegisterDTO{
			FirstName: "Test",
			LastName:  "User",
			Email:     uuid.New().String() + "@test.com",
			Password:  "password123",
		})
		if err != nil {
			t.Fatalf("Erro ao registrar usuário: %v", err)
		}
	}

	testCount, _ := repo.CountByTestGroup(true)
	controlCount, _ := repo.CountByTestGroup(false)
	diff := testCount - controlCount
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		t.Errorf("A diferença entre os grupos é maior que 1: test=%d, control=%d", testCount, controlCount)
	}
}
