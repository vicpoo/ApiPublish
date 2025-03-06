package application

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type CreateOrdenUseCase struct {
	db domain.IOrden
}

func NewCreateOrdenUseCase(db domain.IOrden) *CreateOrdenUseCase {
	return &CreateOrdenUseCase{db: db}
}

func (uc *CreateOrdenUseCase) Run(orden *entities.Orden) (*entities.Orden, error) {
	err := uc.db.Save(*orden)
	return orden, err
}
