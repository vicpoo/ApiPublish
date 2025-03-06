package domain

import "github.com/vicpoo/ApiPublish/src/Ordenes/domain/entities"

type IOrden interface {
	Save(orden entities.Orden) error
	Update(id int, orden entities.Orden) error
	Delete(id int) error
	FindByID(id int) (entities.Orden, error)
	GetAll() ([]entities.Orden, error)
	GetByMesaID(mesaID int) ([]entities.Orden, error)
}
