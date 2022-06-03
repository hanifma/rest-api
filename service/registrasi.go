package service

import (
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
)

type RegistrasiService interface {
	GetRegistrasi(id int) (m *model.Registrasi, err error)
	CreateRegistrasi(registrasi *model.Registrasi) error
	UpdateRegistrasi(id int, registrasi *model.Registrasi) error
	DeleteRegistrasi(id int) error
}

type registrasiService struct {
	registrasis          map[int]*model.Registrasi
	registrasiRepository repository.RegistrasiRepository
}

func NewRegistrasiService(registrasiRepository repository.RegistrasiRepository) RegistrasiService {
	return &registrasiService{
		registrasiRepository: registrasiRepository,
	}
}

func (c *registrasiService) getRegistrasiById(id int) (m *model.Registrasi, err error) {
	m, err = c.registrasiRepository.FindOne(id)
	return
}

func (c *registrasiService) GetRegistrasi(id int) (m *model.Registrasi, err error) {
	return c.getRegistrasiById(id)
}

func (c *registrasiService) CreateRegistrasi(registrasi *model.Registrasi) error {
	return c.registrasiRepository.Create(registrasi)
}

func (c *registrasiService) UpdateRegistrasi(id int, registrasi *model.Registrasi) error {
	existingRegistrasi, err := c.getRegistrasiById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Registrasi{
		NIK: existingRegistrasi.NIK,
	}

	return c.registrasiRepository.Update(existingRegistrasi, updateValue)
}

func (c *registrasiService) DeleteRegistrasi(id int) error {
	existingRegistrasi, err := c.getRegistrasiById(id)
	if err != nil {
		return err
	}

	return c.registrasiRepository.Delete(existingRegistrasi)
}
