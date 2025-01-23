package service

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	repository2 "itl/model/domain"
	"itl/repository"
)

type PostService interface {
	CreatePost(Post *repository2.Post) (*repository2.Post, error)
	GetPost(id uuid.UUID) (*repository2.Post, error)
}

type PostServiceParams struct {
	fx.In
	PostRepository repository.PostRepository
}

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func (s *PostServiceImpl) GetPost(id uuid.UUID) (*repository2.Post, error) {
	return s.PostRepository.Get(id)
}

func (s *PostServiceImpl) CreatePost(post *repository2.Post) (*repository2.Post, error) {
	return s.PostRepository.Create(post)
}

func NewPostService(params PostServiceParams) PostService {
	return &PostServiceImpl{
		PostRepository: params.PostRepository,
	}
}
