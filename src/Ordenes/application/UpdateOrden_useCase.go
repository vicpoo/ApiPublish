package application

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type UpdateOrdenUseCase struct {
	db domain.IOrden
}

func NewUpdateOrdenUseCase(db domain.IOrden) *UpdateOrdenUseCase {
	return &UpdateOrdenUseCase{db: db}
}

func (uc *UpdateOrdenUseCase) Run(id int, orden entities.Orden) error {
	return uc.db.Update(id, orden)
}
