package application

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type ViewPlatilloUseCase struct {
	db domain.IPlatillo
}

func NewViewPlatilloUseCase(db domain.IPlatillo) *ViewPlatilloUseCase {
	return &ViewPlatilloUseCase{db: db}
}

func (uc *ViewPlatilloUseCase) Run(id int) (entities.Platillo, error) {
	return uc.db.FindByID(id)
}
