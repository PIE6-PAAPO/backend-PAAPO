package dto

import "time"

type CreateTrainingSessionDTO struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	Category  *string   `json:"category,omitempty"`
}

type EndTrainingSessionDTO struct {
	Comments string `json:"comments,omitempty"`
}

type TrainingSessionResponseDTO struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	StartTime   time.Time  `json:"start_time"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	Duration    *int64     `json:"duration_seconds,omitempty"`
	Status      string     `json:"status"`
	Comments    *string    `json:"comments,omitempty"`
	Interrupted bool       `json:"interrupted"`
	Category    *string    `json:"category,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
