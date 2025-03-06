package application

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type UpdateDetallesOrdenUseCase struct {
	db domain.IDetallesOrden
}

func NewUpdateDetallesOrdenUseCase(db domain.IDetallesOrden) *UpdateDetallesOrdenUseCase {
	return &UpdateDetallesOrdenUseCase{db: db}
}

func (uc *UpdateDetallesOrdenUseCase) Execute(id int, detalle entities.DetallesOrden) error {
	return uc.db.Update(id, detalle)
}
