// UpdateOrden_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type UpdateOrdenController struct {
	useCase *application.UpdateOrdenUseCase
}

func NewUpdateOrdenController(useCase *application.UpdateOrdenUseCase) *UpdateOrdenController {
	return &UpdateOrdenController{useCase: useCase}
}

func (uoc *UpdateOrdenController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var orden entities.Orden
	if err := c.ShouldBindJSON(&orden); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uoc.useCase.Run(id, orden)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la orden"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden actualizada exitosamente"})
}
