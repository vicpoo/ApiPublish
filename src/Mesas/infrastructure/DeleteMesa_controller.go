package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/ApiPublish/src/Mesas/application"
)

type DeleteMesaController struct {
	deleteUseCase *application.DeleteMesaUseCase
}

func NewDeleteMesaController(deleteUseCase *application.DeleteMesaUseCase) *DeleteMesaController {
	return &DeleteMesaController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteMesaController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar la mesa",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Mesa eliminada exitosamente",
	})
}
