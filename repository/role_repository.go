package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindByID(roleID int) (*entity.Roles, error)
	FindAll() ([]entity.Roles, error)
	Create(role *entity.Roles) error
	DeleteByID(roleID int) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindByID(roleID int) (*entity.Roles, error) {
	var role entity.Roles
	err := r.db.First(&role, roleID).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindAll() ([]entity.Roles, error) {
	var roles []entity.Roles
	err := r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) Create(role *entity.Roles) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) DeleteByID(roleID int) error {
	return r.db.Delete(&entity.Roles{}, roleID).Error
}
