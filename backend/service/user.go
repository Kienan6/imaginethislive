package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	repository2 "itl/model/domain"
	"itl/repository"
)

type UserService interface {
	CreateUser(*repository2.User) (*repository2.User, error)
	GetGroups(uuid.UUID) (*[]repository2.Group, error)
	AddToGroup(uuid.UUID, uuid.UUID) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

type UserServiceParams struct {
	fx.In
	UserRepository repository.UserRepository
}

func (s *UserServiceImpl) GetGroups(id uuid.UUID) (*[]repository2.Group, error) {
	return s.userRepository.FindGroups(id)
}

func (s *UserServiceImpl) CreateUser(user *repository2.User) (*repository2.User, error) {
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
