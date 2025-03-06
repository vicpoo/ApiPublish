package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	detallesOrdenInfra "github.com/vicpoo/ApiPublish/src/DetallesOrden/infrastructure"
	mesasInfra "github.com/vicpoo/ApiPublish/src/Mesas/infrastructure"
	ordenesInfra "github.com/vicpoo/ApiPublish/src/Ordenes/infrastructure"
	platillosInfra "github.com/vicpoo/ApiPublish/src/Platillos/infrastructure"
	"github.com/vicpoo/ApiPublish/src/core"
)

func main() {
	// Inicializar la base de datos
	core.InitDB()

	// Crear un router con Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Cambia esto según la URL de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configuración de rutas para DetallesOrden
	detallesOrdenRouter := detallesOrdenInfra.NewRouter(r)
	detallesOrdenRouter.Run()

	// Configuración de rutas para Mesas
	mesasRouter := mesasInfra.NewRouter(r)
	mesasRouter.Run()

	// Configuración de rutas para Ordenes
	ordenesRouter := ordenesInfra.NewRouter(r)
	ordenesRouter.Run()

	// Configuración de rutas para Platillos
	platillosRouter := platillosInfra.NewRouter(r)
	platillosRouter.Run()

	// Mensaje de inicio
	fmt.Println("¡API en Funcionamiento :D!")

	// Iniciar el servidor en el puerto 8000
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
