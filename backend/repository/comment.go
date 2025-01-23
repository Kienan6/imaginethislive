package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model/domain"
)

type CommentRepository interface {
	CrudRepository[domain.Comment]
}

type CommentRepositoryImpl struct {
	db *gorm.DB
}

type CommentRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *CommentRepositoryImpl) Create(comment *domain.Comment) (*domain.Comment, error) {
	res := u.db.Save(comment)
	return comment, res.Error
}

func (u *CommentRepositoryImpl) Get(id uuid.UUID) (*domain.Comment, error) {
	var comment domain.Comment
	tx := u.db.Model(&domain.Comment{ID: id}).Preload("User").Preload("Post")
	res := tx.First(&comment)
	return &comment, res.Error
}

func (u *CommentRepositoryImpl) Update(comment *domain.Comment) error {
	res := u.db.Update(comment.ID.String(), comment)
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
