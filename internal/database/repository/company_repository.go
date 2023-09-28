package repository

import (
	"crudgen/internal/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.DB.Create(company).Error
}

func (r *CompanyRepository) GetByID(id string) (*models.Company, error) {
	var company models.Company
	if err := r.DB.Where("id = ?", id).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Update(company *models.Company) error {
	return r.DB.Save(company).Error
}

func (r *CompanyRepository) Delete(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.Company{}).Error
}
