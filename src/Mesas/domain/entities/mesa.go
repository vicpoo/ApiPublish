package entities

type Mesa struct {
	ID         int `json:"id"`
	NumeroMesa int `json:"numero_mesa"`
}

func NewMesa(id int, numeroMesa int) *Mesa {
	return &Mesa{
		ID:         id,
		NumeroMesa: numeroMesa,
	}
}

func (m *Mesa) GetID() int {
	return m.ID
}

func (m *Mesa) SetID(id int) {
	m.ID = id
}

func (m *Mesa) GetNumeroMesa() int {
	return m.NumeroMesa
}

func (m *Mesa) SetNumeroMesa(numeroMesa int) {
	m.NumeroMesa = numeroMesa
}
