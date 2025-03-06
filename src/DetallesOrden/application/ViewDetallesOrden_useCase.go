package application

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type ViewDetallesOrdenUseCase struct {
	db domain.IDetallesOrden
}

func NewViewDetallesOrdenUseCase(db domain.IDetallesOrden) *ViewDetallesOrdenUseCase {
	return &ViewDetallesOrdenUseCase{db: db}
}

func (uc *ViewDetallesOrdenUseCase) Execute(id int) (entities.DetallesOrden, error) {
	return uc.db.FindByID(id)
}
