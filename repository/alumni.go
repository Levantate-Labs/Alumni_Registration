package repository

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/models"
	"github.com/akhil-is-watching/techletics_alumni_reg/types"
	"gorm.io/gorm"
)

type AlumniRepository struct {
	db *gorm.DB
}

func NewAlumniRepository(db *gorm.DB) AlumniRepository {
	return AlumniRepository{
		db: db,
	}
}

func (repo AlumniRepository) Create(input types.AlumniCreateInput) error {
	alumni := models.Alumni{
		ID:          input.ID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNo:     input.PhoneNo,
		ImageURL:    input.ImageURL,
		PassoutYear: input.PassoutYear,
		CheckedIn:   false,
	}

	if err := repo.db.Create(&alumni).Error; err != nil {
		return err
	}

	return nil
}

func (repo AlumniRepository) Get(ID string) (models.Alumni, error) {
	var alumni models.Alumni

	if err := repo.db.Where("id = ?", ID).First(&alumni).Error; err != nil {
		return alumni, err
	}

	return alumni, nil
}
