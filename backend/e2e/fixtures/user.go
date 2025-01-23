package fixtures

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	repository2 "itl/model/domain"
	"itl/repository"
	"log"
)

type UserFixture interface {
	CreateAndAddUserToGroup(id uuid.UUID, group *repository2.Group)
	CreateGroup(group *repository2.Group) uuid.UUID
}

type UserFixtureImpl struct {
	userRepository  repository.UserRepository
	groupRepository repository.GroupRepository
}

func (fixture *UserFixtureImpl) CreateGroup(group *repository2.Group) uuid.UUID {
	resp, err := fixture.groupRepository.Create(group)
	if err != nil {
		log.Fatal("Could not create group")
	}
	return resp.ID
}

type UserFixtureParams struct {
	fx.In
	UserRepository  repository.UserRepository
	GroupRepository repository.GroupRepository
}

func (fixture *UserFixtureImpl) CreateAndAddUserToGroup(id uuid.UUID, group *repository2.Group) {
	resp, err := fixture.groupRepository.Create(group)
	if err != nil {
		log.Fatal("Could not create group")
	}

	err = fixture.userRepository.AddToGroup(id, resp.ID)
	if err != nil {
		log.Fatal("Could not create group")
	}
}

func NewUserFixture() UserFixture {
	params := UserFixtureParams{}
	NewSession(&params)
	return &UserFixtureImpl{
		userRepository:  params.UserRepository,
		groupRepository: params.GroupRepository,
	}
}
