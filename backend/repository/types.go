package repository

import "github.com/google/uuid"

type CrudRepository[T any] interface {
	Create(*T) (*T, error)
	Get(uuid.UUID) (*T, error)
	Update(*T) error
	Delete(uuid.UUID) error
}
