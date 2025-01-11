package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(asset *entity.Assets) error
	FindAll() ([]entity.Assets, error)
	FindByID(assetID int) (*entity.Assets, error)
	Update(asset *entity.Assets) error
	Delete(assetID int) error
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset *entity.Assets) error {
	return r.db.Create(asset).Error
}

func (r *assetRepository) FindAll() ([]entity.Assets, error) {
	var assets []entity.Assets
	err := r.db.Preload("AssetCategories").Preload("Statuses").Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) FindByID(assetID int) (*entity.Assets, error) {
	var asset entity.Assets
	err := r.db.Preload("AssetCategories").Preload("Statuses").First(&asset, assetID).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *assetRepository) Update(asset *entity.Assets) error {
	return r.db.Save(asset).Error
}

func (r *assetRepository) Delete(assetID int) error {
	return r.db.Delete(&entity.Assets{}, assetID).Error
}
