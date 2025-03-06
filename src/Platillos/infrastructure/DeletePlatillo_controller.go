package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/ApiPublish/src/Platillos/application"
)

type DeletePlatilloController struct {
	deleteUseCase *application.DeletePlatilloUseCase
}

func NewDeletePlatilloController(deleteUseCase *application.DeletePlatilloUseCase) *DeletePlatilloController {
	return &DeletePlatilloController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeletePlatilloController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el platillo",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Platillo eliminado exitosamente",
	})
}
