package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(userID int) (*entity.Users, error)
	FindByEmail(email string) (*entity.Users, error)
	Create(user *entity.Users) error
	FindAll() ([]entity.Users, error)
	Update(user *entity.Users) error
	FilterUsers(name, email string, roleID int) ([]entity.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(userID int) (*entity.Users, error) {
	var user entity.Users
	err := r.db.Preload("Roles").Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.Users, error) {
	var user entity.Users
	err := r.db.Preload("Roles").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *entity.Users) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *entity.Users) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) FilterUsers(name, email string, roleID int) ([]entity.Users, error) {
	var users []entity.Users
	query := r.db.Preload("Roles")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if roleID != 0 {
		query = query.Where("role_id = ?", roleID)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
