// CreateOrden_useCase.go
package application

import (
	"encoding/json"

	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type CreateOrdenUseCase struct {
	db        domain.IOrden
	messaging domain.MessagingService
}

func NewCreateOrdenUseCase(db domain.IOrden, messaging domain.MessagingService) *CreateOrdenUseCase {
	return &CreateOrdenUseCase{db: db, messaging: messaging}
}

func (uc *CreateOrdenUseCase) Run(orden *entities.Orden) (*entities.Orden, error) {
	err := uc.db.Save(*orden)
	if err != nil {
		return nil, err
	}

	orderJSON, err := json.Marshal(orden)
	if err != nil {
		return nil, err
	}

	err = uc.messaging.PublishOrderCreated(orderJSON)
	if err != nil {
		return nil, err
	}

	return orden, nil
}
