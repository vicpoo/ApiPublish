package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

type ViewOrdenController struct {
	useCase *application.ViewOrdenUseCase
}

func NewViewOrdenController(useCase *application.ViewOrdenUseCase) *ViewOrdenController {
	return &ViewOrdenController{useCase: useCase}
}

func (voc *ViewOrdenController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	orden, err := voc.useCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orden no encontrada"})
		return
	}

	c.JSON(http.StatusOK, orden)
}
