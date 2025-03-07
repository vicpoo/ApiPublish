// DeleteOrden_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

type DeleteOrdenController struct {
	deleteUseCase *application.DeleteOrdenUseCase
}

func NewDeleteOrdenController(deleteUseCase *application.DeleteOrdenUseCase) *DeleteOrdenController {
	return &DeleteOrdenController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteOrdenController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar la orden",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Orden eliminada exitosamente",
	})
}
