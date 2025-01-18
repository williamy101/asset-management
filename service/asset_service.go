package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"go-asset-management/util"

	"gorm.io/gorm"
)

type AssetService interface {
	CreateAsset(assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error
	GetAllAssets() ([]entity.Assets, error)
	GetAssetByID(assetID int) (*entity.Assets, error)
	UpdateAsset(assetID int, assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error
	DeleteAsset(assetID int) error
}

type assetService struct {
	assetRepo         repository.AssetRepository
	assetCategoryRepo repository.AssetCategoryRepository
	maintenanceRepo   repository.MaintenanceRepository
}

func NewAssetService(assetRepo repository.AssetRepository, assetCategoryRepo repository.AssetCategoryRepository, maintenanceRepo repository.MaintenanceRepository) AssetService {
	return &assetService{assetRepo: assetRepo, assetCategoryRepo: assetCategoryRepo, maintenanceRepo: maintenanceRepo}

}

func (s *assetService) CreateAsset(assetName string, categoryID int, statusID int, lastMaintenance, nextMaintenance string) error {
	if assetName == "" {
		return errors.New("asset name cannot be empty")
	}

	// validasi status
	validStatuses := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true} // cek status manual
	if !validStatuses[statusID] {
		return errors.New("invalid status ID")
	}

	// validasi kategori
	_, err := s.assetCategoryRepo.FindByID(categoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("category does not exist")
		}
		return err
	}

	asset := &entity.Assets{
		AssetName:       assetName,
		CategoryID:      categoryID,
		StatusID:        statusID,
		LastMaintenance: nil,
		NextMaintenance: nil,
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
	// Check if the asset exists
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset not found")
		}
		return err
	}

	// Update only if the field is provided
	if assetName != "" {
		asset.AssetName = assetName
	}

	if categoryID > 0 {
		// Validate categoryID exists
		_, err = s.assetCategoryRepo.FindByID(categoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category does not exist")
			}
			return err
		}
		asset.CategoryID = categoryID
	}

	if statusID > 0 {
		// Validate statusID
		validStatuses := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
		if !validStatuses[statusID] {
			return errors.New("invalid status ID")
		}
		asset.StatusID = statusID
	}

	// Parse and update maintenance dates only if provided
	if lastMaintenance != "" {
		parsedLastMaintenance, err := util.ParseDate(lastMaintenance)
		if err != nil {
			return errors.New("invalid last maintenance date")
		}
		asset.LastMaintenance = parsedLastMaintenance
	}

	if nextMaintenance != "" {
		parsedNextMaintenance, err := util.ParseDate(nextMaintenance)
		if err != nil {
			return errors.New("invalid next maintenance date")
		}
		asset.NextMaintenance = parsedNextMaintenance
	}

	// Save updated asset
	return s.assetRepo.Update(asset)
}

func (s *assetService) DeleteAsset(assetID int) error {
	_, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset not found")
		}
		return err
	}

	maintenances, err := s.maintenanceRepo.FindByAssetID(assetID)
	if err != nil {
		return err
	}
	if len(maintenances) > 0 {
		return errors.New("cannot delete asset with existing maintenances")
	}

	return s.assetRepo.Delete(assetID)
}
