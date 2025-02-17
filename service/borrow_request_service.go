package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"

	"gorm.io/gorm"
)

type BorrowAssetRequestService interface {
	CreateBorrowRequest(assetID int, userID int, requestedStartDate, requestedEndDate time.Time) error
	GetAllBorrowRequests() ([]entity.BorrowAssetRequests, error)
	GetBorrowRequestByID(requestID int) (*entity.BorrowAssetRequests, error)
	GetBorrowRequestsByUserID(userID int) ([]entity.BorrowAssetRequests, error)
	GetBorrowRequestsByAssetID(assetID int) ([]entity.BorrowAssetRequests, error)
	GetBorrowRequestsByStatus(statusID int) ([]entity.BorrowAssetRequests, error)
	ApproveBorrowRequest(requestID int, approvedBy int) error
	RejectBorrowRequest(requestID int) error
	DeleteBorrowRequest(requestID int) error
}

type borrowAssetRequestService struct {
	borrowAssetRequestRepo repository.BorrowAssetRequestRepository
	borrowedAssetRepo      repository.BorrowedAssetRepository
	assetRepo              repository.AssetRepository
}

func NewBorrowAssetRequestService(
	borrowAssetRequestRepo repository.BorrowAssetRequestRepository,
	borrowedAssetRepo repository.BorrowedAssetRepository,
	assetRepo repository.AssetRepository,
) BorrowAssetRequestService {
	return &borrowAssetRequestService{
		borrowAssetRequestRepo: borrowAssetRequestRepo,
		borrowedAssetRepo:      borrowedAssetRepo,
		assetRepo:              assetRepo,
	}
}

func (s *borrowAssetRequestService) CreateBorrowRequest(assetID int, userID int, requestedStartDate, requestedEndDate time.Time) error {
	asset, err := s.assetRepo.FindByID(assetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("asset does not exist")
		}
		return err
	}

	if asset.StatusID != 1 { // 1 = Available
		return errors.New("asset is not available for borrowing")
	}

	request := &entity.BorrowAssetRequests{
		AssetID:            assetID,
		UserID:             userID,
		RequestDate:        time.Now(),
		RequestedStartDate: requestedStartDate,
		RequestedEndDate:   requestedEndDate,
		StatusID:           10, // 10 = Pending Approval
	}

	return s.borrowAssetRequestRepo.Create(request)
}

func (s *borrowAssetRequestService) GetAllBorrowRequests() ([]entity.BorrowAssetRequests, error) {
	return s.borrowAssetRequestRepo.FindAll()
}

func (s *borrowAssetRequestService) GetBorrowRequestByID(requestID int) (*entity.BorrowAssetRequests, error) {
	return s.borrowAssetRequestRepo.FindByID(requestID)
}

func (s *borrowAssetRequestService) GetBorrowRequestsByUserID(userID int) ([]entity.BorrowAssetRequests, error) {
	return s.borrowAssetRequestRepo.FindByUserID(userID)
}

func (s *borrowAssetRequestService) GetBorrowRequestsByAssetID(assetID int) ([]entity.BorrowAssetRequests, error) {
	return s.borrowAssetRequestRepo.FindByAssetID(assetID)
}

func (s *borrowAssetRequestService) GetBorrowRequestsByStatus(statusID int) ([]entity.BorrowAssetRequests, error) {
	return s.borrowAssetRequestRepo.FindByStatus(statusID)
}

func (s *borrowAssetRequestService) ApproveBorrowRequest(requestID int, approvedBy int) error {
	request, err := s.borrowAssetRequestRepo.FindByID(requestID)
	if err != nil {
		return errors.New("borrow request not found")
	}

	if request.StatusID != 10 { // 10 = Pending Approval
		return errors.New("request cannot be approved as it is not pending")
	}

	asset, err := s.assetRepo.FindByID(request.AssetID)
	if err != nil {
		return errors.New("asset not found")
	}

	if asset.StatusID != 1 { // 1 = Available
		return errors.New("asset is not available for borrowing")
	}

	now := time.Now()
	borrowedAsset := &entity.BorrowedAssets{
		AssetID:    request.AssetID,
		UserID:     request.UserID,
		BorrowDate: now,
		StatusID:   13, // 13 = Borrowed
	}

	if err := s.borrowedAssetRepo.Create(borrowedAsset); err != nil {
		return err
	}

	asset.StatusID = 2 // 2 = In Use
	if err := s.assetRepo.Update(asset); err != nil {
		return err
	}

	request.StatusID = 11 // 11 = Approved
	request.ApprovedBy = &approvedBy
	if err := s.borrowAssetRequestRepo.Update(request); err != nil {
		return err
	}

	return nil
}

func (s *borrowAssetRequestService) RejectBorrowRequest(requestID int) error {
	request, err := s.borrowAssetRequestRepo.FindByID(requestID)
	if err != nil {
		return errors.New("borrow request not found")
	}

	if request.StatusID != 10 { // 10 = Pending Approval
		return errors.New("request cannot be rejected as it is not pending")
	}

	now := time.Now()
	request.StatusID = 12 // 12 = Rejected
	request.UpdatedAt = now

	return s.borrowAssetRequestRepo.Update(request)
}

func (s *borrowAssetRequestService) DeleteBorrowRequest(requestID int) error {
	request, err := s.borrowAssetRequestRepo.FindByID(requestID)
	if err != nil {
		return errors.New("borrow request not found")
	}

	if request.StatusID != 10 {
		return errors.New("cannot delete a processed borrow request")
	}

	return s.borrowAssetRequestRepo.Delete(requestID)
}
