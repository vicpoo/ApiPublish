// ViewOrden_useCase.go
package application

import (
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain"
	"github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"
)

type ViewOrdenUseCase struct {
	db domain.IOrden
}

func NewViewOrdenUseCase(db domain.IOrden) *ViewOrdenUseCase {
	return &ViewOrdenUseCase{db: db}
}

func (uc *ViewOrdenUseCase) Run(id int) (entities.Orden, error) {
	return uc.db.FindByID(id)
}
