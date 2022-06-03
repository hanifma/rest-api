package service

import (
	"github.com/pushm0v/gorest/model"
	"github.com/pushm0v/gorest/repository"
)

type JenisPenyakitService interface {
	GetJenisPenyakit(id int) (m *model.JenisPenyakit, err error)
	CreateJenisPenyakit(passien *model.JenisPenyakit) error
	UpdateJenisPenyakit(id int, passien *model.JenisPenyakit) error
	DeleteJenisPenyakit(id int) error
}

type jenisPenyakitService struct {
	jenisPenyakits    map[int]*model.JenisPenyakit
	passienRepository repository.JenisPenyakitRepository
}

func NewJenisPenyakitService(passienRepository repository.JenisPenyakitRepository) JenisPenyakitService {
	return &jenisPenyakitService{
		passienRepository: passienRepository,
	}
}

func (c *jenisPenyakitService) getJenisPenyakitById(id int) (m *model.JenisPenyakit, err error) {
	m, err = c.passienRepository.FindOne(id)
	return
}

func (c *jenisPenyakitService) GetJenisPenyakit(id int) (m *model.JenisPenyakit, err error) {
	return c.getJenisPenyakitById(id)
}

func (c *jenisPenyakitService) CreateJenisPenyakit(passien *model.JenisPenyakit) error {
	return c.passienRepository.Create(passien)
}

func (c *jenisPenyakitService) UpdateJenisPenyakit(id int, passien *model.JenisPenyakit) error {
	existingJenisPenyakit, err := c.getJenisPenyakitById(id)
	if err != nil {
		return err
	}

	var updateValue = model.JenisPenyakit{
		Nama: passien.Nama,
	}

	return c.passienRepository.Update(existingJenisPenyakit, updateValue)
}

func (c *jenisPenyakitService) DeleteJenisPenyakit(id int) error {
	existingJenisPenyakit, err := c.getJenisPenyakitById(id)
	if err != nil {
		return err
	}

	return c.passienRepository.Delete(existingJenisPenyakit)
}
