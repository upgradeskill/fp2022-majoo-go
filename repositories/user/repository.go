package user

import (
	"database/sql"
	"mini-pos/core/user/entity"
	"mini-pos/repositories/util/querybuilder"
	"net/url"
	"time"

	"gorm.io/gorm"
)

const TABLE_NAME = "users"

type User struct {
	Id           uint           `gorm:"column:id;primaryKey"`
	Email        sql.NullString `gorm:"column:email;size:64;unique"`
	Password     string         `gorm:"column:password;size:255"`
	PasswordSalt string         `gorm:"column:password_salt"`
	IsAdmin      bool           `gorm:"column:is_admin;default:0"`
	Created      time.Time      `gorm:"column:created;autoCreateTime"`
	Modified     time.Time      `gorm:"column:modified;autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return TABLE_NAME
}

func NewData(data entity.User) *User {
	return &User{
		Id:           data.Id,
		Email:        sql.NullString{String: data.Email},
		Password:     data.Password,
		PasswordSalt: data.PasswordSalt,
		IsAdmin:      data.IsAdmin,
		Created:      data.Created,
		Modified:     data.Modified,
	}
}

func (user *User) Map() entity.User {
	var data entity.User
	data.Id = user.Id
	data.Email = user.Email.String
	data.Password = user.Password
	data.PasswordSalt = user.PasswordSalt
	data.IsAdmin = user.IsAdmin
	data.Created = user.Created
	data.Modified = user.Modified
	return data
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repository, error) {
	repository := Repository{db}
	return &repository, nil
}

func (repository Repository) List(filters url.Values, limit int, offset int) ([]*entity.User, error) {
	users := []*User{}
	querybuilder.GormFilterBuilder(repository.db, filters, limit, offset).Find(&users)
	result := []*entity.User{}
	for _, data := range users {
		newData := data.Map()
		result = append(result, &newData)
	}
	return result, nil
}

func (repository Repository) Create(data entity.User) (*entity.User, error) {
	user := NewData(data)
	inserted := repository.db.Create(&user)
	if inserted.RowsAffected == 0 {
		return nil, inserted.Error
	}
	result := user.Map()
	return &result, nil
}

func (repository Repository) FindById(id string) (*entity.User, error) {
	user := User{}
	finded := repository.db.Find(&user, id)
	if finded.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	result := user.Map()
	return &result, nil
}

func (repository Repository) UpdateById(id string, data entity.User) (*entity.User, error) {
	user := User{}
	finded := repository.db.Find(&user, id)
	if finded.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	user.Email = sql.NullString{String: data.Email}
	user.IsAdmin = data.IsAdmin
	err := repository.db.Save(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	result := user.Map()
	return &result, nil
}

func (repository Repository) DeleteById(id string) error {
	user := User{}
	finded := repository.db.Find(&user, id)
	if finded.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	repository.db.Delete(&user)
	return nil
}
