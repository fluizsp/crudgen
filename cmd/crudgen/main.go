package main

import (
	"crudgen/internal/api"
	initializer "crudgen/internal/tools"
	"crudgen/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Configurar o roteador Gin
	router := gin.Default()

	// Inicializar os repositórios
	err := godotenv.Load()
	companyRepo, err := initializer.InitializeRepositories()
	if err != nil {
		// Lidar com erros de inicialização aqui
		panic(err)
	}

	// Inicializar os casos de uso
	companyUsecase := usecases.NewCompanyUsecase(companyRepo)

	// Configurar as rotas para a entidade Company
	api.SetupCompanyRoutes(router.Group("/api/v1"), *companyUsecase)

	// Iniciar o servidor na porta desejada
	router.Run(":8080")
}
