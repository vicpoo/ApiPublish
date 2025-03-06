package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
)

type GetDetallesOrdenByOrdenIDController struct {
	useCase *application.GetDetallesOrdenByOrdenIDUseCase
}

func NewGetDetallesOrdenByOrdenIDController(useCase *application.GetDetallesOrdenByOrdenIDUseCase) *GetDetallesOrdenByOrdenIDController {
	return &GetDetallesOrdenByOrdenIDController{useCase: useCase}
}

func (ctrl *GetDetallesOrdenByOrdenIDController) Run(c *gin.Context) {
	ordenID, err := strconv.Atoi(c.Param("orden_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de orden inválido"})
		return
	}

	detalles, err := ctrl.useCase.Run(ordenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los detalles de orden"})
		return
	}

	c.JSON(http.StatusOK, detalles)
}
