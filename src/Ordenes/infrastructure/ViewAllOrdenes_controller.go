// ViewAllOrdenes_controller.go
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

type ViewAllOrdenesController struct {
	useCase *application.ViewAllOrdenesUseCase
}

func NewViewAllOrdenesController(useCase *application.ViewAllOrdenesUseCase) *ViewAllOrdenesController {
	return &ViewAllOrdenesController{useCase: useCase}
}

func (vaoc *ViewAllOrdenesController) Execute(c *gin.Context) {
	ordenes, err := vaoc.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las órdenes"})
		return
	}

	c.JSON(http.StatusOK, ordenes)
}
