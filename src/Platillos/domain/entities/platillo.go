package entities

type Platillo struct {
	ID          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
}

func NewPlatillo(id int, nombre, descripcion string, precio float64) *Platillo {
	return &Platillo{
		ID:          id,
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
	}
}

// Getters
func (p *Platillo) GetID() int {
	return p.ID
}

func (p *Platillo) GetNombre() string {
	return p.Nombre
}

func (p *Platillo) GetDescripcion() string {
	return p.Descripcion
}

func (p *Platillo) GetPrecio() float64 {
	return p.Precio
}

// Setters
func (p *Platillo) SetID(id int) {
	p.ID = id
}

func (p *Platillo) SetNombre(nombre string) {
	p.Nombre = nombre
}

func (p *Platillo) SetDescripcion(descripcion string) {
	p.Descripcion = descripcion
}

func (p *Platillo) SetPrecio(precio float64) {
	p.Precio = precio
}
