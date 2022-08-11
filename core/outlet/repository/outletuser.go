package repository

import (
	"mini-pos/core/outlet/entity"
)

type OutletUserRepository interface {
	Create(data entity.OutletUser) (*entity.OutletUser, error)
	FindById(id string) (*entity.OutletUser, error)
}
