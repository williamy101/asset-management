package service

import (
	"errors"
	"go-asset-management/entity"
	"go-asset-management/repository"
	"go-asset-management/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface { //interface service user
	RegisterUser(name, email, password string) error
	Login(email, password string) (string, error)
	GetUserByEmail(email string) (*entity.Users, error)
	GetUserByID(id int) (*entity.UserDTO, error)
	GetAllUsers() ([]entity.UserDTO, error)
	UpdateUserRole(userID int, newRoleID int) error
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

func (s *userService) RegisterUser(name, email, password string) error {

	// validasi jika email sudah ada
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &entity.Users{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		RoleID:   2,
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

func (s *userService) GetUserByID(id int) (*entity.UserDTO, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Mapping ke UserDTO
	userDTO := &entity.UserDTO{
		UserID: user.UserID,
		Name:   user.Name,
		RoleID: user.RoleID,
	}

	return userDTO, nil
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

func (s *userService) GetAllUsers() ([]entity.UserDTO, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Mapping ke UserDTO
	var userDTOs []entity.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, entity.UserDTO{
			UserID: user.UserID,
			Name:   user.Name,
			RoleID: user.RoleID,
		})
	}

	return userDTOs, nil
}

func (s *userService) UpdateUserRole(userID int, newRoleID int) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	role, err := s.roleRepo.FindByID(newRoleID)
	if err != nil || role == nil {
		return errors.New("invalid role ID")
	}

	user.RoleID = newRoleID
	return s.userRepo.Update(user)
}
