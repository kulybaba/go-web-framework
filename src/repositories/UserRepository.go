package repositories

import (
	"github.com/petrokulybaba/go-basic-framework/configs"
	"github.com/petrokulybaba/go-basic-framework/src/models"
)

type UserRepository struct {
	findAll []models.User
	findOne models.User
	findBy  []models.User
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	if res := configs.DB.Find(&r.findAll); res.Error != nil {
		return r.findAll, res.Error
	}
	return r.findAll, nil
}

func (r *UserRepository) FindOneBy(find models.User) (models.User, error) {
	if res := configs.DB.First(&r.findOne, find); res.Error != nil {
		return r.findOne, res.Error
	}
	return r.findOne, nil
}

func (r *UserRepository) FindBy(find models.User) ([]models.User, error) {
	if res := configs.DB.Find(&r.findBy, find); res.Error != nil {
		return r.findBy, res.Error
	}
	return r.findBy, nil
}

func (r *UserRepository) Create(model *models.User) error {
	if res := configs.DB.Create(model); res.Error != nil {
		return res.Error
	}
	return nil
}
