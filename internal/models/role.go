package models

// Role represents a user's role in the system
type Role string

const (
	// Admin role has full access
	Admin   Role = "admin"
	// Patient role has limited access
	Patient Role = "patient"
)
