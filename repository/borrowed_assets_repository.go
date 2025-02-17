package repository

import (
	"go-asset-management/entity"
	"time"

	"gorm.io/gorm"
)

type BorrowedAssetRepository interface {
	Create(borrowed *entity.BorrowedAssets) error
	FindAll() ([]entity.BorrowedAssets, error)
	FindByID(borrowID int) (*entity.BorrowedAssets, error)
	FindByUserID(userID int) ([]entity.BorrowedAssets, error)
	FindByAssetID(assetID int) ([]entity.BorrowedAssets, error)
	FindByStatus(statusID int) ([]entity.BorrowedAssets, error)
	Update(borrowed *entity.BorrowedAssets) error
	UpdateReturnDate(borrowID int, returnDate *time.Time, statusID int) error
	Delete(borrowID int) error
}

type borrowedAssetRepository struct {
	db *gorm.DB
}

func NewBorrowedAssetRepository(db *gorm.DB) BorrowedAssetRepository {
	return &borrowedAssetRepository{db: db}
}

func (r *borrowedAssetRepository) Create(borrowed *entity.BorrowedAssets) error {
	return r.db.Create(borrowed).Error
}

func (r *borrowedAssetRepository) FindAll() ([]entity.BorrowedAssets, error) {
	var borrowedAssets []entity.BorrowedAssets
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Find(&borrowedAssets).Error
	return borrowedAssets, err
}

func (r *borrowedAssetRepository) FindByID(borrowID int) (*entity.BorrowedAssets, error) {
	var borrowed entity.BorrowedAssets
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Preload("Statuses").
		Where("borrow_id = ?", borrowID).
		First(&borrowed).Error
	return &borrowed, err
}

func (r *borrowedAssetRepository) FindByUserID(userID int) ([]entity.BorrowedAssets, error) {
	var borrowedAssets []entity.BorrowedAssets
	err := r.db.
		Preload("Assets").
		Preload("Statuses").
		Where("user_id = ?", userID).
		Find(&borrowedAssets).Error
	return borrowedAssets, err
}

func (r *borrowedAssetRepository) FindByAssetID(assetID int) ([]entity.BorrowedAssets, error) {
	var borrowedAssets []entity.BorrowedAssets
	err := r.db.
		Preload("Users").
		Preload("Statuses").
		Where("asset_id = ?", assetID).
		Find(&borrowedAssets).Error
	return borrowedAssets, err
}

func (r *borrowedAssetRepository) FindByStatus(statusID int) ([]entity.BorrowedAssets, error) {
	var borrowedAssets []entity.BorrowedAssets
	err := r.db.
		Preload("Assets").
		Preload("Users").
		Where("status_id = ?", statusID).
		Find(&borrowedAssets).Error
	return borrowedAssets, err
}

func (r *borrowedAssetRepository) Update(borrowed *entity.BorrowedAssets) error {
	return r.db.Save(borrowed).Error
}

func (r *borrowedAssetRepository) UpdateReturnDate(borrowID int, returnDate *time.Time, statusID int) error {
	updateFields := map[string]interface{}{
		"status_id": statusID,
	}
	if returnDate != nil {
		updateFields["return_date"] = returnDate
	}
	return r.db.Model(&entity.BorrowedAssets{}).
		Where("borrow_id = ?", borrowID).
		Updates(updateFields).Error
}

func (r *borrowedAssetRepository) Delete(borrowID int) error {
	return r.db.Delete(&entity.BorrowedAssets{}, borrowID).Error
}
