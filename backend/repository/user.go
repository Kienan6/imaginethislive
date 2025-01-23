package repository

import (
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"itl/db"
	"itl/model/domain"
)

type UserRepository interface {
	CrudRepository[domain.User]
	FindGroups(id uuid.UUID) (*[]domain.Group, error)
	AddToGroup(id uuid.UUID, groupId uuid.UUID) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

type UserRepositoryParams struct {
	fx.In
	Db *db.PostgresConnection
}

func (u *UserRepositoryImpl) FindGroups(id uuid.UUID) (*[]domain.Group, error) {
	var groups []domain.Group
	tx := u.db.Session(&gorm.Session{})
	err := tx.Model(&domain.User{ID: id}).Association("Groups").Find(&groups)
	return &groups, err
}

func (u *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	res := u.db.Save(user)
	return user, res.Error
}

func (u *UserRepositoryImpl) Get(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	res := u.db.Where("id = ?", id.String()).First(&user)
	return &user, res.Error
}

func (u *UserRepositoryImpl) Update(user *domain.User) error {
	res := u.db.Update(user.ID.String(), user)
	return res.Error
}

func (u *UserRepositoryImpl) Delete(id uuid.UUID) error {
	res := u.db.Table("users").Delete(id.String())
	return res.Error
}

func (u *UserRepositoryImpl) AddToGroup(id uuid.UUID, groupId uuid.UUID) error {
	res := u.db.Model(&domain.User{ID: id}).Association("Groups").Append(&domain.Group{ID: groupId})
	return res
}

func NewUserRepository(params UserRepositoryParams) UserRepository {
	return &UserRepositoryImpl{
		db: params.Db.Db,
	}
}
