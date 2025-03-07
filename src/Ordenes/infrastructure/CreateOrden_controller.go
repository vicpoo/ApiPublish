// CreateOrden_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type CreateOrdenController struct {
	CreateOrdenUseCase *application.CreateOrdenUseCase
}

func NewCreateOrdenController(createOrdenUseCase *application.CreateOrdenUseCase) *CreateOrdenController {
	return &CreateOrdenController{
		CreateOrdenUseCase: createOrdenUseCase,
	}
}

func (ctrl *CreateOrdenController) Run(c *gin.Context) {
	var orden entities.Orden

	if errJSON := c.ShouldBindJSON(&orden); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos de la orden inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	ordenCreada, errAdd := ctrl.CreateOrdenUseCase.Run(&orden)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar la orden",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La orden ha sido agregada",
		"orden":   ordenCreada,
	})
}
