package users

import (
	healthinformation "backend-PAAPO/internal/HealthInformation"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin   Role = "admin"
	Patient Role = "patient"
)

type User struct {
	ID                        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
	FirstName                 string         `gorm:"type:varchar(255);not null"`
	LastName                  string         `gorm:"type:varchar(255);not null"`
	Email                     string         `gorm:"type:varchar(255);not null;unique"`
	IsTestGroup               bool           `gorm:"column:is_test_group"`
	Password                  string         `gorm:"type:varchar(255);not null"`
	Role                      string         `gorm:"type:varchar(255);not null"`
	IsActive                  bool           `gorm:"not null;default:true"`
	IsConfirmed               bool           `gorm:"not null;default:false"`
	LastUsedEmail             string         `gorm:"type:varchar(255);not null"`
	ConfirmationCode          *string        `gorm:"type:varchar(255)"`
	ConfirmationCodeExpiresAt *time.Time
	RecoveryCode              *string `gorm:"type:varchar(255)"`
	RecoveryCodeExpiresAt     *time.Time
	EmailChangeCode           *string `gorm:"type:varchar(255)"`
	EmailChangeCodeExpiresAt  *time.Time
	EmailChangeCodeUsedAt     *time.Time
	EmailChangeCodeUsed       bool                                 `gorm:"not null;default:false"`
	HealthInformation         *healthinformation.HealthInformation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"health_information"`
}
