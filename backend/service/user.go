package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model"
	"itl/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetGroups(id uuid.UUID) (*[]model.Group, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

type UserServiceParams struct {
	fx.In
	UserRepository repository.UserRepository
}

func (s *UserServiceImpl) GetGroups(id uuid.UUID) (*[]model.Group, error) {
	return s.userRepository.FindGroups(id)
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	return s.userRepository.Create(user)
}

func NewUserService(params UserServiceParams) UserService {
	return &UserServiceImpl{
		userRepository: params.UserRepository,
	}
}
