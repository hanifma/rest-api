package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type JenisPenyakitRepository interface {
	Create(dok *model.JenisPenyakit) error
	Update(dok *model.JenisPenyakit, updateValue interface{}) error
	Delete(dok *model.JenisPenyakit) error
	FindOne(id int) (*model.JenisPenyakit, error)
}

type jenisPenyakitRepository struct {
	dbConnection *gorm.DB
}

func NewJenisPenyakitRepository(dbConnection *gorm.DB) JenisPenyakitRepository {
	return &jenisPenyakitRepository{dbConnection: dbConnection}
}

func (c *jenisPenyakitRepository) Create(dok *model.JenisPenyakit) error {
	return c.dbConnection.Create(dok).Error
}

func (c *jenisPenyakitRepository) FindOne(id int) (dok *model.JenisPenyakit, err error) {
	dok = &model.JenisPenyakit{}
	err = c.dbConnection.First(dok, id).Error

	return
}

func (c *jenisPenyakitRepository) Update(dok *model.JenisPenyakit, updateValue interface{}) error {
	return c.dbConnection.Model(dok).Updates(updateValue).Error
}

func (c *jenisPenyakitRepository) Delete(dok *model.JenisPenyakit) error {
	return c.dbConnection.Delete(dok).Error
}
