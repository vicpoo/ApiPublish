package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type UpdatePlatilloController struct {
	useCase *application.UpdatePlatilloUseCase
}

func NewUpdatePlatilloController(useCase *application.UpdatePlatilloUseCase) *UpdatePlatilloController {
	return &UpdatePlatilloController{useCase: useCase}
}

func (upc *UpdatePlatilloController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var platillo entities.Platillo
	if err := c.ShouldBindJSON(&platillo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = upc.useCase.Run(id, platillo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el platillo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Platillo actualizado exitosamente"})
}
