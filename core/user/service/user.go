package service

import (
	"mini-pos/core/user/entity"
	"mini-pos/core/user/util"
	"net/url"

	"github.com/labstack/gommon/log"
)

func (service *service) UserList(filters url.Values, limit int, offset int) ([]*entity.User, error) {
	result, err := service.userRepository.List(filters, limit, offset)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) UserCreate(data entity.User) (*entity.User, error) {
	password, err := util.GeneratePassword(data.Password, data.PasswordSalt)
	if err != nil {
		return nil, err
	}
	data.Password = password
	result, err := service.userRepository.Create(data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) UserFindById(id string) (*entity.User, error) {
	result, err := service.userRepository.FindById(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) UserUpdateById(id string, data entity.User) (*entity.User, error) {
	result, err := service.userRepository.UpdateById(id, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) UserDeleteById(id string) error {
	err := service.userRepository.DeleteById(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
