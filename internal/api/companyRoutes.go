package api

import (
	"net/http"

	"crudgen/internal/models"

	"crudgen/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SetupCompanyRoutes(router *gin.RouterGroup, companyUsecase usecases.CompanyUsecase) {
	// Rota para criar uma nova empresa
	router.POST("/companies", func(c *gin.Context) {
		var company models.Company
		if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := companyUsecase.CreateCompany(&company); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, company)
	})

	// Rota para obter uma empresa pelo ID
	router.GET("/companies/:id", func(c *gin.Context) {
		id := c.Param("id")
		company, err := companyUsecase.GetCompanyByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Empresa não encontrada"})
			return
		}
		c.JSON(http.StatusOK, company)
	})

	// Rota para atualizar uma empresa
	router.PUT("/companies/:id", func(c *gin.Context) {
		id := c.Param("id")
		var company models.Company
		if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		company.ID = id
		if err := companyUsecase.UpdateCompany(&company); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, company)
	})

	// Rota para excluir uma empresa pelo ID
	router.DELETE("/companies/:id", func(c *gin.Context) {
		id := c.Param("id")
		if err := companyUsecase.DeleteCompany(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Empresa excluída com sucesso"})
	})
}
