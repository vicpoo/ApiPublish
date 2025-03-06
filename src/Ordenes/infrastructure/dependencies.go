package infrastructure

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

func InitOrdenDependencies() (
	*CreateOrdenController,
	*ViewOrdenController,
	*UpdateOrdenController,
	*DeleteOrdenController,
	*ViewAllOrdenesController,
) {
	// Inicializar el repositorio
	repo := NewMySQLOrdenRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateOrdenUseCase(repo)
	viewUseCase := application.NewViewOrdenUseCase(repo)
	updateUseCase := application.NewUpdateOrdenUseCase(repo)
	deleteUseCase := application.NewDeleteOrdenUseCase(repo)
	viewAllUseCase := application.NewViewAllOrdenesUseCase(repo)

	// Crear controladores
	createController := NewCreateOrdenController(createUseCase)
	viewController := NewViewOrdenController(viewUseCase)
	updateController := NewUpdateOrdenController(updateUseCase)
	deleteController := NewDeleteOrdenController(deleteUseCase)
	viewAllController := NewViewAllOrdenesController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
