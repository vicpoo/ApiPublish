package domain

import "github.com/vicpoo/ApiPublish/src/Platillos/domain/entities"

type IPlatillo interface {
	Save(platillo entities.Platillo) error
	Update(id int, platillo entities.Platillo) error
	Delete(id int) error
	FindByID(id int) (entities.Platillo, error)
	GetAll() ([]entities.Platillo, error)
	GetByNombre(nombre string) (entities.Platillo, error)
}
