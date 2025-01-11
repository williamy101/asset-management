package service

import (
	"go-asset-management/entity"
	"go-asset-management/repository"
)

type RoleService interface {
	GetRoleByID(roleID int) (*entity.Roles, error)
	GetAllRoles() ([]entity.Roles, error)
	CreateRole(roleName string) error
	DeleteRole(roleID int) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func (s *roleService) GetRoleByID(roleID int) (*entity.Roles, error) {
	return s.roleRepo.FindByID(roleID)
}

func (s *roleService) GetAllRoles() ([]entity.Roles, error) {
	return s.roleRepo.FindAll()
}

func (s *roleService) CreateRole(roleName string) error {
	role := &entity.Roles{
		RoleName: roleName,
	}
	return s.roleRepo.Create(role)
}

func (s *roleService) DeleteRole(roleID int) error {
	return s.roleRepo.DeleteByID(roleID)
}
