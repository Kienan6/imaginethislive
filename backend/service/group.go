package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	repository2 "itl/model/domain"
	"itl/repository"
)

type GroupService interface {
	CreateGroup(Group *repository2.Group) (*repository2.Group, error)
	GetGroup(id uuid.UUID) (*repository2.Group, error)
	FindByOwner(id uuid.UUID) (*[]repository2.Group, error)
	GetUsers(id uuid.UUID) (*[]repository2.User, error)
	GetPosts(id uuid.UUID) (*[]repository2.Post, error)
}

type GroupServiceImpl struct {
	GroupRepository repository.GroupRepository
}

func (s *GroupServiceImpl) GetUsers(id uuid.UUID) (*[]repository2.User, error) {
	return s.GroupRepository.FindUsers(id)
}

func (s *GroupServiceImpl) GetPosts(id uuid.UUID) (*[]repository2.Post, error) {
	return s.GroupRepository.FindPosts(id)
}

func (s *GroupServiceImpl) FindByOwner(id uuid.UUID) (*[]repository2.Group, error) {
	//TODO implement me
	return s.GroupRepository.FindByOwnership(id)
}

func (s *GroupServiceImpl) GetGroup(id uuid.UUID) (*repository2.Group, error) {
	return s.GroupRepository.Get(id)
}

type GroupServiceParams struct {
	fx.In
	GroupRepository repository.GroupRepository
}

func (s *GroupServiceImpl) CreateGroup(group *repository2.Group) (*repository2.Group, error) {
	return s.GroupRepository.Create(group)
}

func NewGroupService(params GroupServiceParams) GroupService {
	return &GroupServiceImpl{
		GroupRepository: params.GroupRepository,
	}
}
