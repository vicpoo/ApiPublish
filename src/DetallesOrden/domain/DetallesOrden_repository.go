package domain

import "github.com/vicpoo/ApiPublish/src/DetallesOrden/domain/entities"

type IDetallesOrden interface {
	Save(detalle entities.DetallesOrden) error
	Update(id int, detalle entities.DetallesOrden) error
	Delete(id int) error
	FindByID(id int) (entities.DetallesOrden, error)
	GetAll() ([]entities.DetallesOrden, error)
	GetByOrdenID(ordenID int) ([]entities.DetallesOrden, error)
}
