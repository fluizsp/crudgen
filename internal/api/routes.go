package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Rotas de configuração
	configRoutes := r.Group("/config")
	{
		configRoutes.GET("/object_type/:id", GetObjectType)
		configRoutes.POST("/object_type", CreateObjectType)
		configRoutes.PUT("/object_type/:id", UpdateObjectType)
		configRoutes.DELETE("/object_type/:id", DeleteObjectType)

		configRoutes.GET("/object_types", GetObjectTypes)
		configRoutes.POST("/object_types", CreateObjectTypes)
		configRoutes.PUT("/object_types/:id", UpdateObjectTypes)
		configRoutes.DELETE("/object_types/:id", DeleteObjectTypes)
	}
}

func GetObjectType(c *gin.Context) {
	// Implemente a lógica para buscar um OBJECT_TYPE por ID
}

func CreateObjectType(c *gin.Context) {
	// Implemente a lógica para criar um novo OBJECT_TYPE
}

func UpdateObjectType(c *gin.Context) {
	// Implemente a lógica para atualizar um OBJECT_TYPE por ID
}

func DeleteObjectType(c *gin.Context) {
	// Implemente a lógica para excluir um OBJECT_TYPE por ID
}

func GetObjectTypes(c *gin.Context) {
	// Implemente a lógica para buscar todos os OBJECT_TYPES
	c.JSON(http.StatusOK, gin.H{"Message": "Hello World"})
}

func CreateObjectTypes(c *gin.Context) {
	// Implemente a lógica para criar novos OBJECT_TYPES
}

func UpdateObjectTypes(c *gin.Context) {
	// Implemente a lógica para atualizar OBJECT_TYPES por ID
}

func DeleteObjectTypes(c *gin.Context) {
	// Implemente a lógica para excluir OBJECT_TYPES por ID
}
