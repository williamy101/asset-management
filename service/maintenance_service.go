package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
)

type MaintenanceService interface {
	CreateMaintenance(requestID int, worker int, description string, cost float64) error
	GetAllMaintenances() ([]entity.Maintenances, error)
	GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error)
	StartMaintenance(maintenanceID int) error
	EndMaintenance(maintenanceID int, description string) (*entity.Maintenances, error)
	DeleteMaintenance(maintenanceID int) error
	GetTotalCost() (float64, error)
	GetTotalCostByAssetID(assetID int) (map[string]interface{}, error)
	GetMaintenancesByWorkerID(workerID int) ([]entity.Maintenances, error)
}

type maintenanceService struct {
	maintenanceRepo        repository.MaintenanceRepository
	assetRepo              repository.AssetRepository
	maintenanceRequestRepo repository.MaintenanceRequestRepository
}

func NewMaintenanceService(
	maintenanceRepo repository.MaintenanceRepository,
	assetRepo repository.AssetRepository,
	maintenanceRequestRepo repository.MaintenanceRequestRepository,
) MaintenanceService {
	return &maintenanceService{
		maintenanceRepo:        maintenanceRepo,
		assetRepo:              assetRepo,
		maintenanceRequestRepo: maintenanceRequestRepo,
	}
}

func (s *maintenanceService) CreateMaintenance(assetID int, worker int, description string, cost float64) error {
	if cost < 0 {
		return errors.New("cost cannot be negative")
	}

	// Cek apakah aset ada dan tidak dalam status "Data Deleted"
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		return errors.New("asset not found")
	}

	// Cek apakah aset sudah dalam perbaikan
	activeMaintenance, _ := s.maintenanceRepo.FindActiveByAssetID(asset.AssetID)
	if activeMaintenance != nil {
		return errors.New("this asset is already under maintenance, complete it first before scheduling a new one")
	}

	// Buat maintenance baru tanpa melalui maintenance request
	maintenance := &entity.Maintenances{
		AssetID:     assetID,
		Worker:      worker,
		Description: description,
		Cost:        cost,
		StatusID:    4, // Scheduled
	}

	err = s.maintenanceRepo.Create(maintenance)
	if err != nil {
		return err
	}

	asset.StatusID = 3
	return s.assetRepo.Update(asset)
}

func (s *maintenanceService) GetAllMaintenances() ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindAll()
}

func (s *maintenanceService) GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error) {
	return s.maintenanceRepo.FindByID(maintenanceID)
}

func (s *maintenanceService) StartMaintenance(maintenanceID int) error {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return err
	}

	if maintenance.StatusID != 4 {
		return errors.New("Only scheduled maintenance can be started")
	}

	maintenance.StatusID = 5 // "In Progress"
	err = s.maintenanceRepo.Update(maintenance)
	if err != nil {
		return err
	}

	return nil
}

func (s *maintenanceService) EndMaintenance(maintenanceID int, description string) (*entity.Maintenances, error) {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return nil, err
	}

	if maintenance.StatusID != 5 { // In Progress
		return nil, errors.New("cannot complete maintenance that is not in progress")
	}

	maintenance.StatusID = 6 // Completed
	maintenance.Description = description
	err = s.maintenanceRepo.Update(maintenance)

	if err != nil {
		return nil, err
	}

	return maintenance, nil
}

func (s *maintenanceService) DeleteMaintenance(maintenanceID int) error {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return err
	}

	if maintenance.StatusID == 6 {
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

	asset.StatusID = 1
	return s.assetRepo.Update(asset)
}

func (s *maintenanceService) GetTotalCost() (float64, error) {
	return s.maintenanceRepo.CalculateTotalCost()
}

func (s *maintenanceService) GetTotalCostByAssetID(assetID int) (map[string]interface{}, error) {
	return s.maintenanceRepo.GetTotalCostByAssetID(assetID)
}

func (s *maintenanceService) GetMaintenancesByWorkerID(workerID int) ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindByUserID(workerID)
}
