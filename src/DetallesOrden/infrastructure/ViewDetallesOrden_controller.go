package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
)

type ViewDetallesOrdenController struct {
	useCase *application.ViewDetallesOrdenUseCase
}

func NewViewDetallesOrdenController(useCase *application.ViewDetallesOrdenUseCase) *ViewDetallesOrdenController {
	return &ViewDetallesOrdenController{useCase: useCase}
}

func (ctrl *ViewDetallesOrdenController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	detalle, err := ctrl.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Detalle de orden no encontrado"})
		return
	}

	c.JSON(http.StatusOK, detalle)
}
