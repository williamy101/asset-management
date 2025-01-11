package service

import (
	"go-asset-management/entity"
	"go-asset-management/repository"
)

type AssetCategoryService interface {
	Create(categoryName string) error
	GetAll() ([]entity.AssetCategories, error)
	GetByID(id int) (*entity.AssetCategories, error)
	Update(id int, categoryName string) error
	Delete(id int) error
}

type assetCategoryService struct {
	repo repository.AssetCategoryRepository
}

func NewAssetCategoryService(repo repository.AssetCategoryRepository) AssetCategoryService {
	return &assetCategoryService{repo: repo}
}

func (s *assetCategoryService) Create(categoryName string) error {
	category := &entity.AssetCategories{
		CategoryName: categoryName,
	}
	return s.repo.Create(category)
}

func (s *assetCategoryService) GetAll() ([]entity.AssetCategories, error) {
	return s.repo.FindAll()
}

func (s *assetCategoryService) GetByID(id int) (*entity.AssetCategories, error) {
	return s.repo.FindByID(id)
}

func (s *assetCategoryService) Update(id int, categoryName string) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	category.CategoryName = categoryName
	return s.repo.Update(category)
}

func (s *assetCategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
