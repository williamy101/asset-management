package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"

	"gorm.io/gorm"
)

type AssetService interface {
	CreateAsset(assetName string, categoryID *int, statusID int, userID *int) error
	GetAllAssets() ([]entity.Assets, error)
	GetAssetByID(assetID int) (*entity.Assets, error)
	UpdateAsset(assetID int, assetName *string, categoryID *int, statusID *int, userID *int) error
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

func (s *assetService) CreateAsset(assetName string, categoryID *int, statusID int, userID *int) error {
	if assetName == "" {
		return errors.New("asset name cannot be empty")
	}

	// validasi status
	validStatuses := map[int]bool{1: true, 2: true, 3: true}
	if !validStatuses[statusID] {
		return errors.New("invalid status ID")
	}

	// userID hanya boleh diisi jika status adalah "In Use"
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
		AssetName:  assetName,
		CategoryID: categoryID,
		StatusID:   statusID,
		UserID:     nil, // default userID ke NULL
	}

	return s.assetRepo.Create(asset)
}

func (s *assetService) GetAllAssets() ([]entity.Assets, error) {
	return s.assetRepo.FindAll()
}

func (s *assetService) GetAssetByID(assetID int) (*entity.Assets, error) {
	return s.assetRepo.FindByID(assetID)
}

func (s *assetService) UpdateAsset(assetID int, assetName *string, categoryID *int, statusID *int, userID *int) error {
	// Check apakah aset ada
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset not found")
		}
		return err
	}

	// Update hanya jika field dikirim
	if assetName != nil {
		asset.AssetName = *assetName
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

	if statusID != nil {
		validStatuses := map[int]bool{1: true, 2: true, 3: true}
		if !validStatuses[*statusID] {
			return errors.New("invalid status ID")
		}
		asset.StatusID = *statusID

		// userID hanya bisa diisi jika status adalah "In Use"
		if *statusID == 2 && userID == nil {
			return errors.New("user ID is required when asset is 'In Use'")
		}
	}

	if userID != nil {
		asset.UserID = userID
	} else if statusID != nil && *statusID == 1 {
		asset.UserID = nil // Hapus user jika status jadi "Available"
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
