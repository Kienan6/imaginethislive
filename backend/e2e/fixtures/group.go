package fixtures

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	repository2 "itl/model/domain"
	"itl/repository"
	"log"
)

type GroupFixture interface {
	GetDefaultUser(user *repository2.User) *repository2.User
	CreatePostInGroup(post *repository2.Post) *repository2.Post
	AddUserToGroup(userId uuid.UUID, groupId uuid.UUID)
}

type GroupFixtureParams struct {
	fx.In
	UserRepository  repository.UserRepository
	GroupRepository repository.GroupRepository
	PostRepository  repository.PostRepository
}

type GroupFixtureImpl struct {
	userRepository  repository.UserRepository
	groupRepository repository.GroupRepository
	postRepository  repository.PostRepository
}

func (g *GroupFixtureImpl) AddUserToGroup(userId uuid.UUID, groupId uuid.UUID) {
	err := g.userRepository.AddToGroup(userId, groupId)
	if err != nil {
		log.Fatal("error creating post")
	}
}

func (g *GroupFixtureImpl) CreatePostInGroup(post *repository2.Post) *repository2.Post {
	resp, err := g.postRepository.Create(post)
	if err != nil {
		log.Fatal("error creating post")
	}
	return resp
}

func (g *GroupFixtureImpl) GetDefaultUser(user *repository2.User) *repository2.User {
	resp, err := g.userRepository.Create(user)
	if err != nil {
		log.Fatal("error creating user")
	}
	return resp
}

func NewGroupFixture() GroupFixture {
	params := GroupFixtureParams{}
	NewSession(&params)
	return &GroupFixtureImpl{
		userRepository:  params.UserRepository,
		groupRepository: params.GroupRepository,
		postRepository:  params.PostRepository,
	}
}
