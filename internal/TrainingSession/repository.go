package trainingsession

import (
	"backend-PAAPO/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(session *models.TrainingSession) error
	GetByID(id string) (*models.TrainingSession, error)
	GetActiveByUserID(userID string) (*models.TrainingSession, error)
	GetAllByUserID(userID string) ([]*models.TrainingSession, error)
	Update(session *models.TrainingSession) error
	Delete(id string) error
}

type trainingSessionRepository struct {
	db *gorm.DB
}

func NewTrainingSessionRepository(db *gorm.DB) Repository {
	return &trainingSessionRepository{db: db}
}

func (r *trainingSessionRepository) Create(session *models.TrainingSession) error {
	return r.db.Create(session).Error
}

func (r *trainingSessionRepository) GetByID(id string) (*models.TrainingSession, error) {
	var session models.TrainingSession
	err := r.db.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *trainingSessionRepository) GetActiveByUserID(userID string) (*models.TrainingSession, error) {
	var session models.TrainingSession
	err := r.db.Where("user_id = ? AND status = ?", userID, models.TrainingSessionActive).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func (r *trainingSessionRepository) GetAllByUserID(userID string) ([]*models.TrainingSession, error) {
	var sessions []*models.TrainingSession
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *trainingSessionRepository) Update(session *models.TrainingSession) error {
	return r.db.Save(session).Error
}

func (r *trainingSessionRepository) Delete(id string) error {
	return r.db.Delete(&models.TrainingSession{}, id).Error
}
