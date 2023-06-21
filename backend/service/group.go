package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model"
	"itl/repository"
)

type GroupService interface {
	CreateGroup(Group *model.Group) error
	GetGroup(id uuid.UUID) (*model.Group, error)
	FindByOwner(id uuid.UUID) (*[]model.Group, error)
	GetUsers(id uuid.UUID) (*[]model.User, error)
}

type GroupServiceImpl struct {
	GroupRepository repository.GroupRepository
}

func (s *GroupServiceImpl) GetUsers(id uuid.UUID) (*[]model.User, error) {
	return s.GroupRepository.FindUsers(id)
}

func (s *GroupServiceImpl) FindByOwner(id uuid.UUID) (*[]model.Group, error) {
	//TODO implement me
	return s.GroupRepository.FindByOwnership(id)
}

func (s *GroupServiceImpl) GetGroup(id uuid.UUID) (*model.Group, error) {
	return s.GroupRepository.Get(id)
}

type GroupServiceParams struct {
	fx.In
	GroupRepository repository.GroupRepository
}

func (s *GroupServiceImpl) CreateGroup(Group *model.Group) error {
	return s.GroupRepository.Create(Group)
}

func NewGroupService(params GroupServiceParams) GroupService {
	return &GroupServiceImpl{
		GroupRepository: params.GroupRepository,
	}
}
