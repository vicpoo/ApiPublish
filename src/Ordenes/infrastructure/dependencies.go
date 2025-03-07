// dependencies.go
package infrastructure

import (
	"log"

	"github.com/vicpoo/ApiPublish/src/Ordenes/application"
)

func InitOrdenDependencies() (
	*CreateOrdenController,
	*ViewOrdenController,
	*UpdateOrdenController,
	*DeleteOrdenController,
	*ViewAllOrdenesController,
) {
	repo := NewMySQLOrdenRepository()

	messaging, err := NewMessagingService()
	if err != nil {
		log.Fatalf("Failed to initialize messaging service: %s", err)
	}

	createUseCase := application.NewCreateOrdenUseCase(repo, messaging)
	viewUseCase := application.NewViewOrdenUseCase(repo)
	updateUseCase := application.NewUpdateOrdenUseCase(repo)
	deleteUseCase := application.NewDeleteOrdenUseCase(repo)
	viewAllUseCase := application.NewViewAllOrdenesUseCase(repo)

	createController := NewCreateOrdenController(createUseCase)
	viewController := NewViewOrdenController(viewUseCase)
	updateController := NewUpdateOrdenController(updateUseCase)
	deleteController := NewDeleteOrdenController(deleteUseCase)
	viewAllController := NewViewAllOrdenesController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
