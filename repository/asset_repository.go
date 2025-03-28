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
	FilterAssets(name, categoryName, statusName string) ([]entity.Assets, error)
	FindAllPaginated(offset int, limit int) ([]entity.Assets, error)
	FilterAssetsPaginated(name, categoryName, statusName string, offset int, limit int) ([]entity.Assets, error)
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

func (r *assetRepository) FilterAssets(name, categoryName, statusName string) ([]entity.Assets, error) {
	var assets []entity.Assets
	query := r.db.Preload("AssetCategories").Preload("Statuses")

	if name != "" {
		query = query.Where("asset_name LIKE ?", "%"+name+"%")
	}

	if categoryName != "" {
		query = query.Joins("JOIN asset_categories ON assets.category_id = asset_categories.category_id").
			Where("asset_categories.category_name LIKE ?", "%"+categoryName+"%")
	}

	if statusName != "" {
		query = query.Joins("JOIN statuses ON assets.status_id = statuses.status_id").
			Where("statuses.status_name LIKE ?", "%"+statusName+"%")
	}

	err := query.Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) FindAllPaginated(offset int, limit int) ([]entity.Assets, error) {
	var assets []entity.Assets
	err := r.db.Preload("AssetCategories").Preload("Statuses").
		Offset(offset).Limit(limit).Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func (r *assetRepository) FilterAssetsPaginated(name, categoryName, statusName string, offset int, limit int) ([]entity.Assets, error) {
	var assets []entity.Assets
	query := r.db.Preload("AssetCategories").Preload("Statuses")

	if name != "" {
		query = query.Where("asset_name LIKE ?", "%"+name+"%")
	}
	if categoryName != "" {
		query = query.Joins("JOIN asset_categories ON assets.category_id = asset_categories.category_id").
			Where("asset_categories.category_name LIKE ?", "%"+categoryName+"%")
	}
	if statusName != "" {
		query = query.Joins("JOIN statuses ON assets.status_id = statuses.status_id").
			Where("statuses.status_name LIKE ?", "%"+statusName+"%")
	}

	err := query.Offset(offset).Limit(limit).Find(&assets).Error
	if err != nil {
		return nil, err
	}
	return assets, nil
}
