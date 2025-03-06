package infrastructure

import (
	"github.com/vicpoo/ApiPublish/src/Mesas/application"
)

func InitMesaDependencies() (
	*CreateMesaController,
	*ViewMesaController,
	*UpdateMesaController,
	*DeleteMesaController,
	*ViewAllMesasController,
) {
	// Inicializar el repositorio
	repo := NewMySQLMesaRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateMesaUseCase(repo)
	viewUseCase := application.NewViewMesaUseCase(repo)
	updateUseCase := application.NewUpdateMesaUseCase(repo)
	deleteUseCase := application.NewDeleteMesaUseCase(repo)
	viewAllUseCase := application.NewViewAllMesasUseCase(repo)

	// Crear controladores
	createController := NewCreateMesaController(createUseCase)
	viewController := NewViewMesaController(viewUseCase)
	updateController := NewUpdateMesaController(updateUseCase)
	deleteController := NewDeleteMesaController(deleteUseCase)
	viewAllController := NewViewAllMesasController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
