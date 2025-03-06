package application

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type GetDetallesOrdenByOrdenIDUseCase struct {
	db domain.IDetallesOrden
}

func NewGetDetallesOrdenByOrdenIDUseCase(db domain.IDetallesOrden) *GetDetallesOrdenByOrdenIDUseCase {
	return &GetDetallesOrdenByOrdenIDUseCase{
		db: db,
	}
}

func (uc *GetDetallesOrdenByOrdenIDUseCase) Run(ordenID int) ([]entities.DetallesOrden, error) {
	return uc.db.GetByOrdenID(ordenID)
}
