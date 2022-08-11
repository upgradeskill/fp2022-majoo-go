package service

import (
	"mini-pos/core/outlet/entity"
	"mini-pos/core/outlet/repository"
	"net/url"
)

type Service interface {
	OutletList(filters url.Values, limit int, offset int) ([]*entity.Outlet, error)
	OutletCreate(data entity.Outlet) (*entity.Outlet, error)
	OutletFindById(id string) (*entity.Outlet, error)
	OutletUpdateById(id string, data entity.Outlet) (*entity.Outlet, error)
	OutletDeleteById(id string) error
	OutletUserCreate(data entity.OutletUser) (*entity.OutletUser, error)
	OutletFindByUserId(id string) (*entity.OutletUser, error)
}

type service struct {
	outletRepository     repository.OutletRepository
	outletUserRepository repository.OutletUserRepository
}

func New(outletUserRepository repository.OutletUserRepository, outletRepository repository.OutletRepository) Service {
	return &service{
		outletRepository:     outletRepository,
		outletUserRepository: outletUserRepository,
	}
}
