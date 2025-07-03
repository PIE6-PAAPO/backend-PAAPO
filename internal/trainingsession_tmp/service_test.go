package trainingsession

import (
	"backend-PAAPO/internal/dto"
	"backend-PAAPO/internal/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *Service {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.TrainingSession{})
	repo := NewTrainingSessionRepository(db)
	return NewService(repo)
}

func TestCriarEEncerrarSessao(t *testing.T) {
	service := setupTestDB()
	userID := uuid.New().String()
	dtoCriar := dto.CreateTrainingSessionDTO{
		StartDate: time.Now(),
		StartTime: time.Now(),
		Category:  nil,
	}
	sessao, err := service.CreateTrainingSession(userID, dtoCriar)
	if err != nil {
		t.Fatalf("Erro ao criar sessão: %v", err)
	}
	if sessao.Status != models.TrainingSessionActive {
		t.Errorf("Sessão não ficou ativa ao criar")
	}

	// Encerrar sessão
	_, err = service.EndTrainingSession(userID, dto.EndTrainingSessionDTO{Comments: "Fim"})
	if err != nil {
		t.Fatalf("Erro ao encerrar sessão: %v", err)
	}
	// Não deve permitir nova sessão ativa
	_, err = service.CreateTrainingSession(userID, dtoCriar)
	if err != nil && err.Error() != "user already has an active training session" {
		t.Errorf("Deveria bloquear nova sessão ativa, erro: %v", err)
	}
}

func TestRelatorioTotalHoras(t *testing.T) {
	service := setupTestDB()
	userID := uuid.New().String()
	now := time.Now()
	// Cria e encerra 2 sessões de 1h
	for i := 0; i < 2; i++ {
		sessao, err := service.CreateTrainingSession(userID, dto.CreateTrainingSessionDTO{
			StartDate: now,
			StartTime: now.Add(time.Duration(i) * time.Hour),
			Category:  nil,
		})
		if err != nil {
			t.Fatalf("Erro ao criar sessão: %v", err)
		}
		// Simula duração de 1h
		end := sessao.StartTime.Add(1 * time.Hour)
		sessao.EndTime = &end
		dur := int64(3600)
		sessao.Duration = &dur
		sessao.Status = models.TrainingSessionFinished
		service.repo.Update(sessao)
	}
	total, err := service.TotalHoursTrained(userID, now.Add(-1*time.Hour), now.Add(3*time.Hour))
	if err != nil {
		t.Fatalf("Erro no relatório: %v", err)
	}
	if total != 2 {
		t.Errorf("Total de horas deveria ser 2, veio %.2f", total)
	}
}
