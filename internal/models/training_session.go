package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrainingSessionStatus string

const (
	TrainingSessionActive   TrainingSessionStatus = "active"
	TrainingSessionFinished TrainingSessionStatus = "finished"
)

type TrainingSession struct {
	ID        uuid.UUID             `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string                `gorm:"not null" json:"user_id"`
	StartDate time.Time             `gorm:"not null" json:"start_date"`
	StartTime time.Time             `gorm:"not null" json:"start_time"`
	EndDate   *time.Time            `json:"end_date,omitempty"`
	EndTime   *time.Time            `json:"end_time,omitempty"`
	Status    TrainingSessionStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	Comments  *string               `gorm:"type:text" json:"comments,omitempty"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt gorm.DeletedAt        `gorm:"index" json:"-"`
}
