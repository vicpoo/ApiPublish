package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

type GetOrdenesController struct {
	useCase *application.GetOrdenesUseCase
}

func NewGetOrdenesController(useCase *application.GetOrdenesUseCase) *GetOrdenesController {
	return &GetOrdenesController{useCase: useCase}
}

func (goc *GetOrdenesController) Run(c *gin.Context) {
	ordenes, err := goc.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las órdenes"})
		return
	}

	c.JSON(http.StatusOK, ordenes)
}
