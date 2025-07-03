package trainingsession

import (
	"backend-PAAPO/internal/dto"
	"backend-PAAPO/internal/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTrainingSession(userID string, dto dto.CreateTrainingSessionDTO) (*models.TrainingSession, error) {
	// Check if user already has an active session
	activeSession, err := s.repo.GetActiveByUserID(userID)
	if err != nil {
		return nil, err
	}
	if activeSession != nil {
		return nil, errors.New("user already has an active training session")
	}

	// Create new training session
	session := &models.TrainingSession{
		ID:        uuid.New(),
		UserID:    userID,
		StartDate: dto.StartDate,
		StartTime: dto.StartTime,
		Status:    models.TrainingSessionActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.Create(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) EndTrainingSession(userID string, dto dto.EndTrainingSessionDTO) (*models.TrainingSession, error) {
	// Find active session for the user
	activeSession, err := s.repo.GetActiveByUserID(userID)
	if err != nil {
		return nil, err
	}
	if activeSession == nil {
		return nil, errors.New("no active training session found")
	}

	// End the session
	now := time.Now()
	activeSession.EndDate = &now
	activeSession.EndTime = &now
	activeSession.Status = models.TrainingSessionFinished
	activeSession.UpdatedAt = now

	if dto.Comments != "" {
		activeSession.Comments = &dto.Comments
	}

	err = s.repo.Update(activeSession)
	if err != nil {
		return nil, err
	}

	return activeSession, nil
}

func (s *Service) GetAllTrainingSessions(userID string) ([]*models.TrainingSession, error) {
	return s.repo.GetAllByUserID(userID)
}

func (s *Service) GetActiveTrainingSession(userID string) (*models.TrainingSession, error) {
	return s.repo.GetActiveByUserID(userID)
}
