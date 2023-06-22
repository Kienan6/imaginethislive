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

func (u *PostRepositoryImpl) Create(Post *model.Post) error {
	res := u.db.Save(Post)
	return res.Error
}

func (u *PostRepositoryImpl) Get(id uuid.UUID) (*model.Post, error) {
	var Post model.Post
	res := u.db.Where("id = ?", id.String()).First(&Post)
	return &Post, res.Error
}

func (u *PostRepositoryImpl) Update(Post *model.Post) error {
	res := u.db.Update(Post.ID.String(), Post)
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
