package outletuser

import (
	"mini-pos/core/outlet/entity"

	"gorm.io/gorm"
)

const TABLE_NAME = "outlets"

type OutletUser struct {
	Id       uint `gorm:"column:id;primaryKey"`
	OutletId uint `gorm:"column:outlet_id"`
	UserId   uint `gorm:"column:user_id"`
}

type Tabler interface {
	TableName() string
}

func (OutletUser) TableName() string {
	return TABLE_NAME
}

func NewData(data entity.OutletUser) *OutletUser {
	return &OutletUser{
		Id:       data.Id,
		OutletId: data.OutletId,
	}
}

func (outletUser *OutletUser) Map() entity.OutletUser {
	var data entity.OutletUser
	data.Id = outletUser.Id
	data.OutletId = outletUser.OutletId
	data.UserId = outletUser.UserId

	return data
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repository, error) {
	repository := Repository{db}
	return &repository, nil
}

func (repository Repository) Create(data entity.OutletUser) (*entity.OutletUser, error) {
	outletUser := NewData(data)
	inserted := repository.db.Create(&outletUser)
	if inserted.RowsAffected == 0 {
		return nil, inserted.Error
	}
	result := outletUser.Map()
	return &result, nil
}

func (repository Repository) FindById(id string) (*entity.OutletUser, error) {
	outletUser := OutletUser{}
	finded := repository.db.Find(&outletUser, id)
	if finded.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	result := outletUser.Map()
	return &result, nil
}
