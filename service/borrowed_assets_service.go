package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"time"
)

type BorrowedAssetService interface {
	GetAllBorrowedAssets() ([]entity.BorrowedAssets, error)
	GetBorrowedAssetByID(borrowID int) (*entity.BorrowedAssets, error)
	GetBorrowedAssetsByUserID(userID int) ([]entity.BorrowedAssets, error)
	GetBorrowedAssetsByAssetID(assetID int) ([]entity.BorrowedAssets, error)
	GetBorrowedAssetsByStatus(statusID int) ([]entity.BorrowedAssets, error)
	UpdateReturnDate(borrowID int, returnDate time.Time) error
}

type borrowedAssetService struct {
	borrowedAssetRepo repository.BorrowedAssetRepository
	borrowRequestRepo repository.BorrowAssetRequestRepository
	assetRepo         repository.AssetRepository
}

func NewBorrowedAssetService(
	borrowedAssetRepo repository.BorrowedAssetRepository,
	borrowRequestRepo repository.BorrowAssetRequestRepository,
	assetRepo repository.AssetRepository,

) BorrowedAssetService {
	return &borrowedAssetService{
		borrowedAssetRepo: borrowedAssetRepo,
		borrowRequestRepo: borrowRequestRepo,
		assetRepo:         assetRepo,
	}
}

func (s *borrowedAssetService) GetAllBorrowedAssets() ([]entity.BorrowedAssets, error) {
	return s.borrowedAssetRepo.FindAll()
}

func (s *borrowedAssetService) GetBorrowedAssetByID(borrowID int) (*entity.BorrowedAssets, error) {
	return s.borrowedAssetRepo.FindByID(borrowID)
}

func (s *borrowedAssetService) GetBorrowedAssetsByUserID(userID int) ([]entity.BorrowedAssets, error) {
	return s.borrowedAssetRepo.FindByUserID(userID)
}

func (s *borrowedAssetService) GetBorrowedAssetsByAssetID(assetID int) ([]entity.BorrowedAssets, error) {
	return s.borrowedAssetRepo.FindByAssetID(assetID)
}

func (s *borrowedAssetService) GetBorrowedAssetsByStatus(statusID int) ([]entity.BorrowedAssets, error) {
	return s.borrowedAssetRepo.FindByStatus(statusID)
}

func (s *borrowedAssetService) UpdateReturnDate(borrowID int, returnDate time.Time) error {
	borrowed, err := s.borrowedAssetRepo.FindByID(borrowID)
	if err != nil {
		return errors.New("borrowed asset not found")
	}

	if borrowed.StatusID != 13 { // 13 = Borrowed
		return errors.New("asset is not currently borrowed")
	}

	request, err := s.borrowRequestRepo.FindByID(borrowed.BorrowRequestID)
	if err != nil {
		return errors.New("associated borrow request not found")
	}

	if returnDate.After(request.RequestedEndDate) {
		borrowed.StatusID = 15 // 15 = Overdue
	} else {
		borrowed.StatusID = 14 // 14 = Returned
	}

	borrowed.ReturnDate = &returnDate

	if err := s.borrowedAssetRepo.Update(borrowed); err != nil {
		return err
	}

	asset, err := s.assetRepo.FindByID(borrowed.AssetID)
	if err != nil {
		return errors.New("asset not found")
	}

	asset.StatusID = 1 // 1 = Available
	asset.UserID = nil // cabut user
	return s.assetRepo.Update(asset)
}
