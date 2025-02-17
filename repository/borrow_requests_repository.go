package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type BorrowAssetRequestRepository interface {
	Create(request *entity.BorrowAssetRequests) error
	FindAll() ([]entity.BorrowAssetRequests, error)
	FindByID(requestID int) (*entity.BorrowAssetRequests, error)
	FindByUserID(userID int) ([]entity.BorrowAssetRequests, error)
	FindByAssetID(assetID int) ([]entity.BorrowAssetRequests, error)
	FindByStatus(statusID int) ([]entity.BorrowAssetRequests, error)
	Update(request *entity.BorrowAssetRequests) error
	UpdateStatus(requestID int, statusID int, approvedBy *int) error
	Delete(requestID int) error
}

type borrowAssetRequestRepository struct {
	db *gorm.DB
}

func NewBorrowAssetRequestRepository(db *gorm.DB) BorrowAssetRequestRepository {
	return &borrowAssetRequestRepository{db: db}
}

func (r *borrowAssetRequestRepository) Create(request *entity.BorrowAssetRequests) error {
	return r.db.Create(request).Error
}

func (r *borrowAssetRequestRepository) FindAll() ([]entity.BorrowAssetRequests, error) {
	var requests []entity.BorrowAssetRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Preload("Approver").
		Find(&requests).Error
	return requests, err
}

func (r *borrowAssetRequestRepository) FindByID(requestID int) (*entity.BorrowAssetRequests, error) {
	var request entity.BorrowAssetRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Preload("Approver").
		Where("borrow_request_id = ?", requestID).
		First(&request).Error
	return &request, err
}

func (r *borrowAssetRequestRepository) FindByUserID(userID int) ([]entity.BorrowAssetRequests, error) {
	var requests []entity.BorrowAssetRequests
	err := r.db.
		Preload("Assets").
		Preload("Statuses").
		Where("user_id = ?", userID).
		Find(&requests).Error
	return requests, err
}

func (r *borrowAssetRequestRepository) FindByAssetID(assetID int) ([]entity.BorrowAssetRequests, error) {
	var requests []entity.BorrowAssetRequests
	err := r.db.
		Preload("Users").
		Preload("Statuses").
		Where("asset_id = ?", assetID).
		Find(&requests).Error
	return requests, err
}

func (r *borrowAssetRequestRepository) FindByStatus(statusID int) ([]entity.BorrowAssetRequests, error) {
	var requests []entity.BorrowAssetRequests
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Where("status_id = ?", statusID).
		Find(&requests).Error
	return requests, err
}

func (r *borrowAssetRequestRepository) Update(request *entity.BorrowAssetRequests) error {
	return r.db.Save(request).Error
}

func (r *borrowAssetRequestRepository) UpdateStatus(requestID int, statusID int, approvedBy *int) error {
	updateFields := map[string]interface{}{
		"status_id": statusID,
	}
	if approvedBy != nil {
		updateFields["approved_by"] = approvedBy
	}
	return r.db.Model(&entity.BorrowAssetRequests{}).
		Where("borrow_request_id = ?", requestID).
		Updates(updateFields).Error
}

func (r *borrowAssetRequestRepository) Delete(requestID int) error {
	return r.db.Delete(&entity.BorrowAssetRequests{}, requestID).Error
}
