package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type PasienRepository interface {
	Create(dok *model.Pasien) error
	Update(dok *model.Pasien, updateValue interface{}) error
	Delete(dok *model.Pasien) error
	FindOne(id string) (*model.Pasien, error)
}

type pasienRepository struct {
	dbConnection *gorm.DB
}

func NewPasienRepository(dbConnection *gorm.DB) PasienRepository {
	return &pasienRepository{dbConnection: dbConnection}
}

func (c *pasienRepository) Create(dok *model.Pasien) error {
	return c.dbConnection.Create(dok).Error
}

func (c *pasienRepository) FindOne(id string) (dok *model.Pasien, err error) {
	dok = &model.Pasien{}
	err = c.dbConnection.First(dok, id).Error

	return
}

func (c *pasienRepository) Update(dok *model.Pasien, updateValue interface{}) error {
	return c.dbConnection.Model(dok).Updates(updateValue).Error
}

func (c *pasienRepository) Delete(dok *model.Pasien) error {
	return c.dbConnection.Delete(dok).Error
}
