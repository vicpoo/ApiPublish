package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
)

type ViewPlatilloController struct {
	useCase *application.ViewPlatilloUseCase
}

func NewViewPlatilloController(useCase *application.ViewPlatilloUseCase) *ViewPlatilloController {
	return &ViewPlatilloController{useCase: useCase}
}

func (vpc *ViewPlatilloController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	platillo, err := vpc.useCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Platillo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, platillo)
}
