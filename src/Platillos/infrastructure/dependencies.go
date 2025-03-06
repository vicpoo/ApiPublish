package infrastructure

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/application"
)

func InitPlatilloDependencies() (
	*CreatePlatilloController,
	*ViewPlatilloController,
	*UpdatePlatilloController,
	*DeletePlatilloController,
	*ViewAllPlatillosController,
) {
	// Inicializar el repositorio
	repo := NewMySQLPlatilloRepository()

	// Crear casos de uso
	createUseCase := application.NewCreatePlatilloUseCase(repo)
	viewUseCase := application.NewViewPlatilloUseCase(repo)
	updateUseCase := application.NewUpdatePlatilloUseCase(repo)
	deleteUseCase := application.NewDeletePlatilloUseCase(repo)
	viewAllUseCase := application.NewViewAllPlatillosUseCase(repo)

	// Crear controladores
	createController := NewCreatePlatilloController(createUseCase)
	viewController := NewViewPlatilloController(viewUseCase)
	updateController := NewUpdatePlatilloController(updateUseCase)
	deleteController := NewDeletePlatilloController(deleteUseCase)
	viewAllController := NewViewAllPlatillosController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
