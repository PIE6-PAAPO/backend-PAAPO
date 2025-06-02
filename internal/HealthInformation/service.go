package healthinformation

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type CreateHealthInformationDTO struct {
	Weight                    float64 `json:"weight"`
	Height                    float64 `json:"height"`
	Smoker                    bool    `json:"smoker"`
	AlcoholConsumption        string  `json:"alcohol_consumption"`
	PhysicalActivityFrequency string  `json:"physical_activity_frequency"`
	PhysicalActivityType      string  `json:"physical_activity_type"`
	UserID                    string  `json:"user_id"`
}

type UpdateHealthInformationDTO struct {
	Weight                    *float64 `json:"weight,omitempty"`
	Height                    *float64 `json:"height,omitempty"`
	Smoker                    *bool    `json:"smoker,omitempty"`
	AlcoholConsumption        *string  `json:"alcohol_consumption,omitempty"`
	PhysicalActivityFrequency *string  `json:"physical_activity_frequency,omitempty"`
	PhysicalActivityType      *string  `json:"physical_activity_type,omitempty"`
}

func (s *Service) CreateHealthInfo(dto CreateHealthInformationDTO) (*HealthInformation, error) {
	healthInfo := &HealthInformation{
		Weight:                    dto.Weight,
		Height:                    dto.Height,
		Smoker:                    dto.Smoker,
		AlcoholConsumption:        dto.AlcoholConsumption,
		PhysicalActivityFrequency: dto.PhysicalActivityFrequency,
		PhysicalActivityType:      dto.PhysicalActivityType,
		UserID:                    dto.UserID,
	}

	err := s.repo.Create(healthInfo)
	if err != nil {
		return nil, err
	}

	return healthInfo, nil
}

func (s *Service) GetHealthInfo(userID string) (*HealthInformation, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateHealthInfo(userID string, dto UpdateHealthInformationDTO) (*HealthInformation, error) {
	healthInfo, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if healthInfo == nil {
		return nil, errors.New("health information not found")
	}

	//Atualizando os campos que chegaram com a DTO
	if dto.Weight != nil {
		healthInfo.Weight = *dto.Weight
	}
	if dto.Height != nil {
		healthInfo.Height = *dto.Height
	}
	if dto.Smoker != nil {
		healthInfo.Smoker = *dto.Smoker
	}
	if dto.AlcoholConsumption != nil {
		healthInfo.AlcoholConsumption = *dto.AlcoholConsumption
	}
	if dto.PhysicalActivityFrequency != nil {
		healthInfo.PhysicalActivityFrequency = *dto.PhysicalActivityFrequency
	}
	if dto.PhysicalActivityType != nil {
		healthInfo.PhysicalActivityType = *dto.PhysicalActivityType
	}

	err = s.repo.Update(healthInfo)
	if err != nil {
		return nil, err
	}

	return healthInfo, nil
}

func (s *Service) DeleteHealthInfo(userID string) error {
	return s.repo.Delete(userID)
}
