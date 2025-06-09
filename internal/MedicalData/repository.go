package medicaldata

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(medicalData *MedicalData) error
	GetByUserID(userID string) (*MedicalData, error)
	Update(medicalData *MedicalData) error
	Delete(userID string) error
}

type medicalDataRepository struct {
	db *gorm.DB
}

func NewMedicalDataRepository(db *gorm.DB) Repository {
	return &medicalDataRepository{db: db}
}

func (r *medicalDataRepository) Create(medicalData *MedicalData) error {
	return r.db.Create(medicalData).Error
}

func (r *medicalDataRepository) GetByUserID(userID string) (*MedicalData, error) {
	var medicalData MedicalData
	err := r.db.Where("user_id = ?", userID).First(&medicalData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &medicalData, nil
}

func (r *medicalDataRepository) Update(medicalData *MedicalData) error {
	return r.db.Save(medicalData).Error
}

func (r *medicalDataRepository) Delete(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&MedicalData{}).Error
}
