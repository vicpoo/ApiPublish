package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Mesas/application"
)

type ViewMesaController struct {
	useCase *application.ViewMesaUseCase
}

func NewViewMesaController(useCase *application.ViewMesaUseCase) *ViewMesaController {
	return &ViewMesaController{useCase: useCase}
}

func (vec *ViewMesaController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	mesa, err := vec.useCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mesa no encontrada"})
		return
	}

	c.JSON(http.StatusOK, mesa)
}
