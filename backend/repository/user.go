package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model"
)

type UserRepository interface {
	CrudRepository[model.User]
	FindGroups(id uuid.UUID) (*[]model.Group, error)
	AddToGroup(id uuid.UUID, groupId uuid.UUID) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) FindGroups(id uuid.UUID) (*[]model.Group, error) {
	var groups []model.Group
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&model.User{ID: id}).Association("Groups").Find(&groups)
	return &groups, err
}

type UserRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {
	res := u.db.Save(user)
	return user, res.Error
}

func (u *UserRepositoryImpl) Get(id uuid.UUID) (*model.User, error) {
	var user model.User
	res := u.db.Where("id = ?", id.String()).First(&user)
	return &user, res.Error
}

func (u *UserRepositoryImpl) Update(user *model.User) error {
	res := u.db.Update(user.ID.String(), user)
	return res.Error
}

func (u *UserRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("users").Delete(id.String())
	return res.Error
}

func (u *UserRepositoryImpl) AddToGroup(id uuid.UUID, groupId uuid.UUID) error {
	res := u.db.Model(&model.User{ID: id}).Association("Groups").Append(&model.Group{ID: groupId})
	return res
}

func NewUserRepository(params UserRepositoryParams) UserRepository {
	return &UserRepositoryImpl{
		db: params.Db.Db,
	}
}
