package application

import (
	"github.com/vicpoo/ApiPublish/src/Platillos/domain"
	"github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"
)

type CreatePlatilloUseCase struct {
	db domain.IPlatillo
}

func NewCreatePlatilloUseCase(db domain.IPlatillo) *CreatePlatilloUseCase {
	return &CreatePlatilloUseCase{db: db}
}

func (uc *CreatePlatilloUseCase) Run(platillo *entities.Platillo) (*entities.Platillo, error) {
	err := uc.db.Save(*platillo)
	return platillo, err
}
