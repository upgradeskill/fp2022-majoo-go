package repository

import (
	"mini-pos/core/user/entity"
	"net/url"
)

type UserRepository interface {
	List(filters url.Values, limit int, offset int) ([]*entity.User, error)
	Create(data entity.User) (*entity.User, error)
	FindById(id string) (*entity.User, error)
	UpdateById(id string, data entity.User) (*entity.User, error)
	DeleteById(id string) error
}
