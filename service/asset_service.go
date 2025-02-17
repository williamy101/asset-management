package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"go-asset-management/util"

	"gorm.io/gorm"
)

type AssetService interface {
	CreateAsset(assetName string, categoryID *int, statusID int, userID *int, lastMaintenance, nextMaintenance string) error
	GetAllAssets() ([]entity.Assets, error)
	GetAssetByID(assetID int) (*entity.Assets, error)
	UpdateAsset(assetID int, assetName string, categoryID *int, statusID int, userID *int, lastMaintenance, nextMaintenance string) error
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

func (s *assetService) CreateAsset(assetName string, categoryID *int, statusID int, userID *int, lastMaintenance, nextMaintenance string) error {
	if assetName == "" {
		return errors.New("asset name cannot be empty")
	}

	// validasi status
	validStatuses := map[int]bool{1: true, 2: true, 3: true}

	if !validStatuses[statusID] {
		return errors.New("invalid status ID")
	}

	if userID != nil && statusID != 2 {
		return errors.New("user can only be assigned when asset is 'In Use'")
	}

	// validasi kategori
	if categoryID != nil {
		_, err := s.assetCategoryRepo.FindByID(*categoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category does not exist")
			}
			return err
		}
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

func (s *assetService) UpdateAsset(assetID int, assetName string, categoryID *int, statusID int, userID *int, lastMaintenance, nextMaintenance string) error {
	// Check apakah aset ada
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset not found")
		}
		return err
	}

	if assetName != "" {
		asset.AssetName = assetName
	}

	if categoryID != nil {
		_, err = s.assetCategoryRepo.FindByID(*categoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category does not exist")
			}
			return err
		}
		asset.CategoryID = categoryID
	}

	validStatuses := map[int]bool{1: true, 2: true, 3: true, 13: true}

	if !validStatuses[statusID] {
		return errors.New("invalid status ID")
	}
	asset.StatusID = statusID

	if statusID == 13 {
		return errors.New("asset status cannot be manually set to 'Borrowed'")
	}

	if userID != nil {
		if statusID != 2 {
			return errors.New("user can only be assigned when asset is 'In Use'")
		}
		asset.UserID = userID
	} else if statusID == 1 {
		asset.UserID = nil
	}

	// Parse date
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
