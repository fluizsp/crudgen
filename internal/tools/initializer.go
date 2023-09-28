package initializer

import (
	"crudgen/internal/database"
	"crudgen/internal/database/repository"
)

func InitializeRepositories() (*repository.CompanyRepository, error) {
	// Conectar ao banco de dados
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// Inicializar o repositório da empresa
	companyRepo := repository.NewCompanyRepository(db)

	return companyRepo, nil
}
