// ordenes_routes.go
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
	createController, viewController, updateController, deleteController, viewAllController := InitOrdenDependencies()

	// Grupo de rutas para órdenes
	ordenGroup := router.engine.Group("/ordenes")
	{
		ordenGroup.POST("/", createController.Run)
		ordenGroup.GET("/:id", viewController.Execute)
		ordenGroup.PUT("/:id", updateController.Execute)
		ordenGroup.DELETE("/:id", deleteController.Run)
		ordenGroup.GET("/", viewAllController.Execute)
	}
}
