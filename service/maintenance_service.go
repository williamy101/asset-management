package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"
)

type MaintenanceService interface {
	CreateMaintenance(requestID int, worker int, description string, cost float64) error
	GetAllMaintenances() ([]entity.Maintenances, error)
	GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error)
	UpdateMaintenance(maintenanceID int, description string, statusID int) error
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

// **Create Maintenance**
func (s *maintenanceService) CreateMaintenance(requestID int, worker int, description string, cost float64) error {
	if cost < 0 {
		return errors.New("cost cannot be negative")
	}

	// Fetch the maintenance request
	request, err := s.maintenanceRequestRepo.FindByID(requestID)
	if err != nil {
		return errors.New("maintenance request not found")
	}

	// Ensure the request has been approved
	if request.StatusID != 8 { // Status 8 = Approved
		return errors.New("maintenance cannot be created from an unapproved request")
	}

	// Fetch the asset associated with the request
	asset, err := s.assetRepo.FindByID(request.AssetID)
	if err != nil {
		return errors.New("asset not found")
	}

	// **Check if the asset already has an active maintenance**
	activeMaintenance, _ := s.maintenanceRepo.FindActiveByAssetID(asset.AssetID)
	if activeMaintenance != nil {
		return errors.New("this asset is already under maintenance, complete it first before scheduling a new one")
	}

	// Proceed to create a new maintenance
	maintenance := &entity.Maintenances{
		AssetID:     request.AssetID,
		Worker:      worker,
		Description: description,
		Cost:        cost,
		StatusID:    3, // Status 3 = In Maintenance
	}

	err = s.maintenanceRepo.Create(maintenance)
	if err != nil {
		return err
	}

	// Update asset status to "In Maintenance"
	asset.StatusID = 3
	err = s.assetRepo.Update(asset)
	if err != nil {
		return err
	}

	// Update the maintenance request status to "In Progress"
	request.StatusID = 4 // Status 4 = In Progress
	err = s.maintenanceRequestRepo.Update(request)
	if err != nil {
		return err
	}

	return nil
}

// **Get All Maintenances**
func (s *maintenanceService) GetAllMaintenances() ([]entity.Maintenances, error) {
	return s.maintenanceRepo.FindAll()
}

// **Get Maintenance By ID**
func (s *maintenanceService) GetMaintenanceByID(maintenanceID int) (*entity.Maintenances, error) {
	return s.maintenanceRepo.FindByID(maintenanceID)
}

// **Update Maintenance**
func (s *maintenanceService) UpdateMaintenance(maintenanceID int, description string, statusID int) error {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
		return err
	}

	// Check if status is valid
	if statusID != 3 && statusID != 4 && statusID != 5 {
		return errors.New("invalid status ID, must be 3 (In Maintenance), 4 (Scheduled), or 5 (Completed)")
	}

	maintenance.Description = description
	maintenance.StatusID = statusID

	err = s.maintenanceRepo.Update(maintenance)
	if err != nil {
		return err
	}

	asset, err := s.assetRepo.FindByID(maintenance.AssetID)
	if err != nil {
		return err
	}

	if statusID == 5 {
		asset.StatusID = 1
		now := time.Now()
		asset.LastMaintenance = &now
		asset.NextMaintenance = nil

		request, err := s.maintenanceRequestRepo.FindByAssetID(maintenance.AssetID)
		if err == nil {
			for i := range request {
				if request[i].StatusID == 8 {
					request[i].StatusID = 5
					request[i].MaintenanceDate = &now
					_ = s.maintenanceRequestRepo.Update(&request[i])
				}
			}
		}
	}

	return s.assetRepo.Update(asset)
}

// **Delete Maintenance**
func (s *maintenanceService) DeleteMaintenance(maintenanceID int) error {
	maintenance, err := s.maintenanceRepo.FindByID(maintenanceID)
	if err != nil {
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
