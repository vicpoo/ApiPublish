package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
)

type ViewAllPlatillosController struct {
	useCase *application.ViewAllPlatillosUseCase
}

func NewViewAllPlatillosController(useCase *application.ViewAllPlatillosUseCase) *ViewAllPlatillosController {
	return &ViewAllPlatillosController{useCase: useCase}
}

func (vapc *ViewAllPlatillosController) Execute(c *gin.Context) {
	platillos, err := vapc.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los platillos"})
		return
	}

	c.JSON(http.StatusOK, platillos)
}
