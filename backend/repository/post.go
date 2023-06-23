package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model"
)

type PostRepository interface {
	CrudRepository[model.Post]
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

type PostRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *PostRepositoryImpl) Create(post *model.Post) (*model.Post, error) {
	res := u.db.Save(post)
	return post, res.Error
}

func (u *PostRepositoryImpl) Get(id uuid.UUID) (*model.Post, error) {
	var post model.Post
	tx := u.db.Model(&model.Post{ID: id}).Preload("Comments")
	res := tx.First(&post)
	return &post, res.Error
}

func (u *PostRepositoryImpl) Update(post *model.Post) error {
	res := u.db.Update(post.ID.String(), post)
	return res.Error
}

func (u *PostRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("Posts").Delete(id.String())
	return res.Error
}

func NewPostRepository(params PostRepositoryParams) PostRepository {
	return &PostRepositoryImpl{
		db: params.Db.Db,
	}
}
