package service

import (
	"mini-pos/core/outlet/entity"
	"net/url"

	"github.com/labstack/gommon/log"
)

func (service *service) OutletList(filters url.Values, limit int, offset int) ([]*entity.Outlet, error) {
	result, err := service.outletRepository.List(filters, limit, offset)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) OutletCreate(data entity.Outlet) (*entity.Outlet, error) {
	result, err := service.outletRepository.Create(data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) OutletFindById(id string) (*entity.Outlet, error) {
	result, err := service.outletRepository.FindById(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) OutletUpdateById(id string, data entity.Outlet) (*entity.Outlet, error) {
	result, err := service.outletRepository.UpdateById(id, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}

func (service *service) OutletDeleteById(id string) error {
	err := service.outletRepository.DeleteById(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
