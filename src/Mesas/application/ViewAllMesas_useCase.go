package application

import (
	"github.com/vicpoo/ApiPublish/src/Mesas/domain"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type ViewAllMesasUseCase struct {
	db domain.IMesa
}

func NewViewAllMesasUseCase(db domain.IMesa) *ViewAllMesasUseCase {
	return &ViewAllMesasUseCase{db: db}
}

func (uc *ViewAllMesasUseCase) Run() ([]entities.Mesa, error) {
	return uc.db.GetAll()
}
