package application

import "github.com/vicpoo/ApiPublish/src/Platillos/domain"

type DeletePlatilloUseCase struct {
	db domain.IPlatillo
}

func NewDeletePlatilloUseCase(db domain.IPlatillo) *DeletePlatilloUseCase {
	return &DeletePlatilloUseCase{db: db}
}

func (uc *DeletePlatilloUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
