package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(userID int) (*entity.Users, error)
	FindByEmail(email string) (*entity.Users, error)
	Create(user *entity.Users) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(userID int) (*entity.Users, error) {
	var user entity.Users
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.Users, error) {
	var user entity.Users
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *entity.Users) error {
	return r.db.Create(user).Error
}
