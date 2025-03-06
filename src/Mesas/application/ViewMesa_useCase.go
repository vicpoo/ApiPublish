package application

import (
	"github.com/vicpoo/ApiPublish/src/Mesas/domain"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type ViewMesaUseCase struct {
	db domain.IMesa
}

func NewViewMesaUseCase(db domain.IMesa) *ViewMesaUseCase {
	return &ViewMesaUseCase{db: db}
}

func (uc *ViewMesaUseCase) Run(id int) (entities.Mesa, error) {
	return uc.db.GetByNumeroMesa(id)
}
