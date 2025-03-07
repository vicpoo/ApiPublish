// DeleteOrden_useCase.go
package application

import "github.com/vicpoo/ApiPublish/src/Ordenes/domain"

type DeleteOrdenUseCase struct {
	db domain.IOrden
}

func NewDeleteOrdenUseCase(db domain.IOrden) *DeleteOrdenUseCase {
	return &DeleteOrdenUseCase{db: db}
}

func (uc *DeleteOrdenUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
