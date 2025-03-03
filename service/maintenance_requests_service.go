package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"

	"gorm.io/gorm"
)

type MaintenanceRequestService interface {
	CreateMaintenanceRequest(assetID int, userID int, issueDescription string) error
	GetAllMaintenanceRequests() ([]entity.MaintenanceRequests, error)
	GetMaintenanceRequestByID(requestID int) (*entity.MaintenanceRequests, error)
	GetMaintenanceRequestsByUserID(userID int) ([]entity.MaintenanceRequests, error)
	GetMaintenanceRequestsByAssetID(assetID int) ([]entity.MaintenanceRequests, error)
	GetMaintenanceRequestsByStatus(statusID int) ([]entity.MaintenanceRequests, error)
	ApproveMaintenanceRequest(
		requestID int,
		worker int,
		description string,
		cost float64,
		maintenanceDate time.Time,
	) error
	RejectMaintenanceRequest(requestID int) error
	DeleteMaintenanceRequest(requestID int) error
}

type maintenanceRequestService struct {
	maintenanceRequestRepo repository.MaintenanceRequestRepository
	assetRepo              repository.AssetRepository
	maintenanceRepo        repository.MaintenanceRepository
	userRepo               repository.UserRepository
}

func NewMaintenanceRequestService(
	maintenanceRequestRepo repository.MaintenanceRequestRepository,
	assetRepo repository.AssetRepository,
	maintenanceRepo repository.MaintenanceRepository,
	userRepo repository.UserRepository,
) MaintenanceRequestService {
	return &maintenanceRequestService{
		maintenanceRequestRepo: maintenanceRequestRepo,
		assetRepo:              assetRepo,
		maintenanceRepo:        maintenanceRepo,
		userRepo:               userRepo,
	}
}

func (s *maintenanceRequestService) CreateMaintenanceRequest(assetID int, userID int, issueDescription string) error {
	_, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset does not exist")
		}
		return err
	}

	request := &entity.MaintenanceRequests{
		AssetID:          assetID,
		UserID:           userID,
		RequestDate:      time.Now(),
		IssueDescription: issueDescription,
		StatusID:         7, // Status "Pending Approval"
	}

	return s.maintenanceRequestRepo.Create(request)
}

func (s *maintenanceRequestService) GetAllMaintenanceRequests() ([]entity.MaintenanceRequests, error) {
	return s.maintenanceRequestRepo.FindAll()
}

func (s *maintenanceRequestService) GetMaintenanceRequestByID(requestID int) (*entity.MaintenanceRequests, error) {
	return s.maintenanceRequestRepo.FindByID(requestID)
}

func (s *maintenanceRequestService) GetMaintenanceRequestsByUserID(userID int) ([]entity.MaintenanceRequests, error) {
	return s.maintenanceRequestRepo.FindByUserID(userID)
}

func (s *maintenanceRequestService) GetMaintenanceRequestsByAssetID(assetID int) ([]entity.MaintenanceRequests, error) {
	return s.maintenanceRequestRepo.FindByAssetID(assetID)
}

func (s *maintenanceRequestService) GetMaintenanceRequestsByStatus(statusID int) ([]entity.MaintenanceRequests, error) {
	return s.maintenanceRequestRepo.FindByStatusID(statusID)
}

func (s *maintenanceRequestService) ApproveMaintenanceRequest(
	requestID int,
	worker int,
	description string,
	cost float64,
	maintenanceDate time.Time,
) error {
	request, err := s.maintenanceRequestRepo.FindByID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("maintenance request not found")
		}
		return err
	}

	if request.StatusID != 7 {
		return errors.New("request cannot be approved as it is not pending")
	}

	asset, err := s.assetRepo.FindByID(request.AssetID)
	if err != nil {
		return errors.New("asset not found")
	}

	user, err := s.userRepo.FindByID(worker)
	if err != nil {
		return errors.New("assigned worker not found")
	}

	if user.RoleID != 1 && user.RoleID != 2 {
		return errors.New("only Admins or Technicians can perform maintenance")
	}

	activeMaintenance, err := s.maintenanceRepo.FindActiveByAssetID(asset.AssetID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if activeMaintenance != nil {
		return errors.New("this asset is already under maintenance, complete it first before scheduling a new one")
	}

	now := time.Now()
	request.StatusID = 8 // Status "Approved"
	request.DecisionDate = &now
	request.MaintenanceDate = &maintenanceDate

	if err := s.maintenanceRequestRepo.Update(request); err != nil {
		return err
	}

	maintenance := &entity.Maintenances{
		AssetID:     request.AssetID,
		Worker:      worker,
		Description: description,
		Cost:        cost,
		StatusID:    4, // Status "Scheduled"
	}

	if err := s.maintenanceRepo.Create(maintenance); err != nil {
		return err
	}

	asset.StatusID = 3
	if err := s.assetRepo.Update(asset); err != nil {
		return err
	}

	request.StatusID = 4
	return s.maintenanceRequestRepo.Update(request)
}

func (s *maintenanceRequestService) RejectMaintenanceRequest(requestID int) error {
	request, err := s.maintenanceRequestRepo.FindByID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("maintenance request not found")
		}
		return err
	}

	if request.StatusID != 7 {
		return errors.New("request cannot be rejected as it is not pending")
	}

	now := time.Now()
	request.StatusID = 9
	request.DecisionDate = &now

	return s.maintenanceRequestRepo.Update(request)
}

func (s *maintenanceRequestService) DeleteMaintenanceRequest(requestID int) error {
	request, err := s.maintenanceRequestRepo.FindByID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("maintenance request not found")
		}
		return err
	}

	if request.StatusID != 7 {
		return errors.New("cannot delete a processed maintenance request")
	}

	return s.maintenanceRequestRepo.Delete(requestID)
}
