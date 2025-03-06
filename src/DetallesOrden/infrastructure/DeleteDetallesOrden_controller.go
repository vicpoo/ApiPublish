package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
)

type DeleteDetallesOrdenController struct {
	deleteUseCase *application.DeleteDetallesOrdenUseCase
}

func NewDeleteDetallesOrdenController(deleteUseCase *application.DeleteDetallesOrdenUseCase) *DeleteDetallesOrdenController {
	return &DeleteDetallesOrdenController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteDetallesOrdenController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el detalle de orden",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Detalle de orden eliminado exitosamente",
	})
}
