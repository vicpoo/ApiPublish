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
	createController, viewController, updateController, deleteController, viewAllController := InitPlatilloDependencies()

	// Grupo de rutas para platillos
	platilloGroup := router.engine.Group("/platillos")
	{
		platilloGroup.POST("/", createController.Run)
		platilloGroup.GET("/:id", viewController.Execute)
		platilloGroup.PUT("/:id", updateController.Execute)
		platilloGroup.DELETE("/:id", deleteController.Run)
		platilloGroup.GET("/", viewAllController.Execute)
	}
}
