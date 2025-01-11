package service

import (
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"
)

type AssetService interface {
	CreateAsset(assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error
	GetAllAssets() ([]entity.Assets, error)
	GetAssetByID(assetID int) (*entity.Assets, error)
	UpdateAsset(assetID int, assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error
	DeleteAsset(assetID int) error
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(assetRepo repository.AssetRepository) AssetService {
	return &assetService{assetRepo: assetRepo}
}

func (s *assetService) CreateAsset(assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error {
	parsedLastMaintenance, err := time.Parse("2006-01-02", lastMaintenance)
	if err != nil {
		return err
	}
	parsedNextMaintenance, err := time.Parse("2006-01-02", nextMaintenance)
	if err != nil {
		return err
	}

	asset := &entity.Assets{
		AssetName:       assetName,
		CategoryID:      categoryID,
		StatusID:        statusID,
		LastMaintenance: &parsedLastMaintenance,
		NextMaintenance: &parsedNextMaintenance,
	}
	return s.assetRepo.Create(asset)
}

func (s *assetService) GetAllAssets() ([]entity.Assets, error) {
	return s.assetRepo.FindAll()
}

func (s *assetService) GetAssetByID(assetID int) (*entity.Assets, error) {
	return s.assetRepo.FindByID(assetID)
}

func (s *assetService) UpdateAsset(assetID int, assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error {
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		return err
	}

	parsedLastMaintenance, err := time.Parse("2006-01-02", lastMaintenance)
	if err != nil {
		return err
	}

	parsedNextMaintenance, err := time.Parse("2006-01-02", nextMaintenance)
	if err != nil {
		return err
	}

	asset.AssetName = assetName
	asset.CategoryID = categoryID
	asset.StatusID = statusID
	asset.LastMaintenance = &parsedLastMaintenance
	asset.NextMaintenance = &parsedNextMaintenance

	return s.assetRepo.Update(asset)
}

func (s *assetService) DeleteAsset(assetID int) error {
	return s.assetRepo.Delete(assetID)
}
