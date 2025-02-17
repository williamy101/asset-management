package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type MaintenanceRequestRepository interface {
	Create(request *entity.MaintenanceRequests) error
	FindAll() ([]entity.MaintenanceRequests, error)
	FindByID(requestID int) (*entity.MaintenanceRequests, error)
	FindByUserID(userID int) ([]entity.MaintenanceRequests, error)
	FindByAssetID(assetID int) ([]entity.MaintenanceRequests, error)
	FindLatestByAssetID(assetID int) (*entity.MaintenanceRequests, error)
	FindByStatusID(statusID int) ([]entity.MaintenanceRequests, error)
	Update(request *entity.MaintenanceRequests) error
	UpdateStatus(requestID int, statusID int) error
	Delete(requestID int) error
}

type maintenanceRequestRepository struct {
	db *gorm.DB
}

func NewMaintenanceRequestRepository(db *gorm.DB) MaintenanceRequestRepository {
	return &maintenanceRequestRepository{db: db}
}

func (r *maintenanceRequestRepository) Create(request *entity.MaintenanceRequests) error {
	return r.db.Create(request).Error
}

func (r *maintenanceRequestRepository) FindAll() ([]entity.MaintenanceRequests, error) {
	var requests []entity.MaintenanceRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Find(&requests).Error
	return requests, err
}

func (r *maintenanceRequestRepository) FindByID(requestID int) (*entity.MaintenanceRequests, error) {
	var request entity.MaintenanceRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Where("request_id = ?", requestID).
		First(&request).Error
	return &request, err
}

func (r *maintenanceRequestRepository) FindByUserID(userID int) ([]entity.MaintenanceRequests, error) {
	var requests []entity.MaintenanceRequests
	err := r.db.
		Preload("Assets").
		Preload("Statuses").
		Where("user_id = ?", userID).
		Find(&requests).Error
	return requests, err
}

func (r *maintenanceRequestRepository) FindByAssetID(assetID int) ([]entity.MaintenanceRequests, error) {
	var requests []entity.MaintenanceRequests
	err := r.db.
		Preload("Users").
		Preload("Statuses").
		Where("asset_id = ?", assetID).
		Find(&requests).Error
	return requests, err
}

func (r *maintenanceRequestRepository) FindLatestByAssetID(assetID int) (*entity.MaintenanceRequests, error) {
	var request entity.MaintenanceRequests
	err := r.db.
		Preload("Users").
		Preload("Statuses").
		Where("asset_id = ?", assetID).
		Order("request_date DESC").
		First(&request).Error
	return &request, err
}

func (r *maintenanceRequestRepository) FindByStatusID(statusID int) ([]entity.MaintenanceRequests, error) {
	var requests []entity.MaintenanceRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Where("status_id = ?", statusID).
		Find(&requests).Error
	return requests, err
}

func (r *maintenanceRequestRepository) Update(request *entity.MaintenanceRequests) error {
	return r.db.Save(request).Error
}

func (r *maintenanceRequestRepository) UpdateStatus(requestID int, statusID int) error {
	return r.db.Model(&entity.MaintenanceRequests{}).
		Where("request_id = ?", requestID).
		Update("status_id", statusID).Error
}

func (r *maintenanceRequestRepository) Delete(requestID int) error {
	return r.db.Delete(&entity.MaintenanceRequests{}, requestID).Error
}
