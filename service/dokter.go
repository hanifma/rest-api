package service

import (
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
)

type DokterService interface {
	GetDokter(id int) (m *model.Dokter, err error)
	CreateDokter(dokter *model.Dokter) error
	UpdateDokter(id int, dokter *model.Dokter) error
	DeleteDokter(id int) error
}

type dokterService struct {
	dokters          map[int]*model.Dokter
	dokterRepository repository.DokterRepository
}

func NewDokterService(dokterRepository repository.DokterRepository) DokterService {
	return &dokterService{
		dokterRepository: dokterRepository,
	}
}

func (c *dokterService) getDokterById(id int) (m *model.Dokter, err error) {
	m, err = c.dokterRepository.FindOne(id)
	return
}

func (c *dokterService) GetDokter(id int) (m *model.Dokter, err error) {
	return c.getDokterById(id)
}

func (c *dokterService) CreateDokter(dokter *model.Dokter) error {
	return c.dokterRepository.Create(dokter)
}

func (c *dokterService) UpdateDokter(id int, dokter *model.Dokter) error {
	existingDokter, err := c.getDokterById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Dokter{
		Nama:  dokter.Nama,
		Email: dokter.Email,
	}

	return c.dokterRepository.Update(existingDokter, updateValue)
}

func (c *dokterService) DeleteDokter(id int) error {
	existingDokter, err := c.getDokterById(id)
	if err != nil {
		return err
	}

	return c.dokterRepository.Delete(existingDokter)
}
