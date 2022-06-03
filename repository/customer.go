package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(cust *model.Dokter) error
	Update(cust *model.Dokter, updateValue interface{}) error
	Delete(cust *model.Dokter) error
	FindOne(id int) (*model.Dokter, error)
}

type customerRepository struct {
	dbConnection *gorm.DB
}

func NewCustomerRepository(dbConnection *gorm.DB) CustomerRepository {
	return &customerRepository{dbConnection: dbConnection}
}

func (c *customerRepository) Create(cust *model.Dokter) error {
	return c.dbConnection.Create(cust).Error
}

func (c *customerRepository) FindOne(id int) (cust *model.Dokter, err error) {
	cust = &model.Dokter{}
	err = c.dbConnection.First(cust, id).Error

	return
}

func (c *customerRepository) Update(cust *model.Dokter, updateValue interface{}) error {
	return c.dbConnection.Model(cust).Updates(updateValue).Error
}

func (c *customerRepository) Delete(cust *model.Dokter) error {
	return c.dbConnection.Delete(cust).Error
}
