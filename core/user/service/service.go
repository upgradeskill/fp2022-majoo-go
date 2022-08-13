package service

import (
	"mini-pos/core/user/entity"
	"mini-pos/core/user/repository"
	"net/url"
)

type Service interface {
	UserList(filters url.Values, limit int, offset int) ([]*entity.User, error)
	UserCreate(data entity.User) (*entity.User, error)
	UserFindById(id string) (*entity.User, error)
	UserUpdateById(id string, data entity.User) (*entity.User, error)
	UserDeleteById(id string) error
}

type service struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) Service {
	return &service{
		userRepository: userRepository,
	}
}
