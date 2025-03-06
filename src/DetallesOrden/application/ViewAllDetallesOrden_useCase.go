package application

import (
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"
	"github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"
)

type ViewAllDetallesOrdenUseCase struct {
	db domain.IDetallesOrden
}

func NewViewAllDetallesOrdenUseCase(db domain.IDetallesOrden) *ViewAllDetallesOrdenUseCase {
	return &ViewAllDetallesOrdenUseCase{db: db}
}

func (uc *ViewAllDetallesOrdenUseCase) Execute() ([]entities.DetallesOrden, error) {
	return uc.db.GetAll()
}
