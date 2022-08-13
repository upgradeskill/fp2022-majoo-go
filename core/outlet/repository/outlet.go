package repository

import (
	"mini-pos/core/outlet/entity"
	"net/url"
)

type OutletRepository interface {
	List(filters url.Values, limit int, offset int) ([]*entity.Outlet, error)
	Create(data entity.Outlet) (*entity.Outlet, error)
	FindById(id string) (*entity.Outlet, error)
	UpdateById(id string, data entity.Outlet) (*entity.Outlet, error)
	DeleteById(id string) error
}
