package healthinformation

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *Service {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&HealthInformation{})
	repo := NewHealthInformationRepository(db)
	return NewService(repo)
}

func TestCRUDHealthInfo(t *testing.T) {
	service := setupTestDB()
	userID := "user-test-1"
	// Criar
	dto := CreateHealthInformationDTO{
		Weight:                    70.5,
		Height:                    1.75,
		Smoker:                    false,
		AlcoholConsumption:        "Ocasional",
		PhysicalActivityFrequency: "3x semana",
		PhysicalActivityType:      "Caminhada",
		UserID:                    userID,
	}
	h, err := service.CreateHealthInfo(dto)
	if err != nil {
		t.Fatalf("Erro ao criar: %v", err)
	}
	if h.Weight != 70.5 {
		t.Errorf("Peso não bate")
	}
	// Consultar
	h2, err := service.GetHealthInfo(userID)
	if err != nil || h2 == nil {
		t.Fatalf("Erro ao consultar: %v", err)
	}
	// Atualizar
	newWeight := 72.0
	upd := UpdateHealthInformationDTO{Weight: &newWeight}
	h3, err := service.UpdateHealthInfo(userID, upd)
	if err != nil {
		t.Fatalf("Erro ao atualizar: %v", err)
	}
	if h3.Weight != 72.0 {
		t.Errorf("Atualização não funcionou")
	}
	// Deletar
	err = service.DeleteHealthInfo(userID)
	if err != nil {
		t.Fatalf("Erro ao deletar: %v", err)
	}
	h4, _ := service.GetHealthInfo(userID)
	if h4 != nil {
		t.Error("Deveria retornar nil após deletar")
	}
}
