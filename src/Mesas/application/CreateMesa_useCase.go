package application

import (
	"github.com/vicpoo/ApiPublish/src/Mesas/domain"
	"github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"
)

type CreateMesaUseCase struct {
	db domain.IMesa
}

func NewCreateMesaUseCase(db domain.IMesa) *CreateMesaUseCase {
	return &CreateMesaUseCase{db: db}
}

func (uc *CreateMesaUseCase) Run(mesa *entities.Mesa) (*entities.Mesa, error) {
	err := uc.db.Save(*mesa)
	return mesa, err
}
