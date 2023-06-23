package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"itl/model"
	"itl/repository"
)

type PostService interface {
	CreatePost(Post *model.Post) (*model.Post, error)
	GetPost(id uuid.UUID) (*model.Post, error)
}

type PostServiceParams struct {
	fx.In
	PostRepository repository.PostRepository
}

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func (s *PostServiceImpl) GetPost(id uuid.UUID) (*model.Post, error) {
	return s.PostRepository.Get(id)
}

func (s *PostServiceImpl) CreatePost(post *model.Post) (*model.Post, error) {
	return s.PostRepository.Create(post)
}

func NewPostService(params PostServiceParams) PostService {
	return &PostServiceImpl{
		PostRepository: params.PostRepository,
	}
}
