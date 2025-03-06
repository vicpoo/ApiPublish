package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Mesas/application"
)

type ViewAllMesasController struct {
	useCase *application.ViewAllMesasUseCase
}

func NewViewAllMesasController(useCase *application.ViewAllMesasUseCase) *ViewAllMesasController {
	return &ViewAllMesasController{useCase: useCase}
}

func (vec *ViewAllMesasController) Execute(c *gin.Context) {
	mesas, err := vec.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las mesas"})
		return
	}

	c.JSON(http.StatusOK, mesas)
}
