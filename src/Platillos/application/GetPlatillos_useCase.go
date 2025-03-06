package application

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type GetPlatillosUseCase struct {
	db domain.IPlatillo
}

func NewGetPlatillosUseCase(db domain.IPlatillo) *GetPlatillosUseCase {
	return &GetPlatillosUseCase{db: db}
}

func (uc *GetPlatillosUseCase) Run() ([]entities.Platillo, error) {
	return uc.db.GetAll()
}
