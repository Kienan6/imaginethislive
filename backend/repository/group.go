package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model"
)

type GroupRepository interface {
	CrudRepository[model.Group]
	FindByOwnership(ownerId uuid.UUID) (*[]model.Group, error)
	FindUsers(id uuid.UUID) (*[]model.User, error)
	FindPosts(id uuid.UUID) (*[]model.Post, error)
}

type GroupRepositoryImpl struct {
	db *gorm.DB
}

type GroupRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *GroupRepositoryImpl) Create(group *model.Group) (*model.Group, error) {
	res := u.db.Save(group)
	return group, res.Error
}

func (u *GroupRepositoryImpl) Get(id uuid.UUID) (*model.Group, error) {
	var group model.Group
	tx := u.db.Preload("Owner").Session(&gorm.Session{})
	res := tx.Where("id = ?", id.String()).First(&group)
	return &group, res.Error
}

func (u *GroupRepositoryImpl) Update(group *model.Group) error {
	res := u.db.Update(group.ID.String(), group)
	return res.Error
}

func (u *GroupRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("groups").Delete(id.String())
	return res.Error
}

func (u *GroupRepositoryImpl) FindByOwnership(ownerId uuid.UUID) (*[]model.Group, error) {
	var groups []model.Group
	tx := u.db.Preload("Owner").Session(&gorm.Session{})
	res := tx.Find(&groups, "owner_id = ?", ownerId.String())
	return &groups, res.Error
}

func (u *GroupRepositoryImpl) FindPosts(id uuid.UUID) (*[]model.Post, error) {
	var posts []model.Post
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&model.Group{ID: id}).Preload("User").Association("Posts").Find(&posts)
	return &posts, err
}

func (u *GroupRepositoryImpl) FindUsers(id uuid.UUID) (*[]model.User, error) {
	var users []model.User
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&model.Group{ID: id}).Association("Users").Find(&users)
	return &users, err
}

func NewGroupRepository(params GroupRepositoryParams) GroupRepository {
	return &GroupRepositoryImpl{
		db: params.Db.Db,
	}
}
