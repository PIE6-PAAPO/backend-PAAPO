package users

import (
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type Repository interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByProvider(provider, providerId string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	GetByConfirmationCode(code string) (*User, error)
	GetByRecoveryCode(code string) (*User, error)
	GetByEmailChangeCode(code string) (*User, error)
	GetByEmailChangeCodeUsedAt(code string) (*User, error)
	GetByEmailChangeCodeExpiresAt(code string) (*User, error)
	GetByEmailChangeCodeUsed(code string) (*User, error)
	GetByEmailChangeCodeUsedAtAndUsed(code string, used bool) (*User, error)
	GetByEmailChangeCodeExpiresAtAndUsed(code string, used bool) (*User, error)
	GetByEmailChangeCodeUsedAtAndExpiresAt(code string, used bool, expiresAt *time.Time) (*User, error)
	CountByTestGroup(isTestGroup bool) (int64, error)
	GetDB() *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

// GetDB returns the underlying database connection for transactions
func (u *userRepository) GetDB() *gorm.DB {
	return u.db
}

// Create implements Repository.
func (u *userRepository) Create(user *User) (*User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete implements Repository.
func (u *userRepository) Delete(id string) error {
	if err := u.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetByConfirmationCode implements Repository.
func (u *userRepository) GetByConfirmationCode(code string) (*User, error) {
	var user User
	if err := u.db.Where("confirmation_code = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail implements Repository.
func (u *userRepository) GetByEmail(email string) (*User, error) {
	var user User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCode implements Repository.
func (u *userRepository) GetByEmailChangeCode(code string) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeExpiresAt implements Repository.
func (u *userRepository) GetByEmailChangeCodeExpiresAt(code string) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_expires_at = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeExpiresAtAndUsed implements Repository.
func (u *userRepository) GetByEmailChangeCodeExpiresAtAndUsed(code string, used bool) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_expires_at = ? AND email_change_code_used = ?", code, used).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeUsed implements Repository.
func (u *userRepository) GetByEmailChangeCodeUsed(code string) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_used = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeUsedAt implements Repository.
func (u *userRepository) GetByEmailChangeCodeUsedAt(code string) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_used_at = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeUsedAtAndExpiresAt implements Repository.
func (u *userRepository) GetByEmailChangeCodeUsedAtAndExpiresAt(code string, used bool, expiresAt *time.Time) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_used_at = ? AND email_change_code_expires_at = ?", code, expiresAt).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmailChangeCodeUsedAtAndUsed implements Repository.
func (u *userRepository) GetByEmailChangeCodeUsedAtAndUsed(code string, used bool) (*User, error) {
	var user User
	if err := u.db.Where("email_change_code_used_at = ? AND email_change_code_used = ?", code, used).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetById implements Repository.
func (u *userRepository) GetById(id string) (*User, error) {
	var user User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByProvider implements Repository.
func (u *userRepository) GetByProvider(provider string, providerId string) (*User, error) {
	var user User
	if err := u.db.Where("provider = ? AND provider_id = ?", provider, providerId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByRecoveryCode implements Repository.
func (u *userRepository) GetByRecoveryCode(code string) (*User, error) {
	var user User
	if err := u.db.Where("recovery_code = ?", code).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements Repository.
func (u *userRepository) Update(user *User) (*User, error) {
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) CountByTestGroup(isTestGroup bool) (int64, error) {
	var count int64
	err := u.db.
		Model(&User{}).
		Where("is_test_group = ?", isTestGroup).
		Count(&count).Error

	return count, err
}
