package outlet

import (
	"mini-pos/core/outlet/entity"
	"mini-pos/repositories/util/querybuilder"
	"net/url"

	"gorm.io/gorm"
)

const TABLE_NAME = "outlets"

type Outlet struct {
	Id   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:password;size:255"`
}

type Tabler interface {
	TableName() string
}

func (Outlet) TableName() string {
	return TABLE_NAME
}

func NewData(data entity.Outlet) *Outlet {
	return &Outlet{
		Id:   data.Id,
		Name: data.Name,
	}
}

func (outlet *Outlet) Map() entity.Outlet {
	var data entity.Outlet
	data.Id = outlet.Id
	data.Name = outlet.Name

	return data
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repository, error) {
	repository := Repository{db}
	return &repository, nil
}

func (repository Repository) List(filters url.Values, limit int, offset int) ([]*entity.Outlet, error) {
	outlets := []*Outlet{}
	querybuilder.GormFilterBuilder(repository.db, filters, limit, offset).Find(&outlets)
	result := []*entity.Outlet{}
	for _, data := range outlets {
		newData := data.Map()
		result = append(result, &newData)
	}
	return result, nil
}

func (repository Repository) Create(data entity.Outlet) (*entity.Outlet, error) {
	outlet := NewData(data)
	inserted := repository.db.Create(&outlet)
	if inserted.RowsAffected == 0 {
		return nil, inserted.Error
	}
	result := outlet.Map()
	return &result, nil
}

func (repository Repository) FindById(id string) (*entity.Outlet, error) {
	outlet := Outlet{}
	finded := repository.db.Find(&outlet, id)
	if finded.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	result := outlet.Map()
	return &result, nil
}

func (repository Repository) UpdateById(id string, data entity.Outlet) (*entity.Outlet, error) {
	outlet := Outlet{}
	finded := repository.db.Find(&outlet, id)
	if finded.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	outlet.Name = data.Name
	err := repository.db.Save(&outlet)
	if err.Error != nil {
		return nil, err.Error
	}
	result := outlet.Map()
	return &result, nil
}

func (repository Repository) DeleteById(id string) error {
	outlet := Outlet{}
	finded := repository.db.Find(&outlet, id)
	if finded.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	repository.db.Delete(&outlet)
	return nil
}
