package application

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type GetOrdenesUseCase struct {
	db domain.IOrden
}

func NewGetOrdenesUseCase(db domain.IOrden) *GetOrdenesUseCase {
	return &GetOrdenesUseCase{db: db}
}

func (uc *GetOrdenesUseCase) Run() ([]entities.Orden, error) {
	return uc.db.GetAll()
}
