package healthinformation

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(healthInfo *HealthInformation) error
	GetByUserID(userID string) (*HealthInformation, error)
	Update(healthInfo *HealthInformation) error
	Delete(userID string) error
}

type healthInformationRepository struct {
	db *gorm.DB
}

func NewHealthInformationRepository(db *gorm.DB) Repository {
	return &healthInformationRepository{db: db}
}

func (r *healthInformationRepository) Create(healthInfo *HealthInformation) error {
	return r.db.Create(healthInfo).Error
}

func (r *healthInformationRepository) GetByUserID(userID string) (*HealthInformation, error) {
	var healthInfo HealthInformation
	err := r.db.Where("user_id = ?", userID).First(&healthInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &healthInfo, nil
}

func (r *healthInformationRepository) Update(healthInfo *HealthInformation) error {
	return r.db.Save(healthInfo).Error
}

func (r *healthInformationRepository) Delete(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&HealthInformation{}).Error
}
