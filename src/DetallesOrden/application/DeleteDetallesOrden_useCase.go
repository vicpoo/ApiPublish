package application

import "github.com/vicpoo/ApiPublish/src/DetallesOrden/domain"

type DeleteDetallesOrdenUseCase struct {
	db domain.IDetallesOrden
}

func NewDeleteDetallesOrdenUseCase(db domain.IDetallesOrden) *DeleteDetallesOrdenUseCase {
	return &DeleteDetallesOrdenUseCase{
		db: db,
	}
}

func (uc *DeleteDetallesOrdenUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
