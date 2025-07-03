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
	activeSession, err := s.repo.GetActiveByUserID(userID)
	if err != nil {
		return nil, err
	}
	if activeSession != nil {
		return activeSession, &ConflictError{Session: activeSession}
	}

	session := &models.TrainingSession{
		ID:        uuid.New().String(),
		UserID:    userID,
		StartDate: dto.StartDate,
		StartTime: dto.StartTime,
		Status:    models.TrainingSessionActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Category:  dto.Category,
	}

	err = s.repo.Create(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) EndTrainingSession(userID string, dto dto.EndTrainingSessionDTO) (*models.TrainingSession, error) {
	activeSession, err := s.repo.GetActiveByUserID(userID)
	if err != nil {
		return nil, err
	}
	if activeSession == nil {
		return nil, errors.New("no active training session found")
	}

	now := time.Now()
	activeSession.EndDate = &now
	activeSession.EndTime = &now
	activeSession.Status = models.TrainingSessionFinished
	activeSession.UpdatedAt = now

	if dto.Comments != "" {
		activeSession.Comments = &dto.Comments
	}

	// Calculate duration in seconds
	if activeSession.StartTime.Before(now) {
		dur := int64(now.Sub(activeSession.StartTime).Seconds())
		activeSession.Duration = &dur
	}

	err = s.repo.Update(activeSession)
	if err != nil {
		return nil, err
	}

	return activeSession, nil
}

func (s *Service) GetAllTrainingSessions(userID string, filters SessionFilters, page, limit int) ([]*models.TrainingSession, int64, error) {
	return s.repo.GetAllByUserID(userID, filters, page, limit)
}

func (s *Service) GetActiveTrainingSession(userID string) (*models.TrainingSession, error) {
	return s.repo.GetActiveByUserID(userID)
}

// Abandoned session detection (over 12 hours)
func (s *Service) FlagAbandonedSessions() error {
	sessions, err := s.repo.FindAbandonedSessions(12)
	if err != nil {
		return err
	}
	for _, session := range sessions {
		session.Status = models.TrainingSessionInterrupted
		session.Interrupted = true
		session.UpdatedAt = time.Now()
		dur := int64(time.Since(session.StartTime).Seconds())
		session.Duration = &dur
		s.repo.Update(session)
	}
	return nil
}

// Reporting helpers
func (s *Service) TotalHoursTrained(userID string, from, to time.Time) (float64, error) {
	filters := SessionFilters{From: &from, To: &to}
	sessions, _, err := s.repo.GetAllByUserID(userID, filters, 0, 0)
	if err != nil {
		return 0, err
	}
	total := int64(0)
	for _, s := range sessions {
		if s.Duration != nil {
			total += *s.Duration
		}
	}
	return float64(total) / 3600.0, nil
}

func (s *Service) AverageSessionDuration(userID string, from, to time.Time) (float64, error) {
	filters := SessionFilters{From: &from, To: &to}
	sessions, _, err := s.repo.GetAllByUserID(userID, filters, 0, 0)
	if err != nil {
		return 0, err
	}
	total := int64(0)
	count := 0
	for _, s := range sessions {
		if s.Duration != nil {
			total += *s.Duration
			count++
		}
	}
	if count == 0 {
		return 0, nil
	}
	return float64(total) / float64(count) / 60.0, nil // in minutes
}

// For future: session count per category
func (s *Service) SessionCountPerCategory(userID string) (map[string]int, error) {
	filters := SessionFilters{}
	sessions, _, err := s.repo.GetAllByUserID(userID, filters, 0, 0)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int)
	for _, s := range sessions {
		if s.Category != nil {
			result[*s.Category]++
		}
	}
	return result, nil
}

// Conflict error for duplicate active session
// (for handler to return 409)
type ConflictError struct {
	Session *models.TrainingSession
}

func (e *ConflictError) Error() string {
	return "user already has an active training session"
}
