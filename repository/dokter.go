package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type DokterRepository interface {
	Create(dok *model.Dokter) error
	Update(dok *model.Dokter, updateValue interface{}) error
	Delete(dok *model.Dokter) error
	FindOne(id int) (*model.Dokter, error)
}

type dokterRepository struct {
	dbConnection *gorm.DB
}

func NewDokterRepository(dbConnection *gorm.DB) DokterRepository {
	return &dokterRepository{dbConnection: dbConnection}
}

func (c *dokterRepository) Create(dok *model.Dokter) error {
	return c.dbConnection.Create(dok).Error
}

func (c *dokterRepository) FindOne(id int) (dok *model.Dokter, err error) {
	dok = &model.Dokter{}
	err = c.dbConnection.First(dok, id).Error

	return
}

func (c *dokterRepository) Update(dok *model.Dokter, updateValue interface{}) error {
	return c.dbConnection.Model(dok).Updates(updateValue).Error
}

func (c *dokterRepository) Delete(dok *model.Dokter) error {
	return c.dbConnection.Delete(dok).Error
}
