package service

import (
	"go-asset-management/entity"
	"go-asset-management/repository"
)

type StatusService interface {
	Create(statusName string) error
	GetAll() ([]entity.Statuses, error)
	GetByID(id int) (*entity.Statuses, error)
	Update(id int, statusName string) error
	Delete(id int) error
}

type statusService struct {
	repo repository.StatusRepository
}

func NewStatusService(repo repository.StatusRepository) StatusService {
	return &statusService{repo: repo}
}

func (s *statusService) Create(statusName string) error {
	status := &entity.Statuses{
		StatusName: statusName,
	}
	return s.repo.Create(status)
}

func (s *statusService) GetAll() ([]entity.Statuses, error) {
	return s.repo.FindAll()
}

func (s *statusService) GetByID(id int) (*entity.Statuses, error) {
	return s.repo.FindByID(id)
}

func (s *statusService) Update(id int, statusName string) error {
	status, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	status.StatusName = statusName
	return s.repo.Update(status)
}

func (s *statusService) Delete(id int) error {
	return s.repo.Delete(id)
}
