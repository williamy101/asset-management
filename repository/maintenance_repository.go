package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type MaintenanceRepository interface {
	Create(maintenance *entity.Maintenances) error
	FindAll() ([]entity.Maintenances, error)
	FindByID(maintenanceID int) (*entity.Maintenances, error)
	Update(maintenance *entity.Maintenances) error
	Delete(maintenanceID int) error
}

type maintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) MaintenanceRepository {
	return &maintenanceRepository{db: db}
}

func (r *maintenanceRepository) Create(maintenance *entity.Maintenances) error {
	return r.db.Create(maintenance).Error
}

func (r *maintenanceRepository) FindAll() ([]entity.Maintenances, error) {
	var maintenances []entity.Maintenances
	err := r.db.Find(&maintenances).Error
	if err != nil {
		return nil, err
	}
	return maintenances, nil
}

func (r *maintenanceRepository) FindByID(maintenanceID int) (*entity.Maintenances, error) {
	var maintenance entity.Maintenances
	err := r.db.First(&maintenance, maintenanceID).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

func (r *maintenanceRepository) Update(maintenance *entity.Maintenances) error {
	return r.db.Save(maintenance).Error
}

func (r *maintenanceRepository) Delete(maintenanceID int) error {
	return r.db.Delete(&entity.Maintenances{}, maintenanceID).Error
}
