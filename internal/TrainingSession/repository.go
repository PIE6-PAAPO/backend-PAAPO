package trainingsession

import (
	"backend-PAAPO/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(session *models.TrainingSession) error
	GetByID(id string) (*models.TrainingSession, error)
	GetActiveByUserID(userID string) (*models.TrainingSession, error)
	GetAllByUserID(userID string, filters SessionFilters, page, limit int) ([]*models.TrainingSession, int64, error)
	Update(session *models.TrainingSession) error
	Delete(id string) error
	FindAbandonedSessions(maxDurationHours int) ([]*models.TrainingSession, error)
}

type SessionFilters struct {
	From     *time.Time
	To       *time.Time
	Category *string
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

func (r *trainingSessionRepository) GetAllByUserID(userID string, filters SessionFilters, page, limit int) ([]*models.TrainingSession, int64, error) {
	var sessions []*models.TrainingSession
	var count int64
	query := r.db.Model(&models.TrainingSession{}).Where("user_id = ?", userID)

	if filters.From != nil {
		query = query.Where("start_date >= ?", filters.From)
	}
	if filters.To != nil {
		query = query.Where("start_date <= ?", filters.To)
	}
	if filters.Category != nil {
		query = query.Where("category = ?", filters.Category)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Order("created_at DESC").Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return sessions, count, nil
}

func (r *trainingSessionRepository) Update(session *models.TrainingSession) error {
	return r.db.Save(session).Error
}

func (r *trainingSessionRepository) Delete(id string) error {
	return r.db.Delete(&models.TrainingSession{}, id).Error
}

func (r *trainingSessionRepository) FindAbandonedSessions(maxDurationHours int) ([]*models.TrainingSession, error) {
	var sessions []*models.TrainingSession
	cutoff := time.Now().Add(-time.Duration(maxDurationHours) * time.Hour)
	err := r.db.Where("status = ? AND start_time < ?", models.TrainingSessionActive, cutoff).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}
