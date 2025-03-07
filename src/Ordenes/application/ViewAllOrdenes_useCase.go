// ViewAllOrdenes_useCase.go
package application

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type ViewAllOrdenesUseCase struct {
	db domain.IOrden
}

func NewViewAllOrdenesUseCase(db domain.IOrden) *ViewAllOrdenesUseCase {
	return &ViewAllOrdenesUseCase{db: db}
}

func (uc *ViewAllOrdenesUseCase) Run() ([]entities.Orden, error) {
	return uc.db.GetAll()
}
