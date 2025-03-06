package application

import "github.com/vicpoo/ApiPublish/src/Mesas/domain"

type DeleteMesaUseCase struct {
	db domain.IMesa
}

func NewDeleteMesaUseCase(db domain.IMesa) *DeleteMesaUseCase {
	return &DeleteMesaUseCase{db: db}
}

func (uc *DeleteMesaUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
