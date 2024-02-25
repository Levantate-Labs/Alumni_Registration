package repository

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/helpers"
	"github.com/akhil-is-watching/techletics_alumni_reg/models"
	"github.com/akhil-is-watching/techletics_alumni_reg/types"
	"gorm.io/gorm"
)

type VolunteerRepository struct {
	db *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) VolunteerRepository {
	return VolunteerRepository{
		db: db,
	}
}

func (repo VolunteerRepository) Create(input types.VolunteerCreateInput) error {
	volunteer := models.Volunteer{
		ID:       helpers.UIDGen().GenerateID("V"),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := repo.db.Create(&volunteer).Error; err != nil {
		return err
	}

	return nil
}

func (repo VolunteerRepository) GetByAuth(input types.LoginRequest) (models.Volunteer, error) {
	var volunteer models.Volunteer

	if err := repo.db.Where("email = ?", input.Email).Where("password = ?", input.Password).First(&volunteer).Error; err != nil {
		return volunteer, err
	}

	return volunteer, nil
}
