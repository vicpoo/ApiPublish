package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type CreatePlatilloController struct {
	CreatePlatilloUseCase *application.CreatePlatilloUseCase
}

func NewCreatePlatilloController(createPlatilloUseCase *application.CreatePlatilloUseCase) *CreatePlatilloController {
	return &CreatePlatilloController{
		CreatePlatilloUseCase: createPlatilloUseCase,
	}
}

func (ctrl *CreatePlatilloController) Run(c *gin.Context) {
	var platillo entities.Platillo

	if errJSON := c.ShouldBindJSON(&platillo); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del platillo inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	platilloCreado, errAdd := ctrl.CreatePlatilloUseCase.Run(&platillo)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el platillo",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "El platillo ha sido agregado",
		"platillo": platilloCreado,
	})
}
