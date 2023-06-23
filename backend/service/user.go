package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model"
	"itl/repository"
)

type UserService interface {
	CreateUser(*model.User) (*model.User, error)
	GetGroups(uuid.UUID) (*[]model.Group, error)
	AddToGroup(uuid.UUID, uuid.UUID) error
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

func (s *UserServiceImpl) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepository.Create(user)
}

func (s *UserServiceImpl) AddToGroup(u uuid.UUID, u2 uuid.UUID) error {
	return s.userRepository.AddToGroup(u, u2)
}

func NewUserService(params UserServiceParams) UserService {
	return &UserServiceImpl{
		userRepository: params.UserRepository,
	}
}
