package application

import (
	"github.com/vicpoo/ApiPublish/src/Mesas/domain"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type UpdateMesaUseCase struct {
	db domain.IMesa
}

func NewUpdateMesaUseCase(db domain.IMesa) *UpdateMesaUseCase {
	return &UpdateMesaUseCase{db: db}
}

func (uc *UpdateMesaUseCase) Run(id int, mesa entities.Mesa) error {
	return uc.db.Update(id, mesa)
}
