package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type CreateDetallesOrdenController struct {
	CreateDetallesOrdenUseCase *application.CreateDetallesOrdenUseCase
}

func NewCreateDetallesOrdenController(createDetallesOrdenUseCase *application.CreateDetallesOrdenUseCase) *CreateDetallesOrdenController {
	return &CreateDetallesOrdenController{
		CreateDetallesOrdenUseCase: createDetallesOrdenUseCase,
	}
}

func (ctrl *CreateDetallesOrdenController) Run(c *gin.Context) {
	var detalle entities.DetallesOrden

	if errJSON := c.ShouldBindJSON(&detalle); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del detalle de orden inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	detalleCreado, errAdd := ctrl.CreateDetallesOrdenUseCase.Run(&detalle)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el detalle de orden",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "El detalle de orden ha sido agregado",
		"detalle": detalleCreado,
	})
}
