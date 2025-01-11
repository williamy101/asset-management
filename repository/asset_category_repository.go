package repository

import (
	"errors"
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type AssetCategoryRepository interface {
	Create(category *entity.AssetCategories) error
	FindAll() ([]entity.AssetCategories, error)
	FindByID(id int) (*entity.AssetCategories, error)
	Update(category *entity.AssetCategories) error
	Delete(id int) error
}

type assetCategoryRepository struct {
	db *gorm.DB
}

func NewAssetCategoryRepository(db *gorm.DB) AssetCategoryRepository {
	return &assetCategoryRepository{db: db}
}

func (r *assetCategoryRepository) Create(category *entity.AssetCategories) error {
	result := r.db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *assetCategoryRepository) FindAll() ([]entity.AssetCategories, error) {
	var categories []entity.AssetCategories
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *assetCategoryRepository) FindByID(id int) (*entity.AssetCategories, error) {
	var category entity.AssetCategories
	result := r.db.First(&category, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, result.Error
	}
	return &category, nil
}

func (r *assetCategoryRepository) Update(category *entity.AssetCategories) error {
	result := r.db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *assetCategoryRepository) Delete(id int) error {
	var category entity.AssetCategories
	result := r.db.Delete(&category, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
