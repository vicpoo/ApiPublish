package domain

import "github.com/vicpoo/ApiPublish/src/Mesas/domain/entities"

type IMesa interface {
	Save(mesa entities.Mesa) error
	Update(id int, mesa entities.Mesa) error
	Delete(id int) error
	GetAll() ([]entities.Mesa, error)
	GetByNumeroMesa(numeroMesa int) (entities.Mesa, error)
}
