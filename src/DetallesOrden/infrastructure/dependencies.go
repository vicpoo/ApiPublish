package infrastructure

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/application"
)

func InitDetallesOrdenDependencies() (
	*CreateDetallesOrdenController,
	*ViewDetallesOrdenController,
	*UpdateDetallesOrdenController,
	*DeleteDetallesOrdenController,
	*ViewAllDetallesOrdenController,
	*GetDetallesOrdenByOrdenIDController,
) {
	repo := NewMysqlDetallesOrdenRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateDetallesOrdenUseCase(repo)
	viewUseCase := application.NewViewDetallesOrdenUseCase(repo)
	updateUseCase := application.NewUpdateDetallesOrdenUseCase(repo)
	deleteUseCase := application.NewDeleteDetallesOrdenUseCase(repo)
	viewAllUseCase := application.NewViewAllDetallesOrdenUseCase(repo)
	getByOrdenIDUseCase := application.NewGetDetallesOrdenByOrdenIDUseCase(repo)

	// Crear controladores
	createController := NewCreateDetallesOrdenController(createUseCase)
	viewController := NewViewDetallesOrdenController(viewUseCase)
	updateController := NewUpdateDetallesOrdenController(updateUseCase)
	deleteController := NewDeleteDetallesOrdenController(deleteUseCase)
	viewAllController := NewViewAllDetallesOrdenController(viewAllUseCase)
	getByOrdenIDController := NewGetDetallesOrdenByOrdenIDController(getByOrdenIDUseCase)

	return createController, viewController, updateController, deleteController, viewAllController, getByOrdenIDController
}
