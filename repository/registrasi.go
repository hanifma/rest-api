package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type RegistrasiRepository interface {
	Create(dok *model.Registrasi) error
	Update(dok *model.Registrasi, updateValue interface{}) error
	Delete(dok *model.Registrasi) error
	FindOne(id int) (*model.Registrasi, error)
}

type registrasiRepository struct {
	dbConnection *gorm.DB
}

func NewRegistrasiRepository(dbConnection *gorm.DB) RegistrasiRepository {
	return &registrasiRepository{dbConnection: dbConnection}
}

func (c *registrasiRepository) Create(dok *model.Registrasi) error {
	return c.dbConnection.Create(dok).Error
}

func (c *registrasiRepository) FindOne(id int) (dok *model.Registrasi, err error) {
	dok = &model.Registrasi{}
	err = c.dbConnection.First(dok, id).Error

	return
}

func (c *registrasiRepository) Update(dok *model.Registrasi, updateValue interface{}) error {
	return c.dbConnection.Model(dok).Updates(updateValue).Error
}

func (c *registrasiRepository) Delete(dok *model.Registrasi) error {
	return c.dbConnection.Delete(dok).Error
}
