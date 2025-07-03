package models

import (
	"time"

	"gorm.io/gorm"
)

type TrainingSessionStatus string

const (
	TrainingSessionActive      TrainingSessionStatus = "active"
	TrainingSessionFinished    TrainingSessionStatus = "finished"
	TrainingSessionInterrupted TrainingSessionStatus = "interrupted"
)

type TrainingSession struct {
	ID          string                `gorm:"primaryKey" json:"id"`
	UserID      string                `gorm:"not null" json:"user_id"`
	StartDate   time.Time             `gorm:"not null" json:"start_date"`
	StartTime   time.Time             `gorm:"not null" json:"start_time"`
	EndDate     *time.Time            `json:"end_date,omitempty"`
	EndTime     *time.Time            `json:"end_time,omitempty"`
	Duration    *int64                `json:"duration_seconds,omitempty"` // Duration in seconds
	Status      TrainingSessionStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	Comments    *string               `gorm:"type:text" json:"comments,omitempty"`
	Interrupted bool                  `gorm:"not null;default:false" json:"interrupted"`
	Category    *string               `gorm:"type:varchar(50)" json:"category,omitempty"` // For future use
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   gorm.DeletedAt        `gorm:"index" json:"-"`
}
