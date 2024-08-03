package service

import (
	"CRUD-SQL/model"
	"CRUD-SQL/repository"
	"errors"
)

// Interface
type UserService interface {
	RegisterUser(*model.UserRegister) error
	LoginUser(string, string) (*model.UserRegister, error)
	GetUser(string) (*model.UserRegister, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// Concrete-Code

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(u *model.UserRegister) error {
	return s.userRepo.CreateUser(u)
}

func (s *userService) GetUser(email string) (*model.UserRegister, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) LoginUser(email string, password string) (*model.UserRegister, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
