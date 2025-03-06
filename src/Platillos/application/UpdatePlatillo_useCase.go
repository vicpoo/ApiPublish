package application

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type UpdatePlatilloUseCase struct {
	db domain.IPlatillo
}

func NewUpdatePlatilloUseCase(db domain.IPlatillo) *UpdatePlatilloUseCase {
	return &UpdatePlatilloUseCase{db: db}
}

func (uc *UpdatePlatilloUseCase) Run(id int, platillo entities.Platillo) error {
	return uc.db.Update(id, platillo)
}
