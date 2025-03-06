package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Mesas/application"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type CreateMesaController struct {
	CreateMesaUseCase *application.CreateMesaUseCase
}

func NewCreateMesaController(createMesaUseCase *application.CreateMesaUseCase) *CreateMesaController {
	return &CreateMesaController{
		CreateMesaUseCase: createMesaUseCase,
	}
}

func (ctrl *CreateMesaController) Run(c *gin.Context) {
	var mesa entities.Mesa

	if errJSON := c.ShouldBindJSON(&mesa); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos de la mesa inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	mesaCreada, errAdd := ctrl.CreateMesaUseCase.Run(&mesa)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar la mesa",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La mesa ha sido agregada",
		"mesa":    mesaCreada,
	})
}
