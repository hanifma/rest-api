package service

import (
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
)

type PasienService interface {
	GetPasien(id string) (m *model.Pasien, err error)
	CreatePasien(pasien *model.Pasien) error
	UpdatePasien(id string, pasien *model.Pasien) error
	DeletePasien(id string) error
}

type pasienService struct {
	pasiens          map[int]*model.Pasien
	pasienRepository repository.PasienRepository
}

func NewPasienService(pasienRepository repository.PasienRepository) PasienService {
	return &pasienService{
		pasienRepository: pasienRepository,
	}
}

func (c *pasienService) getPasienById(id string) (m *model.Pasien, err error) {
	m, err = c.pasienRepository.FindOne(id)
	return
}

func (c *pasienService) GetPasien(id string) (m *model.Pasien, err error) {
	return c.getPasienById(id)
}

func (c *pasienService) CreatePasien(pasien *model.Pasien) error {
	return c.pasienRepository.Create(pasien)
}

func (c *pasienService) UpdatePasien(id string, pasien *model.Pasien) error {
	existingPasien, err := c.getPasienById(id)
	if err != nil {
		return err
	}

	var updateValue = model.Pasien{
		Nama:   pasien.Nama,
		Alamat: pasien.Alamat,
		NoHp:   pasien.NoHp,
	}

	return c.pasienRepository.Update(existingPasien, updateValue)
}

func (c *pasienService) DeletePasien(id string) error {
	existingPasien, err := c.getPasienById(id)
	if err != nil {
		return err
	}

	return c.pasienRepository.Delete(existingPasien)
}
