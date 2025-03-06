package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
)

type ViewAllDetallesOrdenController struct {
	useCase *application.ViewAllDetallesOrdenUseCase
}

func NewViewAllDetallesOrdenController(useCase *application.ViewAllDetallesOrdenUseCase) *ViewAllDetallesOrdenController {
	return &ViewAllDetallesOrdenController{useCase: useCase}
}

func (ctrl *ViewAllDetallesOrdenController) Execute(c *gin.Context) {
	detalles, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los detalles de orden"})
		return
	}

	c.JSON(http.StatusOK, detalles)
}
