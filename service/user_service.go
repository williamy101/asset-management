package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"go-asset-management/util"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface { //interface service user
	RegisterUser(name, email, password string, roleId int) error
	Login(email, password string) (string, error)
	GetUserByEmail(email string) (*entity.Users, error)
	GetUserByID(id int) (*entity.Users, error)
}

type userService struct { // struct user terhubung dengan repo user dan role
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *userService) RegisterUser(name, email, password string, roleId int) error {
	// validasi jika email sudah ada
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}

	// periksa adanya role
	role, err := s.roleRepo.FindByID(roleId)
	if err != nil || role == nil {
		return errors.New("invalid role ID")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &entity.Users{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		RoleID:   roleId,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}

func (s *userService) GetUserByEmail(email string) (*entity.Users, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *userService) GetUserByID(id int) (*entity.Users, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) Login(email, password string) (string, error) {
	// Cari data user berdasarkan email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	// Cek password yang dimasukkan dengan password yang ada di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token setelah berhasil login
	token, err := util.GenerateToken(user.UserID, user.RoleID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
