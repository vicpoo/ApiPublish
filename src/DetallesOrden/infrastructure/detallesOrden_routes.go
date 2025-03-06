package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	// Inicializar dependencias
	createController, viewController, updateController, deleteController, viewAllController, getByOrdenIDController := InitDetallesOrdenDependencies()

	// Grupo de rutas para detalles de orden
	detallesOrdenGroup := router.engine.Group("/detalles-orden")
	{
		detallesOrdenGroup.POST("/", createController.Run)
		detallesOrdenGroup.GET("/:id", viewController.Execute)
		detallesOrdenGroup.PUT("/:id", updateController.Execute)
		detallesOrdenGroup.DELETE("/:id", deleteController.Run)
		detallesOrdenGroup.GET("/", viewAllController.Execute)
		detallesOrdenGroup.GET("/orden/:orden_id", getByOrdenIDController.Run)
	}
}
