package application

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type ViewAllPlatillosUseCase struct {
	db domain.IPlatillo
}

func NewViewAllPlatillosUseCase(db domain.IPlatillo) *ViewAllPlatillosUseCase {
	return &ViewAllPlatillosUseCase{db: db}
}

func (uc *ViewAllPlatillosUseCase) Run() ([]entities.Platillo, error) {
	return uc.db.GetAll()
}
