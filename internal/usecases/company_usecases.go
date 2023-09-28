package usecases

import (
	"crudgen/internal/database/repository"
	"crudgen/internal/models"
)

// CompanyUsecase representa o caso de uso para empresas
type CompanyUsecase struct {
	companyRepo repository.CompanyRepository
}

// NewCompanyUsecase cria uma nova instância de CompanyUsecase
func NewCompanyUsecase(companyRepo *repository.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{companyRepo: *companyRepo}
}

// CreateCompany cria uma nova empresa
func (uc *CompanyUsecase) CreateCompany(company *models.Company) error {
	// Implemente a lógica de validação e regras de negócios aqui, se necessário.
	return uc.companyRepo.Create(company)
}

// GetCompanyByID obtém uma empresa pelo ID
func (uc *CompanyUsecase) GetCompanyByID(id string) (*models.Company, error) {
	return uc.companyRepo.GetByID(id)
}

// UpdateCompany atualiza os dados de uma empresa existente
func (uc *CompanyUsecase) UpdateCompany(company *models.Company) error {
	// Implemente a lógica de validação e regras de negócios aqui, se necessário.
	return uc.companyRepo.Update(company)
}

// DeleteCompany exclui uma empresa pelo ID
func (uc *CompanyUsecase) DeleteCompany(id string) error {
	return uc.companyRepo.Delete(id)
}
