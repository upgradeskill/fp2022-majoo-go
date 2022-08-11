package service

import (
	"mini-pos/core/outlet/entity"

	"github.com/labstack/gommon/log"
)

func (service *service) OutletUserCreate(data entity.OutletUser) (*entity.OutletUser, error) {
	result, err := service.outletUserRepository.Create(data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) OutletFindByUserId(id string) (*entity.OutletUser, error) {
	result, err := service.outletUserRepository.FindById(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}
