package application

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type CreateDetallesOrdenUseCase struct {
	db domain.IDetallesOrden
}

func NewCreateDetallesOrdenUseCase(db domain.IDetallesOrden) *CreateDetallesOrdenUseCase {
	return &CreateDetallesOrdenUseCase{
		db: db,
	}
}

func (uc *CreateDetallesOrdenUseCase) Run(detalle *entities.DetallesOrden) (*entities.DetallesOrden, error) {
	err := uc.db.Save(*detalle)
	return detalle, err
}
