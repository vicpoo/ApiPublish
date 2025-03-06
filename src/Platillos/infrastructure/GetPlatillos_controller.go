package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
)

type GetPlatillosController struct {
	useCase *application.GetPlatillosUseCase
}

func NewGetPlatillosController(useCase *application.GetPlatillosUseCase) *GetPlatillosController {
	return &GetPlatillosController{useCase: useCase}
}

func (gpc *GetPlatillosController) Run(c *gin.Context) {
	platillos, err := gpc.useCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los platillos"})
		return
	}

	c.JSON(http.StatusOK, platillos)
}
