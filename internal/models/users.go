package models

import (
	"backend-PAAPO/pkg"
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	Admin   Role = "admin"
	Patient Role = "patient"
)

type User struct {
	ID                string `gorm:"primaryKey" json:"id"`
	Email             string `gorm:"uniqueIndex" json:"email"`
	Password          string `json:"password"`
	FullName          string `json:"full_name"`
	Birthdate         string `json:"birthdate"`
	PhoneNumber       string `json:"phone_number"`
	Address           string `json:"address"`
	Occupation        string `json:"occupation"`
	IsActive          bool   `json:"is_active"`
	isConfirmed       bool
	ConfirmedAt       time.Time `json:"confirmed_at"`
	recoveryToken     string
	recoveryTokenAt   time.Time
	UpdatedAt         time.Time          `json:"updated_at"`
	RegisteredAt      time.Time          `json:"registered_at"`
	Role              Role               `json:"role"`
	HealthInformation *HealthInformation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"health_information"`
}

// Utilizar para construir um novo usuario
func NewUser(user User) User {
	return User{
		ID:           uuid.New().String(),
		Email:        user.Email,
		Password:     pkg.EncodePassword(user.Password),
		FullName:     user.FullName,
		Birthdate:    user.Birthdate,
		PhoneNumber:  user.PhoneNumber,
		Address:      user.Address,
		Occupation:   user.Occupation,
		Role:         Patient,
		isConfirmed:  false,
		IsActive:     true,
		UpdatedAt:    time.Now(),
		RegisteredAt: time.Now(),
	}
}
