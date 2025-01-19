package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"

	"gorm.io/gorm"
)

type MaintenanceService interface {
	CreateMaintenance(assetID int, userID int, description string, cost float64) error
	GetAllMaintenances() ([]entity.Maintenances, error)
	GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error)
	UpdateMaintenance(maintenanceID int, description string, statusID int) error
	DeleteMaintenance(maintenanceID int) error
	GetTotalCost() (float64, error)
	GetTotalCostByAssetID(assetID int) (map[string]interface{}, error)
	GetMaintenancesByUserID(userID int) ([]entity.Maintenances, error)
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

func (s *maintenanceService) CreateMaintenance(assetID int, userID int, description string, cost float64) error {

	if cost < 0 {
		return errors.New("cost cannot be negative")
	}

	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset does not exist")
		}
		return err
	}

	// cek ketersediaan aset untuk di-maintenance
	if asset.StatusID != 1 { // 1 = Available
		return errors.New("maintenance cannot be created because the asset is not available")
	}

	maintenance := &entity.Maintenances{
		AssetID:     assetID,
		UserID:      userID,
		Description: description,
		Cost:        cost,
		StatusID:    4, // StatusID 4 adalah "Scheduled" untuk default
	}

	err = s.maintenanceRepo.Create(maintenance)
	if err != nil {
		return err
	}

	asset.StatusID = 4
	asset.NextMaintenance = &maintenance.CreatedAt
	err = s.assetRepo.Update(asset)
	if err != nil {
		return err
	}

	return nil
}

func (s *maintenanceService) GetAllMaintenances() ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindAll()
}

func (s *maintenanceService) GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error) {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return nil, err
	}
	return maintenance, nil
}

func (s *maintenanceService) UpdateMaintenance(maintenanceID int, description string, statusID int) error {
	if statusID != 3 && statusID != 4 && statusID != 5 {
		return errors.New("Invalid status ID. Only 3 ('In Maintenance'), 4 ('Scheduled'), or 5 ('Completed') are allowed for maintenance status")
	}

	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return err
	}

	maintenance.Description = description
	maintenance.StatusID = statusID

	err = s.maintenanceRepo.Update(maintenance)
	if err != nil {
		return err
	}

	if statusID == 3 {
		asset, err := s.assetRepo.FindByID(maintenance.AssetID)
		if err != nil {
			return err
		}

		asset.StatusID = 3
		err = s.assetRepo.Update(asset)
		if err != nil {
			return err
		}
	}

	if statusID == 5 {
		asset, err := s.assetRepo.FindByID(maintenance.AssetID)
		if err != nil {
			return err
		}
		asset.StatusID = 1
		err = s.assetRepo.Update(asset)
		if err != nil {
			return err
		}
		now := time.Now()
		asset.LastMaintenance = &now
		asset.NextMaintenance = nil // Clear NextMaintenance saat selesai
		err = s.assetRepo.Update(asset)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *maintenanceService) DeleteMaintenance(maintenanceID int) error {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("maintenance not found")
		}
		return err
	}

	if maintenance.StatusID == 5 {
		return errors.New("cannot delete completed maintenance")
	}

	err = s.maintenanceRepo.Delete(maintenanceID)
	if err != nil {
		return err
	}

	asset, err := s.assetRepo.FindByID(maintenance.AssetID)
	if err != nil {
		return err
	}

	asset.StatusID = 1 // Status aset dikembalikan ke 'Available'
	err = s.assetRepo.Update(asset)
	if err != nil {
		return err
	}

	return nil
}

func (s *maintenanceService) GetTotalCost() (float64, error) {
	return s.maintenanceRepo.CalculateTotalCost()
}

func (s *maintenanceService) GetTotalCostByAssetID(assetID int) (map[string]interface{}, error) {
	return s.maintenanceRepo.GetTotalCostByAssetID(assetID)
}

func (s *maintenanceService) GetMaintenancesByUserID(userID int) ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindByUserID(userID)
}
