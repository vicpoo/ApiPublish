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
	createController, viewController, updateController, deleteController, viewAllController := InitMesaDependencies()

	// Grupo de rutas para mesas
	mesaGroup := router.engine.Group("/mesas")
	{
		mesaGroup.POST("/", createController.Run)
		mesaGroup.GET("/:id", viewController.Execute)
		mesaGroup.PUT("/:id", updateController.Execute)
		mesaGroup.DELETE("/:id", deleteController.Run)
		mesaGroup.GET("/", viewAllController.Execute)
	}
}
