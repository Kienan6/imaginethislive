package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model/domain"
)

type GroupRepository interface {
	CrudRepository[domain.Group]
	FindByOwnership(ownerId uuid.UUID) (*[]domain.Group, error)
	FindUsers(id uuid.UUID) (*[]domain.User, error)
	FindPosts(id uuid.UUID) (*[]domain.Post, error)
}

type GroupRepositoryImpl struct {
	db *gorm.DB
}

type GroupRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *GroupRepositoryImpl) Create(group *domain.Group) (*domain.Group, error) {
	res := u.db.Save(group)
	return group, res.Error
}

func (u *GroupRepositoryImpl) Get(id uuid.UUID) (*domain.Group, error) {
	var group domain.Group
	tx := u.db.Preload("Owner").Session(&gorm.Session{})
	res := tx.Where("id = ?", id.String()).First(&group)
	return &group, res.Error
}

func (u *GroupRepositoryImpl) Update(group *domain.Group) error {
	res := u.db.Update(group.ID.String(), group)
	return res.Error
}

func (u *GroupRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("groups").Delete(id.String())
	return res.Error
}

func (u *GroupRepositoryImpl) FindByOwnership(ownerId uuid.UUID) (*[]domain.Group, error) {
	var groups []domain.Group
	tx := u.db.Preload("Owner").Session(&gorm.Session{})
	res := tx.Find(&groups, "owner_id = ?", ownerId.String())
	return &groups, res.Error
}

func (u *GroupRepositoryImpl) FindPosts(id uuid.UUID) (*[]domain.Post, error) {
	var posts []domain.Post
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&domain.Group{ID: id}).Preload("User").Association("Posts").Find(&posts)
	return &posts, err
}

func (u *GroupRepositoryImpl) FindUsers(id uuid.UUID) (*[]domain.User, error) {
	var users []domain.User
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&domain.Group{ID: id}).Association("Users").Find(&users)
	return &users, err
}

func NewGroupRepository(params GroupRepositoryParams) GroupRepository {
	return &GroupRepositoryImpl{
		db: params.Db.Db,
	}
}
