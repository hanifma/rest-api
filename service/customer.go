package service

import (
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
)

type CustomerService interface {
	GetCustomer(id int) (m *model.Dokter, err error)
	CreateCustomer(cust *model.Dokter) error
	UpdateCustomer(id int, cust *model.Dokter) error
	DeleteCustomer(id int) error
}

type customerService struct {
	customers      map[int]*model.Dokter
	custRepository repository.CustomerRepository
}

func NewCustomerService(custRepository repository.CustomerRepository) CustomerService {
	return &customerService{
		custRepository: custRepository,
	}
}

func (c *customerService) getCustomerById(id int) (m *model.Dokter, err error) {
	m, err = c.custRepository.FindOne(id)
	return
}

func (c *customerService) GetCustomer(id int) (m *model.Dokter, err error) {
	return c.getCustomerById(id)
}

func (c *customerService) CreateCustomer(cust *model.Dokter) error {
	return c.custRepository.Create(cust)
}

func (c *customerService) UpdateCustomer(id int, cust *model.Dokter) error {
	existingCustomer, err := c.getCustomerById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Dokter{
		Nama:  cust.Nama,
		Email: cust.Email,
	}

	return c.custRepository.Update(existingCustomer, updateValue)
}

func (c *customerService) DeleteCustomer(id int) error {
	existingCustomer, err := c.getCustomerById(id)
	if err != nil {
		return err
	}

	return c.custRepository.Delete(existingCustomer)
}
