package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Mesas/application"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type UpdateMesaController struct {
	useCase *application.UpdateMesaUseCase
}

func NewUpdateMesaController(useCase *application.UpdateMesaUseCase) *UpdateMesaController {
	return &UpdateMesaController{useCase: useCase}
}

func (uec *UpdateMesaController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var mesa entities.Mesa
	if err := c.ShouldBindJSON(&mesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uec.useCase.Run(id, mesa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la mesa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mesa actualizada exitosamente"})
}
