package medicaldata

import (
	"encoding/json"
	"errors"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type CreateMedicalDataDTO struct {
	DiagnosisDate         *time.Time `json:"diagnosis_date,omitempty"`
	SurgeryDate           *time.Time `json:"surgery_date,omitempty"`
	SurgeryType           string     `json:"surgery_type,omitempty"`
	PhysiotherapyReferral bool       `json:"physiotherapy_referral"`
	PhysiotherapyDuration *time.Time `json:"physiotherapy_duration,omitempty"`
	SufficientRecovery    bool       `json:"sufficient_recovery"`
	PresentSequela        Sequela    `json:"present_sequela,omitempty"`
	RemainingSequela      Sequela    `json:"remaining_sequela,omitempty"`
	ImpactOnDailyLife     string     `json:"impact_on_daily_life,omitempty"`
	AffectsIndependence   bool       `json:"affects_independence"`
	PostSurgeryActivities string     `json:"post_surgery_activities,omitempty"`
	MedicalConditions     []string   `json:"medical_conditions,omitempty"`
	OtherDiagnoses        string     `json:"other_diagnoses,omitempty"`
	Medications           string     `json:"medications,omitempty"`
}

type UpdateMedicalDataDTO struct {
	DiagnosisDate         *time.Time `json:"diagnosis_date,omitempty"`
	SurgeryDate           *time.Time `json:"surgery_date,omitempty"`
	SurgeryType           *string    `json:"surgery_type,omitempty"`
	PhysiotherapyReferral *bool      `json:"physiotherapy_referral,omitempty"`
	PhysiotherapyDuration *time.Time `json:"physiotherapy_duration,omitempty"`
	SufficientRecovery    *bool      `json:"sufficient_recovery,omitempty"`
	PresentSequela        *Sequela   `json:"present_sequela,omitempty"`
	RemainingSequela      *Sequela   `json:"remaining_sequela,omitempty"`
	ImpactOnDailyLife     *string    `json:"impact_on_daily_life,omitempty"`
	AffectsIndependence   *bool      `json:"affects_independence,omitempty"`
	PostSurgeryActivities *string    `json:"post_surgery_activities,omitempty"`
	MedicalConditions     []string   `json:"medical_conditions,omitempty"`
	OtherDiagnoses        *string    `json:"other_diagnoses,omitempty"`
	Medications           *string    `json:"medications,omitempty"`
}

func (s *Service) CreateMedicalData(userID string, dto CreateMedicalDataDTO) (*MedicalData, error) {
	// Convert medical conditions to JSON
	mcJSON, marshalErr := json.Marshal(dto.MedicalConditions)
	if marshalErr != nil {
		return nil, marshalErr
	}

	medicalData := &MedicalData{
		UserID:                userID,
		DiagnosisDate:         dto.DiagnosisDate,
		SurgeryDate:           dto.SurgeryDate,
		SurgeryType:           dto.SurgeryType,
		PhysiotherapyReferral: dto.PhysiotherapyReferral,
		PhysiotherapyDuration: dto.PhysiotherapyDuration,
		SufficientRecovery:    dto.SufficientRecovery,
		PresentSequela:        dto.PresentSequela,
		RemainingSequela:      dto.RemainingSequela,
		ImpactOnDailyLife:     dto.ImpactOnDailyLife,
		AffectsIndependence:   dto.AffectsIndependence,
		PostSurgeryActivities: dto.PostSurgeryActivities,
		MedicalConditions:     mcJSON,
		OtherDiagnoses:        dto.OtherDiagnoses,
		Medications:           dto.Medications,
	}

	err := s.repo.Create(medicalData)
	if err != nil {
		return nil, err
	}

	return medicalData, nil
}

func (s *Service) GetMedicalData(userID string) (*MedicalData, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateMedicalData(userID string, dto UpdateMedicalDataDTO) (*MedicalData, error) {
	medicalData, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if medicalData == nil {
		return nil, errors.New("medical data not found")
	}

	if dto.DiagnosisDate != nil {
		medicalData.DiagnosisDate = dto.DiagnosisDate
	}
	if dto.SurgeryDate != nil {
		medicalData.SurgeryDate = dto.SurgeryDate
	}
	if dto.SurgeryType != nil {
		medicalData.SurgeryType = *dto.SurgeryType
	}
	if dto.PhysiotherapyReferral != nil {
		medicalData.PhysiotherapyReferral = *dto.PhysiotherapyReferral
	}
	if dto.PhysiotherapyDuration != nil {
		medicalData.PhysiotherapyDuration = dto.PhysiotherapyDuration
	}
	if dto.SufficientRecovery != nil {
		medicalData.SufficientRecovery = *dto.SufficientRecovery
	}
	if dto.PresentSequela != nil {
		medicalData.PresentSequela = *dto.PresentSequela
	}
	if dto.RemainingSequela != nil {
		medicalData.RemainingSequela = *dto.RemainingSequela
	}
	if dto.ImpactOnDailyLife != nil {
		medicalData.ImpactOnDailyLife = *dto.ImpactOnDailyLife
	}
	if dto.AffectsIndependence != nil {
		medicalData.AffectsIndependence = *dto.AffectsIndependence
	}
	if dto.PostSurgeryActivities != nil {
		medicalData.PostSurgeryActivities = *dto.PostSurgeryActivities
	}
	if dto.MedicalConditions != nil {
		mcJSON, jsonErr := json.Marshal(dto.MedicalConditions)
		if jsonErr != nil {
			return nil, jsonErr
		}
		medicalData.MedicalConditions = mcJSON
	}
	if dto.OtherDiagnoses != nil {
		medicalData.OtherDiagnoses = *dto.OtherDiagnoses
	}
	if dto.Medications != nil {
		medicalData.Medications = *dto.Medications
	}

	err = s.repo.Update(medicalData)
	if err != nil {
		return nil, err
	}

	return medicalData, nil
}

func (s *Service) DeleteMedicalData(userID string) error {
	return s.repo.Delete(userID)
}
