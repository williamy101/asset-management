package repository

import (
	"errors"
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type StatusRepository interface {
	Create(status *entity.Statuses) error
	FindAll() ([]entity.Statuses, error)
	FindByID(id int) (*entity.Statuses, error)
	Update(status *entity.Statuses) error
	Delete(id int) error
}

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) StatusRepository {
	return &statusRepository{db: db}
}

func (r *statusRepository) Create(status *entity.Statuses) error {
	result := r.db.Create(status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *statusRepository) FindAll() ([]entity.Statuses, error) {
	var statuses []entity.Statuses
	result := r.db.Find(&statuses)
	if result.Error != nil {
		return nil, result.Error
	}
	return statuses, nil
}

func (r *statusRepository) FindByID(id int) (*entity.Statuses, error) {
	var status entity.Statuses
	result := r.db.First(&status, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("status not found")
		}
		return nil, result.Error
	}
	return &status, nil
}

func (r *statusRepository) Update(status *entity.Statuses) error {
	result := r.db.Save(status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *statusRepository) Delete(id int) error {
	var status entity.Statuses
	result := r.db.Delete(&status, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
