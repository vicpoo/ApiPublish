package entities

type DetallesOrden struct {
	ID         int `json:"id"`
	OrdenID    int `json:"orden_id"`
	PlatilloID int `json:"platillo_id"`
	Cantidad   int `json:"cantidad"`
}

func NewDetallesOrden(id, ordenID, platilloID, cantidad int) *DetallesOrden {
	return &DetallesOrden{
		ID:         id,
		OrdenID:    ordenID,
		PlatilloID: platilloID,
		Cantidad:   cantidad,
	}
}

func (d *DetallesOrden) GetID() int {
	return d.ID
}

func (d *DetallesOrden) SetID(id int) {
	d.ID = id
}

func (d *DetallesOrden) GetOrdenID() int {
	return d.OrdenID
}

func (d *DetallesOrden) SetOrdenID(ordenID int) {
	d.OrdenID = ordenID
}

func (d *DetallesOrden) GetPlatilloID() int {
	return d.PlatilloID
}

func (d *DetallesOrden) SetPlatilloID(platilloID int) {
	d.PlatilloID = platilloID
}

func (d *DetallesOrden) GetCantidad() int {
	return d.Cantidad
}

func (d *DetallesOrden) SetCantidad(cantidad int) {
	d.Cantidad = cantidad
}
