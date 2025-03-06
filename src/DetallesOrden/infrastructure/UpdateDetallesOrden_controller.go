package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type UpdateDetallesOrdenController struct {
	useCase *application.UpdateDetallesOrdenUseCase
}

func NewUpdateDetallesOrdenController(useCase *application.UpdateDetallesOrdenUseCase) *UpdateDetallesOrdenController {
	return &UpdateDetallesOrdenController{useCase: useCase}
}

func (uc *UpdateDetallesOrdenController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var detalle entities.DetallesOrden
	if err := c.ShouldBindJSON(&detalle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.useCase.Execute(id, detalle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el detalle de orden"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Detalle de orden actualizado exitosamente"})
}
