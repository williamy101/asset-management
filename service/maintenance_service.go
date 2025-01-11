package service

import (
	"go-asset-management/entity"
	"go-asset-management/repository"
)

type MaintenanceService interface {
	CreateMaintenance(assetID int, userID int, description string, cost float64, statusID int) error
	GetAllMaintenances() ([]entity.Maintenances, error)
	// GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error)
	// UpdateMaintenance(maintenanceID int, description string, cost float64, statusID int) error
	// DeleteMaintenance(maintenanceID int) error
}

type maintenanceService struct {
	maintenanceRepo repository.MaintenanceRepository
	assetRepo       repository.AssetRepository
}

func NewMaintenanceService(maintenanceRepo repository.MaintenanceRepository, assetRepo repository.AssetRepository) MaintenanceService {
	return &maintenanceService{
		maintenanceRepo: maintenanceRepo,
		assetRepo:       assetRepo,
	}
}

func (s *maintenanceService) CreateMaintenance(assetID int, userID int, description string, cost float64, statusID int) error {

	maintenance := &entity.Maintenances{
		AssetID:     assetID,
		UserID:      userID,
		Description: description,
		Cost:        cost,
		StatusID:    4, // StatusID 4 adalah "Scheduled"
	}

	err := s.maintenanceRepo.Create(maintenance)
	if err != nil {
		return err
	}

	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		return err
	}

	asset.StatusID = 3 // Mengubah status asset yang terjadwal maintenance jadi "In maintenance"
	err = s.assetRepo.Update(asset)
	if err != nil {
		return err
	}

	return nil
}

func (s *maintenanceService) GetAllMaintenances() ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindAll()
}
