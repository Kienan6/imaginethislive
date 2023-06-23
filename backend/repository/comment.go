package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model"
)

type CommentRepository interface {
	CrudRepository[model.Comment]
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

type CommentRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *CommentRepositoryImpl) Create(Comment *model.Comment) error {
	res := u.db.Save(Comment)
	return res.Error
}

func (u *CommentRepositoryImpl) Get(id uuid.UUID) (*model.Comment, error) {
	var comment model.Comment
	tx := u.db.Model(&model.Comment{ID: id}).Preload("User").Preload("Post")
	res := tx.First(&comment)
	return &comment, res.Error
}

func (u *CommentRepositoryImpl) Update(Comment *model.Comment) error {
	res := u.db.Update(Comment.ID.String(), Comment)
	return res.Error
}

func (u *CommentRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("Comments").Delete(id.String())
	return res.Error
}

func NewCommentRepository(params CommentRepositoryParams) CommentRepository {
	return &CommentRepositoryImpl{
		db: params.Db.Db,
	}
}
