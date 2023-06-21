package repository

import "github.com/google/uuid"

type CrudRepository[T any] interface {
	Create(user *T) error
	Get(id uuid.UUID) (*T, error)
	Update(user *T) error
	Delete(id uuid.UUID) error
}
