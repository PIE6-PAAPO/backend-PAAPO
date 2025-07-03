package medicaldata

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *Service {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&MedicalData{})
	repo := NewMedicalDataRepository(db)
	return NewService(repo)
}

func TestCRUDMedicalData(t *testing.T) {
	service := setupTestDB()
	userID := "user-med-1"
	dto := CreateMedicalDataDTO{
		DiagnosisDate:         nil,
		SurgeryDate:           nil,
		SurgeryType:           "Tipo X",
		PhysiotherapyReferral: true,
		SufficientRecovery:    false,
		PresentSequela:        Sequela{Fatigue: true},
		RemainingSequela:      Sequela{},
		ImpactOnDailyLife:     "Moderado",
		AffectsIndependence:   false,
		PostSurgeryActivities: "Atividade Y",
		MedicalConditions:     []string{"Cond1"},
		OtherDiagnoses:        "Diag extra",
		Medications:           "Med1",
	}
	// Criar
	m, err := service.CreateMedicalData(userID, dto)
	if err != nil {
		t.Fatalf("Erro ao criar: %v", err)
	}
	if m.SurgeryType != "Tipo X" {
		t.Errorf("Tipo de cirurgia não bate")
	}
	// Consultar
	m2, err := service.GetMedicalData(userID)
	if err != nil || m2 == nil {
		t.Fatalf("Erro ao consultar: %v", err)
	}
	// Atualizar
	novoTipo := "Tipo Y"
	upd := UpdateMedicalDataDTO{SurgeryType: &novoTipo}
	m3, err := service.UpdateMedicalData(userID, upd)
	if err != nil {
		t.Fatalf("Erro ao atualizar: %v", err)
	}
	if m3.SurgeryType != "Tipo Y" {
		t.Errorf("Atualização não funcionou")
	}
	// Deletar
	err = service.DeleteMedicalData(userID)
	if err != nil {
		t.Fatalf("Erro ao deletar: %v", err)
	}
	m4, _ := service.GetMedicalData(userID)
	if m4 != nil {
		t.Error("Deveria retornar nil após deletar")
	}
}
